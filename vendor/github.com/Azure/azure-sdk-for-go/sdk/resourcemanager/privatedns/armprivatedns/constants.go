//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	moduleVersion = "v1.3.0"
)

// ProvisioningState - The provisioning state of the resource. This is a read-only property and any attempt to set this value
// will be ignored.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

type RecordType string

const (
	RecordTypeA     RecordType = "A"
	RecordTypeAAAA  RecordType = "AAAA"
	RecordTypeCNAME RecordType = "CNAME"
	RecordTypeMX    RecordType = "MX"
	RecordTypePTR   RecordType = "PTR"
	RecordTypeSOA   RecordType = "SOA"
	RecordTypeSRV   RecordType = "SRV"
	RecordTypeTXT   RecordType = "TXT"
)

// PossibleRecordTypeValues returns the possible values for the RecordType const type.
func PossibleRecordTypeValues() []RecordType {
	return []RecordType{
		RecordTypeA,
		RecordTypeAAAA,
		RecordTypeCNAME,
		RecordTypeMX,
		RecordTypePTR,
		RecordTypeSOA,
		RecordTypeSRV,
		RecordTypeTXT,
	}
}

// ResolutionPolicy - The resolution policy on the virtual network link. Only applicable for virtual network links to privatelink
// zones, and for A,AAAA,CNAME queries. When set to 'NxDomainRedirect', Azure DNS resolver
// falls back to public resolution if private dns query resolution results in non-existent domain response.
type ResolutionPolicy string

const (
	ResolutionPolicyDefault          ResolutionPolicy = "Default"
	ResolutionPolicyNxDomainRedirect ResolutionPolicy = "NxDomainRedirect"
)

// PossibleResolutionPolicyValues returns the possible values for the ResolutionPolicy const type.
func PossibleResolutionPolicyValues() []ResolutionPolicy {
	return []ResolutionPolicy{
		ResolutionPolicyDefault,
		ResolutionPolicyNxDomainRedirect,
	}
}

// VirtualNetworkLinkState - The status of the virtual network link to the Private DNS zone. Possible values are 'InProgress'
// and 'Done'. This is a read-only property and any attempt to set this value will be ignored.
type VirtualNetworkLinkState string

const (
	VirtualNetworkLinkStateCompleted  VirtualNetworkLinkState = "Completed"
	VirtualNetworkLinkStateInProgress VirtualNetworkLinkState = "InProgress"
)

// PossibleVirtualNetworkLinkStateValues returns the possible values for the VirtualNetworkLinkState const type.
func PossibleVirtualNetworkLinkStateValues() []VirtualNetworkLinkState {
	return []VirtualNetworkLinkState{
		VirtualNetworkLinkStateCompleted,
		VirtualNetworkLinkStateInProgress,
	}
}