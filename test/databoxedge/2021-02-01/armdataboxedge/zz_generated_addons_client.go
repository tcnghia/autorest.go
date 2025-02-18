//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"context"
	"errors"
	"fmt"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AddonsClient contains the methods for the Addons group.
// Don't use this type directly, use NewAddonsClient() instead.
type AddonsClient struct {
	con            *connection
	subscriptionID string
}

// NewAddonsClient creates a new instance of AddonsClient with the specified values.
func NewAddonsClient(con *connection, subscriptionID string) *AddonsClient {
	return &AddonsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a addon.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, options *AddonsBeginCreateOrUpdateOptions) (AddonsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, roleName, addonName, resourceGroupName, addon, options)
	if err != nil {
		return AddonsCreateOrUpdatePollerResponse{}, err
	}
	result := AddonsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AddonsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return AddonsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &AddonsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a addon.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) createOrUpdate(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, options *AddonsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, addon, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AddonsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon AddonClassification, options *AddonsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, addon)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AddonsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the addon on the device.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) BeginDelete(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsBeginDeleteOptions) (AddonsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, roleName, addonName, resourceGroupName, options)
	if err != nil {
		return AddonsDeletePollerResponse{}, err
	}
	result := AddonsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AddonsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return AddonsDeletePollerResponse{}, err
	}
	result.Poller = &AddonsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the addon on the device.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) deleteOperation(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AddonsClient) deleteCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AddonsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a specific addon by name.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) Get(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsGetOptions) (AddonsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, roleName, addonName, resourceGroupName, options)
	if err != nil {
		return AddonsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AddonsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AddonsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AddonsClient) getCreateRequest(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *AddonsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons/{addonName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if addonName == "" {
		return nil, errors.New("parameter addonName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{addonName}", url.PathEscape(addonName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AddonsClient) getHandleResponse(resp *http.Response) (AddonsGetResponse, error) {
	result := AddonsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AddonsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AddonsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByRole - Lists all the addons configured in the role.
// If the operation fails it returns the *CloudError error type.
func (client *AddonsClient) ListByRole(deviceName string, roleName string, resourceGroupName string, options *AddonsListByRoleOptions) *AddonsListByRolePager {
	return &AddonsListByRolePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByRoleCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp AddonsListByRoleResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AddonList.NextLink)
		},
	}
}

// listByRoleCreateRequest creates the ListByRole request.
func (client *AddonsClient) listByRoleCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *AddonsListByRoleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/addons"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRoleHandleResponse handles the ListByRole response.
func (client *AddonsClient) listByRoleHandleResponse(resp *http.Response) (AddonsListByRoleResponse, error) {
	result := AddonsListByRoleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddonList); err != nil {
		return AddonsListByRoleResponse{}, err
	}
	return result, nil
}

// listByRoleHandleError handles the ListByRole error response.
func (client *AddonsClient) listByRoleHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
