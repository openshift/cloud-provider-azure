//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerLoadBalancingRulesClient contains the methods for the LoadBalancerLoadBalancingRules group.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesClient() instead.
type LoadBalancerLoadBalancingRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLoadBalancerLoadBalancingRulesClient creates a new instance of LoadBalancerLoadBalancingRulesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancerLoadBalancingRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancerLoadBalancingRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancerLoadBalancingRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the specified load balancer load balancing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - loadBalancingRuleName - The name of the load balancing rule.
//   - options - LoadBalancerLoadBalancingRulesClientGetOptions contains the optional parameters for the LoadBalancerLoadBalancingRulesClient.Get
//     method.
func (client *LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientGetOptions) (LoadBalancerLoadBalancingRulesClientGetResponse, error) {
	var err error
	const operationName = "LoadBalancerLoadBalancingRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, loadBalancingRuleName, options)
	if err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerLoadBalancingRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if loadBalancingRuleName == "" {
		return nil, errors.New("parameter loadBalancingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerLoadBalancingRulesClient) getHandleResponse(resp *http.Response) (LoadBalancerLoadBalancingRulesClientGetResponse, error) {
	result := LoadBalancerLoadBalancingRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancingRule); err != nil {
		return LoadBalancerLoadBalancingRulesClientGetResponse{}, err
	}
	return result, nil
}

// BeginHealth - Get health details of a load balancing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - groupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - loadBalancingRuleName - The name of the load balancing rule.
//   - options - LoadBalancerLoadBalancingRulesClientBeginHealthOptions contains the optional parameters for the LoadBalancerLoadBalancingRulesClient.BeginHealth
//     method.
func (client *LoadBalancerLoadBalancingRulesClient) BeginHealth(ctx context.Context, groupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientBeginHealthOptions) (*runtime.Poller[LoadBalancerLoadBalancingRulesClientHealthResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.health(ctx, groupName, loadBalancerName, loadBalancingRuleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancerLoadBalancingRulesClientHealthResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LoadBalancerLoadBalancingRulesClientHealthResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Health - Get health details of a load balancing rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
func (client *LoadBalancerLoadBalancingRulesClient) health(ctx context.Context, groupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientBeginHealthOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancerLoadBalancingRulesClient.BeginHealth"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.healthCreateRequest(ctx, groupName, loadBalancerName, loadBalancingRuleName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// healthCreateRequest creates the Health request.
func (client *LoadBalancerLoadBalancingRulesClient) healthCreateRequest(ctx context.Context, groupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesClientBeginHealthOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}/health"
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if loadBalancingRuleName == "" {
		return nil, errors.New("parameter loadBalancingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Gets all the load balancing rules in a load balancer.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancerLoadBalancingRulesClientListOptions contains the optional parameters for the LoadBalancerLoadBalancingRulesClient.NewListPager
//     method.
func (client *LoadBalancerLoadBalancingRulesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesClientListOptions) *runtime.Pager[LoadBalancerLoadBalancingRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancerLoadBalancingRulesClientListResponse]{
		More: func(page LoadBalancerLoadBalancingRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancerLoadBalancingRulesClientListResponse) (LoadBalancerLoadBalancingRulesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancerLoadBalancingRulesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			}, nil)
			if err != nil {
				return LoadBalancerLoadBalancingRulesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancerLoadBalancingRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerLoadBalancingRulesClient) listHandleResponse(resp *http.Response) (LoadBalancerLoadBalancingRulesClientListResponse, error) {
	result := LoadBalancerLoadBalancingRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerLoadBalancingRuleListResult); err != nil {
		return LoadBalancerLoadBalancingRulesClientListResponse{}, err
	}
	return result, nil
}