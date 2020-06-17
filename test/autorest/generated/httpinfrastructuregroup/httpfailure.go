// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// HTTPFailureOperations contains the methods for the HTTPFailure group.
type HTTPFailureOperations interface {
	// GetEmptyError - Get empty error form server
	GetEmptyError(ctx context.Context) (*BoolResponse, error)
	// GetNoModelEmpty - Get empty response from server
	GetNoModelEmpty(ctx context.Context) (*BoolResponse, error)
	// GetNoModelError - Get empty error form server
	GetNoModelError(ctx context.Context) (*BoolResponse, error)
}

// httpFailureOperations implements the HTTPFailureOperations interface.
type httpFailureOperations struct {
	*Client
}

// GetEmptyError - Get empty error form server
func (client *httpFailureOperations) GetEmptyError(ctx context.Context) (*BoolResponse, error) {
	req, err := client.getEmptyErrorCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getEmptyErrorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getEmptyErrorCreateRequest creates the GetEmptyError request.
func (client *httpFailureOperations) getEmptyErrorCreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/emptybody/error"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getEmptyErrorHandleResponse handles the GetEmptyError response.
func (client *httpFailureOperations) getEmptyErrorHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getEmptyErrorHandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getEmptyErrorHandleError handles the GetEmptyError error response.
func (client *httpFailureOperations) getEmptyErrorHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetNoModelEmpty - Get empty response from server
func (client *httpFailureOperations) GetNoModelEmpty(ctx context.Context) (*BoolResponse, error) {
	req, err := client.getNoModelEmptyCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNoModelEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNoModelEmptyCreateRequest creates the GetNoModelEmpty request.
func (client *httpFailureOperations) getNoModelEmptyCreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNoModelEmptyHandleResponse handles the GetNoModelEmpty response.
func (client *httpFailureOperations) getNoModelEmptyHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNoModelEmptyHandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getNoModelEmptyHandleError handles the GetNoModelEmpty error response.
func (client *httpFailureOperations) getNoModelEmptyHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// GetNoModelError - Get empty error form server
func (client *httpFailureOperations) GetNoModelError(ctx context.Context) (*BoolResponse, error) {
	req, err := client.getNoModelErrorCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNoModelErrorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNoModelErrorCreateRequest creates the GetNoModelError request.
func (client *httpFailureOperations) getNoModelErrorCreateRequest() (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/error"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNoModelErrorHandleResponse handles the GetNoModelError response.
func (client *httpFailureOperations) getNoModelErrorHandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getNoModelErrorHandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// getNoModelErrorHandleError handles the GetNoModelError error response.
func (client *httpFailureOperations) getNoModelErrorHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
