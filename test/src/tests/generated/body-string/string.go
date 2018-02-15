package stringgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// StringClient is the test Infrastructure for AutoRest Swagger BAT
type StringClient struct {
	ManagementClient
}

// NewStringClient creates an instance of the StringClient client.
func NewStringClient(p pipeline.Pipeline) StringClient {
	return StringClient{NewManagementClient(p)}
}

// GetBase64Encoded get value that is base64 encoded
func (client StringClient) GetBase64Encoded(ctx context.Context) (*GetBase64EncodedResponse, error) {
	req, err := client.getBase64EncodedPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getBase64EncodedResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetBase64EncodedResponse), err
}

// getBase64EncodedPreparer prepares the GetBase64Encoded request.
func (client StringClient) getBase64EncodedPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/base64Encoding"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getBase64EncodedResponder handles the response to the GetBase64Encoded request.
func (client StringClient) getBase64EncodedResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetBase64EncodedResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetBase64URLEncoded get value that is base64url encoded
func (client StringClient) GetBase64URLEncoded(ctx context.Context) (*GetBase64URLEncodedResponse, error) {
	req, err := client.getBase64URLEncodedPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getBase64URLEncodedResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetBase64URLEncodedResponse), err
}

// getBase64URLEncodedPreparer prepares the GetBase64URLEncoded request.
func (client StringClient) getBase64URLEncodedPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/base64UrlEncoding"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getBase64URLEncodedResponder handles the response to the GetBase64URLEncoded request.
func (client StringClient) getBase64URLEncodedResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetBase64URLEncodedResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetEmpty get empty string value value ''
func (client StringClient) GetEmpty(ctx context.Context) (*GetEmptyResponse, error) {
	req, err := client.getEmptyPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getEmptyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetEmptyResponse), err
}

// getEmptyPreparer prepares the GetEmpty request.
func (client StringClient) getEmptyPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/empty"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getEmptyResponder handles the response to the GetEmpty request.
func (client StringClient) getEmptyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetEmptyResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetMbcs get mbcs string value '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
func (client StringClient) GetMbcs(ctx context.Context) (*GetMbcsResponse, error) {
	req, err := client.getMbcsPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getMbcsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetMbcsResponse), err
}

// getMbcsPreparer prepares the GetMbcs request.
func (client StringClient) getMbcsPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/mbcs"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getMbcsResponder handles the response to the GetMbcs request.
func (client StringClient) getMbcsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetMbcsResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetNotProvided get String value when no string value is sent in response payload
func (client StringClient) GetNotProvided(ctx context.Context) (*GetNotProvidedResponse, error) {
	req, err := client.getNotProvidedPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNotProvidedResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetNotProvidedResponse), err
}

// getNotProvidedPreparer prepares the GetNotProvided request.
func (client StringClient) getNotProvidedPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/notProvided"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getNotProvidedResponder handles the response to the GetNotProvided request.
func (client StringClient) getNotProvidedResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetNotProvidedResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetNull get null string value value
func (client StringClient) GetNull(ctx context.Context) (*GetNullResponse, error) {
	req, err := client.getNullPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNullResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetNullResponse), err
}

// getNullPreparer prepares the GetNull request.
func (client StringClient) getNullPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/null"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getNullResponder handles the response to the GetNull request.
func (client StringClient) getNullResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetNullResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetNullBase64URLEncoded get null value that is expected to be base64url encoded
func (client StringClient) GetNullBase64URLEncoded(ctx context.Context) (*GetNullBase64URLEncodedResponse, error) {
	req, err := client.getNullBase64URLEncodedPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getNullBase64URLEncodedResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetNullBase64URLEncodedResponse), err
}

// getNullBase64URLEncodedPreparer prepares the GetNullBase64URLEncoded request.
func (client StringClient) getNullBase64URLEncodedPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/nullBase64UrlEncoding"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getNullBase64URLEncodedResponder handles the response to the GetNullBase64URLEncoded request.
func (client StringClient) getNullBase64URLEncodedResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetNullBase64URLEncodedResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetWhitespace get string value with leading and trailing whitespace '<tab><space><space>Now is the time for all good
// men to come to the aid of their country<tab><space><space>'
func (client StringClient) GetWhitespace(ctx context.Context) (*GetWhitespaceResponse, error) {
	req, err := client.getWhitespacePreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getWhitespaceResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetWhitespaceResponse), err
}

// getWhitespacePreparer prepares the GetWhitespace request.
func (client StringClient) getWhitespacePreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/whitespace"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getWhitespaceResponder handles the response to the GetWhitespace request.
func (client StringClient) getWhitespaceResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetWhitespaceResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// PutBase64URLEncoded put value that is base64url encoded
//
func (client StringClient) PutBase64URLEncoded(ctx context.Context, stringBody string) (*http.Response, error) {
	req, err := client.putBase64URLEncodedPreparer(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putBase64URLEncodedResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// putBase64URLEncodedPreparer prepares the PutBase64URLEncoded request.
func (client StringClient) putBase64URLEncodedPreparer(stringBody string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/base64UrlEncoding"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(stringBody)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putBase64URLEncodedResponder handles the response to the PutBase64URLEncoded request.
func (client StringClient) putBase64URLEncodedResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// PutEmpty set string value empty ''
//
// stringBody is
func (client StringClient) PutEmpty(ctx context.Context, stringBody string) (*http.Response, error) {
	req, err := client.putEmptyPreparer(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putEmptyResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// putEmptyPreparer prepares the PutEmpty request.
func (client StringClient) putEmptyPreparer(stringBody string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/empty"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(stringBody)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putEmptyResponder handles the response to the PutEmpty request.
func (client StringClient) putEmptyResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// PutMbcs set string value mbcs '啊齄丂狛狜隣郎隣兀﨩ˊ〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€'
//
// stringBody is
func (client StringClient) PutMbcs(ctx context.Context, stringBody string) (*http.Response, error) {
	req, err := client.putMbcsPreparer(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putMbcsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// putMbcsPreparer prepares the PutMbcs request.
func (client StringClient) putMbcsPreparer(stringBody string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/mbcs"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(stringBody)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putMbcsResponder handles the response to the PutMbcs request.
func (client StringClient) putMbcsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// PutNull set string value null
//
// stringBody is
func (client StringClient) PutNull(ctx context.Context, stringBody string) (*http.Response, error) {
	req, err := client.putNullPreparer(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putNullResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// putNullPreparer prepares the PutNull request.
func (client StringClient) putNullPreparer(stringBody string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/null"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(stringBody)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putNullResponder handles the response to the PutNull request.
func (client StringClient) putNullResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// PutWhitespace set String value with leading and trailing whitespace '<tab><space><space>Now is the time for all good
// men to come to the aid of their country<tab><space><space>'
//
// stringBody is
func (client StringClient) PutWhitespace(ctx context.Context, stringBody string) (*http.Response, error) {
	req, err := client.putWhitespacePreparer(stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.putWhitespaceResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// putWhitespacePreparer prepares the PutWhitespace request.
func (client StringClient) putWhitespacePreparer(stringBody string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/string/whitespace"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(stringBody)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// putWhitespaceResponder handles the response to the PutWhitespace request.
func (client StringClient) putWhitespaceResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}
