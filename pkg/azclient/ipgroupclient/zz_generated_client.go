// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package ipgroupclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/metrics"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.IPGroupsClient
	subscriptionID string
	tracer         tracing.Tracer
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}
	tr := options.TracingProvider.NewTracer(utils.ModuleName, utils.ModuleVersion)

	client, err := armnetwork.NewIPGroupsClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		IPGroupsClient: client,
		subscriptionID: subscriptionID,
		tracer:         tr,
	}, nil
}

const GetOperationName = "IPGroupsClient.Get"

// Get gets the IPGroup
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, expand *string) (result *armnetwork.IPGroup, err error) {
	var ops *armnetwork.IPGroupsClientGetOptions
	if expand != nil {
		ops = &armnetwork.IPGroupsClientGetOptions{Expand: expand}
	}
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "IPGroup", "get")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, GetOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := client.IPGroupsClient.Get(ctx, resourceGroupName, resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.IPGroup, nil
}

const CreateOrUpdateOperationName = "IPGroupsClient.Create"

// CreateOrUpdate creates or updates a IPGroup.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.IPGroup) (result *armnetwork.IPGroup, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "IPGroup", "create_or_update")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, CreateOrUpdateOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := utils.NewPollerWrapper(client.IPGroupsClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.IPGroup, nil
	}
	return nil, nil
}

const DeleteOperationName = "IPGroupsClient.Delete"

// Delete deletes a IPGroup by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, resourceName string) (err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "IPGroup", "delete")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, DeleteOperationName, client.tracer, nil)
	defer endSpan(err)
	_, err = utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName, resourceName, nil)).WaitforPollerResp(ctx)
	return err
}

const ListOperationName = "IPGroupsClient.List"

// List gets a list of IPGroup in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string) (result []*armnetwork.IPGroup, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "IPGroup", "list")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, ListOperationName, client.tracer, nil)
	defer endSpan(err)
	pager := client.IPGroupsClient.NewListByResourceGroupPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}