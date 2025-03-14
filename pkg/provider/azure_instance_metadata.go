/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"k8s.io/klog/v2"

	azcache "sigs.k8s.io/cloud-provider-azure/pkg/cache"
	"sigs.k8s.io/cloud-provider-azure/pkg/consts"
)

// NetworkMetadata contains metadata about an instance's network
type NetworkMetadata struct {
	Interface []*NetworkInterface `json:"interface"`
}

// NetworkInterface represents an instances network interface.
type NetworkInterface struct {
	IPV4 NetworkData `json:"ipv4"`
	IPV6 NetworkData `json:"ipv6"`
	MAC  string      `json:"macAddress"`
}

// NetworkData contains IP information for a armnetwork.
type NetworkData struct {
	IPAddress []IPAddress `json:"ipAddress"`
	Subnet    []Subnet    `json:"subnet"`
}

// IPAddress represents IP address information.
type IPAddress struct {
	PrivateIP string `json:"privateIpAddress"`
	PublicIP  string `json:"publicIpAddress"`
}

// Subnet represents subnet information.
type Subnet struct {
	Address string `json:"address"`
	Prefix  string `json:"prefix"`
}

// ComputeMetadata represents compute information
type ComputeMetadata struct {
	Environment            string `json:"azEnvironment,omitempty"`
	SKU                    string `json:"SKU,omitempty"`
	Name                   string `json:"name,omitempty"`
	Zone                   string `json:"zone,omitempty"`
	VMSize                 string `json:"vmSize,omitempty"`
	OSType                 string `json:"osType,omitempty"`
	Location               string `json:"location,omitempty"`
	FaultDomain            string `json:"platformFaultDomain,omitempty"`
	PlatformSubFaultDomain string `json:"platformSubFaultDomain,omitempty"`
	UpdateDomain           string `json:"platformUpdateDomain,omitempty"`
	ResourceGroup          string `json:"resourceGroupName,omitempty"`
	VMScaleSetName         string `json:"vmScaleSetName,omitempty"`
	SubscriptionID         string `json:"subscriptionId,omitempty"`
	ResourceID             string `json:"resourceId,omitempty"`
}

// InstanceMetadata represents instance information.
type InstanceMetadata struct {
	Compute *ComputeMetadata `json:"compute,omitempty"`
	Network *NetworkMetadata `json:"network,omitempty"`
}

// PublicIPMetadata represents the public IP metadata.
type PublicIPMetadata struct {
	FrontendIPAddress string `json:"frontendIpAddress,omitempty"`
	PrivateIPAddress  string `json:"privateIpAddress,omitempty"`
}

// LoadbalancerProfile represents load balancer profile in IMDS.
type LoadbalancerProfile struct {
	PublicIPAddresses []PublicIPMetadata `json:"publicIpAddresses,omitempty"`
}

// LoadBalancerMetadata represents load balancer metadata.
type LoadBalancerMetadata struct {
	LoadBalancer *LoadbalancerProfile `json:"loadbalancer,omitempty"`
}

// InstanceMetadataService knows how to query the Azure instance metadata server.
type InstanceMetadataService struct {
	imdsServer string
	imsCache   azcache.Resource
}

// NewInstanceMetadataService creates an instance of the InstanceMetadataService accessor object.
func NewInstanceMetadataService(imdsServer string) (*InstanceMetadataService, error) {
	ims := &InstanceMetadataService{
		imdsServer: imdsServer,
	}

	imsCache, err := azcache.NewTimedCache(consts.MetadataCacheTTL, ims.getMetadata, false)
	if err != nil {
		return nil, err
	}

	ims.imsCache = imsCache
	return ims, nil
}

// fillNetInterfacePublicIPs finds PIPs from imds load balancer and fills them into net interface config.
func fillNetInterfacePublicIPs(publicIPs []PublicIPMetadata, netInterface *NetworkInterface) {
	// IPv6 IPs from imds load balancer are wrapped by brackets while those from imds are not.
	trimIP := func(ip string) string {
		return strings.Trim(strings.Trim(ip, "["), "]")
	}

	if len(netInterface.IPV4.IPAddress) > 0 && len(netInterface.IPV4.IPAddress[0].PrivateIP) > 0 {
		for _, pip := range publicIPs {
			if pip.PrivateIPAddress == netInterface.IPV4.IPAddress[0].PrivateIP {
				netInterface.IPV4.IPAddress[0].PublicIP = pip.FrontendIPAddress
				break
			}
		}
	}
	if len(netInterface.IPV6.IPAddress) > 0 && len(netInterface.IPV6.IPAddress[0].PrivateIP) > 0 {
		for _, pip := range publicIPs {
			privateIP := trimIP(pip.PrivateIPAddress)
			frontendIP := trimIP(pip.FrontendIPAddress)
			if privateIP == netInterface.IPV6.IPAddress[0].PrivateIP {
				netInterface.IPV6.IPAddress[0].PublicIP = frontendIP
				break
			}
		}
	}
}

func (ims *InstanceMetadataService) getMetadata(_ context.Context, key string) (interface{}, error) {
	instanceMetadata, err := ims.getInstanceMetadata(key)
	if err != nil {
		return nil, err
	}

	if instanceMetadata.Network != nil && len(instanceMetadata.Network.Interface) > 0 {
		netInterface := instanceMetadata.Network.Interface[0]
		if (len(netInterface.IPV4.IPAddress) > 0 && len(netInterface.IPV4.IPAddress[0].PublicIP) > 0) ||
			(len(netInterface.IPV6.IPAddress) > 0 && len(netInterface.IPV6.IPAddress[0].PublicIP) > 0) {
			// Return if public IP address has already part of instance metadata.
			return instanceMetadata, nil
		}

		loadBalancerMetadata, err := ims.getLoadBalancerMetadata()
		if err != nil || loadBalancerMetadata == nil || loadBalancerMetadata.LoadBalancer == nil {
			// Log a warning since loadbalancer metadata may not be available when the VM
			// is not in standard LoadBalancer backend address pool.
			klog.V(4).Infof("Warning: failed to get loadbalancer metadata: %v", err)
			return instanceMetadata, nil
		}

		publicIPs := loadBalancerMetadata.LoadBalancer.PublicIPAddresses
		fillNetInterfacePublicIPs(publicIPs, netInterface)
	}

	return instanceMetadata, nil
}

func (ims *InstanceMetadataService) getInstanceMetadata(_ string) (*InstanceMetadata, error) {
	req, err := http.NewRequest("GET", ims.imdsServer+consts.ImdsInstanceURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Metadata", "True")
	req.Header.Add("User-Agent", "golang/kubernetes-cloud-provider")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", consts.ImdsInstanceAPIVersion)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failure of getting instance metadata with response %q", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := InstanceMetadata{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (ims *InstanceMetadataService) getLoadBalancerMetadata() (*LoadBalancerMetadata, error) {
	req, err := http.NewRequest("GET", ims.imdsServer+consts.ImdsLoadBalancerURI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Metadata", "True")
	req.Header.Add("User-Agent", "golang/kubernetes-cloud-provider")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", consts.ImdsLoadBalancerAPIVersion)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failure of getting loadbalancer metadata with response %q", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := LoadBalancerMetadata{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

// GetMetadata gets instance metadata from cache.
// crt determines if we can get data from stalled cache/need fresh if cache expired.
func (ims *InstanceMetadataService) GetMetadata(ctx context.Context, crt azcache.AzureCacheReadType) (*InstanceMetadata, error) {
	cache, err := ims.imsCache.Get(ctx, consts.MetadataCacheKey, crt)
	if err != nil {
		return nil, err
	}

	// Cache shouldn't be nil, but added a check in case something is wrong.
	if cache == nil {
		return nil, fmt.Errorf("failure of getting instance metadata")
	}

	if metadata, ok := cache.(*InstanceMetadata); ok {
		return metadata, nil
	}

	return nil, fmt.Errorf("failure of getting instance metadata")
}

// GetPlatformSubFaultDomain returns the PlatformSubFaultDomain from IMDS if set.
func (az *Cloud) GetPlatformSubFaultDomain(ctx context.Context) (string, error) {
	if az.UseInstanceMetadata {
		metadata, err := az.Metadata.GetMetadata(ctx, azcache.CacheReadTypeUnsafe)
		if err != nil {
			klog.Errorf("GetPlatformSubFaultDomain: failed to GetMetadata: %s", err.Error())
			return "", err
		}
		if metadata.Compute == nil {
			_ = az.Metadata.imsCache.Delete(consts.MetadataCacheKey)
			return "", errors.New("failure of getting compute information from instance metadata")
		}
		return metadata.Compute.PlatformSubFaultDomain, nil
	}
	return "", nil
}
