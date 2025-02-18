//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPServerFailureClient contains the methods for the HTTPServerFailure group.
// Don't use this type directly, use NewHTTPServerFailureClient() instead.
type HTTPServerFailureClient struct {
	con *Connection
}

// NewHTTPServerFailureClient creates a new instance of HTTPServerFailureClient with the specified values.
func NewHTTPServerFailureClient(con *Connection) *HTTPServerFailureClient {
	return &HTTPServerFailureClient{con: con}
}

// Delete505 - Return 505 status code - should be represented in the client as an error
// If the operation fails it returns the *Error error type.
func (client *HTTPServerFailureClient) Delete505(ctx context.Context, options *HTTPServerFailureDelete505Options) (HTTPServerFailureDelete505Response, error) {
	req, err := client.delete505CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureDelete505Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPServerFailureDelete505Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailureDelete505Response{}, client.delete505HandleError(resp)
	}
	return HTTPServerFailureDelete505Response{RawResponse: resp}, nil
}

// delete505CreateRequest creates the Delete505 request.
func (client *HTTPServerFailureClient) delete505CreateRequest(ctx context.Context, options *HTTPServerFailureDelete505Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// delete505HandleError handles the Delete505 error response.
func (client *HTTPServerFailureClient) delete505HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get501 - Return 501 status code - should be represented in the client as an error
// If the operation fails it returns the *Error error type.
func (client *HTTPServerFailureClient) Get501(ctx context.Context, options *HTTPServerFailureGet501Options) (HTTPServerFailureGet501Response, error) {
	req, err := client.get501CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureGet501Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPServerFailureGet501Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailureGet501Response{}, client.get501HandleError(resp)
	}
	return HTTPServerFailureGet501Response{RawResponse: resp}, nil
}

// get501CreateRequest creates the Get501 request.
func (client *HTTPServerFailureClient) get501CreateRequest(ctx context.Context, options *HTTPServerFailureGet501Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// get501HandleError handles the Get501 error response.
func (client *HTTPServerFailureClient) get501HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Head501 - Return 501 status code - should be represented in the client as an error
// If the operation fails it returns the *Error error type.
func (client *HTTPServerFailureClient) Head501(ctx context.Context, options *HTTPServerFailureHead501Options) (HTTPServerFailureHead501Response, error) {
	req, err := client.head501CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailureHead501Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPServerFailureHead501Response{}, err
	}
	result := HTTPServerFailureHead501Response{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head501CreateRequest creates the Head501 request.
func (client *HTTPServerFailureClient) head501CreateRequest(ctx context.Context, options *HTTPServerFailureHead501Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Post505 - Return 505 status code - should be represented in the client as an error
// If the operation fails it returns the *Error error type.
func (client *HTTPServerFailureClient) Post505(ctx context.Context, options *HTTPServerFailurePost505Options) (HTTPServerFailurePost505Response, error) {
	req, err := client.post505CreateRequest(ctx, options)
	if err != nil {
		return HTTPServerFailurePost505Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPServerFailurePost505Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPServerFailurePost505Response{}, client.post505HandleError(resp)
	}
	return HTTPServerFailurePost505Response{RawResponse: resp}, nil
}

// post505CreateRequest creates the Post505 request.
func (client *HTTPServerFailureClient) post505CreateRequest(ctx context.Context, options *HTTPServerFailurePost505Options) (*policy.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, true)
}

// post505HandleError handles the Post505 error response.
func (client *HTTPServerFailureClient) post505HandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
