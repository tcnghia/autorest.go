//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectionMonitorsClient contains the methods for the ConnectionMonitors group.
// Don't use this type directly, use NewConnectionMonitorsClient() instead.
type ConnectionMonitorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConnectionMonitorsClient creates a new instance of ConnectionMonitorsClient with the specified values.
func NewConnectionMonitorsClient(con *arm.Connection, subscriptionID string) *ConnectionMonitorsClient {
	return &ConnectionMonitorsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (ConnectionMonitorsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return ConnectionMonitorsCreateOrUpdatePollerResponse{}, err
	}
	result := ConnectionMonitorsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectionMonitorsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ConnectionMonitorsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ConnectionMonitorsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionMonitorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters ConnectionMonitor, options *ConnectionMonitorsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ConnectionMonitorsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (ConnectionMonitorsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorsDeletePollerResponse{}, err
	}
	result := ConnectionMonitorsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectionMonitorsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ConnectionMonitorsDeletePollerResponse{}, err
	}
	result.Poller = &ConnectionMonitorsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionMonitorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConnectionMonitorsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a connection monitor by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (ConnectionMonitorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionMonitorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionMonitorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionMonitorsClient) getHandleResponse(resp *http.Response) (ConnectionMonitorsGetResponse, error) {
	result := ConnectionMonitorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorResult); err != nil {
		return ConnectionMonitorsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConnectionMonitorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all connection monitors for the specified Network Watcher.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (ConnectionMonitorsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
	if err != nil {
		return ConnectionMonitorsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionMonitorsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ConnectionMonitorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *ConnectionMonitorsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectionMonitorsClient) listHandleResponse(resp *http.Response) (ConnectionMonitorsListResponse, error) {
	result := ConnectionMonitorsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorListResult); err != nil {
		return ConnectionMonitorsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ConnectionMonitorsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginQuery - Query a snapshot of the most recent connection states.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) BeginQuery(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (ConnectionMonitorsQueryPollerResponse, error) {
	resp, err := client.query(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorsQueryPollerResponse{}, err
	}
	result := ConnectionMonitorsQueryPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectionMonitorsClient.Query", "location", resp, client.pl, client.queryHandleError)
	if err != nil {
		return ConnectionMonitorsQueryPollerResponse{}, err
	}
	result.Poller = &ConnectionMonitorsQueryPoller{
		pt: pt,
	}
	return result, nil
}

// Query - Query a snapshot of the most recent connection states.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) query(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (*http.Response, error) {
	req, err := client.queryCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.queryHandleError(resp)
	}
	return resp, nil
}

// queryCreateRequest creates the Query request.
func (client *ConnectionMonitorsClient) queryCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/query"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// queryHandleError handles the Query error response.
func (client *ConnectionMonitorsClient) queryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStart - Starts the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) BeginStart(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (ConnectionMonitorsStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorsStartPollerResponse{}, err
	}
	result := ConnectionMonitorsStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectionMonitorsClient.Start", "location", resp, client.pl, client.startHandleError)
	if err != nil {
		return ConnectionMonitorsStartPollerResponse{}, err
	}
	result.Poller = &ConnectionMonitorsStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Starts the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) start(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ConnectionMonitorsClient) startCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleError handles the Start error response.
func (client *ConnectionMonitorsClient) startHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginStop - Stops the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (ConnectionMonitorsStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return ConnectionMonitorsStopPollerResponse{}, err
	}
	result := ConnectionMonitorsStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ConnectionMonitorsClient.Stop", "location", resp, client.pl, client.stopHandleError)
	if err != nil {
		return ConnectionMonitorsStopPollerResponse{}, err
	}
	result.Poller = &ConnectionMonitorsStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stops the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ConnectionMonitorsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, options *ConnectionMonitorsBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *ConnectionMonitorsClient) stopHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Update tags of the specified connection monitor.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectionMonitorsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (ConnectionMonitorsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkWatcherName, connectionMonitorName, parameters, options)
	if err != nil {
		return ConnectionMonitorsUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionMonitorsUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionMonitorsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ConnectionMonitorsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, connectionMonitorName string, parameters TagsObject, options *ConnectionMonitorsUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if connectionMonitorName == "" {
		return nil, errors.New("parameter connectionMonitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionMonitorName}", url.PathEscape(connectionMonitorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ConnectionMonitorsClient) updateTagsHandleResponse(resp *http.Response) (ConnectionMonitorsUpdateTagsResponse, error) {
	result := ConnectionMonitorsUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionMonitorResult); err != nil {
		return ConnectionMonitorsUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ConnectionMonitorsClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
