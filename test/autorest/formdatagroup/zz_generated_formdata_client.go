//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package formdatagroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
)

// FormdataClient contains the methods for the Formdata group.
// Don't use this type directly, use NewFormdataClient() instead.
type FormdataClient struct {
	con *Connection
}

// NewFormdataClient creates a new instance of FormdataClient with the specified values.
func NewFormdataClient(con *Connection) *FormdataClient {
	return &FormdataClient{con: con}
}

// UploadFile - Upload file
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFile(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *FormdataUploadFileOptions) (FormdataUploadFileResponse, error) {
	req, err := client.uploadFileCreateRequest(ctx, fileContent, fileName, options)
	if err != nil {
		return FormdataUploadFileResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FormdataUploadFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataUploadFileResponse{}, client.uploadFileHandleError(resp)
	}
	return FormdataUploadFileResponse{RawResponse: resp}, nil
}

// uploadFileCreateRequest creates the UploadFile request.
func (client *FormdataClient) uploadFileCreateRequest(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *FormdataUploadFileOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	if err := runtime.SetMultipartFormData(req, map[string]interface{}{
		"fileContent": fileContent,
		"fileName":    fileName,
	}); err != nil {
		return nil, err
	}
	return req, nil
}

// uploadFileHandleError handles the UploadFile error response.
func (client *FormdataClient) uploadFileHandleError(resp *http.Response) error {
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

// UploadFileViaBody - Upload file
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFileViaBody(ctx context.Context, fileContent io.ReadSeekCloser, options *FormdataUploadFileViaBodyOptions) (FormdataUploadFileViaBodyResponse, error) {
	req, err := client.uploadFileViaBodyCreateRequest(ctx, fileContent, options)
	if err != nil {
		return FormdataUploadFileViaBodyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FormdataUploadFileViaBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataUploadFileViaBodyResponse{}, client.uploadFileViaBodyHandleError(resp)
	}
	return FormdataUploadFileViaBodyResponse{RawResponse: resp}, nil
}

// uploadFileViaBodyCreateRequest creates the UploadFileViaBody request.
func (client *FormdataClient) uploadFileViaBodyCreateRequest(ctx context.Context, fileContent io.ReadSeekCloser, options *FormdataUploadFileViaBodyOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	return req, req.SetBody(fileContent, "application/octet-stream")
}

// uploadFileViaBodyHandleError handles the UploadFileViaBody error response.
func (client *FormdataClient) uploadFileViaBodyHandleError(resp *http.Response) error {
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

// UploadFiles - Upload multiple files
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFiles(ctx context.Context, fileContent []io.ReadSeekCloser, options *FormdataUploadFilesOptions) (FormdataUploadFilesResponse, error) {
	req, err := client.uploadFilesCreateRequest(ctx, fileContent, options)
	if err != nil {
		return FormdataUploadFilesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FormdataUploadFilesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataUploadFilesResponse{}, client.uploadFilesHandleError(resp)
	}
	return FormdataUploadFilesResponse{RawResponse: resp}, nil
}

// uploadFilesCreateRequest creates the UploadFiles request.
func (client *FormdataClient) uploadFilesCreateRequest(ctx context.Context, fileContent []io.ReadSeekCloser, options *FormdataUploadFilesOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfiles"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	if err := runtime.SetMultipartFormData(req, map[string]interface{}{
		"fileContent": fileContent,
	}); err != nil {
		return nil, err
	}
	return req, nil
}

// uploadFilesHandleError handles the UploadFiles error response.
func (client *FormdataClient) uploadFilesHandleError(resp *http.Response) error {
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
