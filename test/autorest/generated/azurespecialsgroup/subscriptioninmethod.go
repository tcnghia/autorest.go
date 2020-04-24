// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInMethodOperations contains the methods for the SubscriptionInMethod group.
type SubscriptionInMethodOperations interface {
	// PostMethodLocalNull - POST method with subscriptionId modeled in the method.  pass in subscription id = null, client-side validation should prevent you from making this call
	PostMethodLocalNull(ctx context.Context, subscriptionId string) (*http.Response, error)
	// PostMethodLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
	PostMethodLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error)
	// PostPathLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
	PostPathLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error)
	// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
	PostSwaggerLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error)
}

// subscriptionInMethodOperations implements the SubscriptionInMethodOperations interface.
type subscriptionInMethodOperations struct {
	*Client
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method.  pass in subscription id = null, client-side validation should prevent you from making this call
func (client *subscriptionInMethodOperations) PostMethodLocalNull(ctx context.Context, subscriptionId string) (*http.Response, error) {
	req, err := client.postMethodLocalNullCreateRequest(subscriptionId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMethodLocalNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMethodLocalNullCreateRequest creates the PostMethodLocalNull request.
func (client *subscriptionInMethodOperations) postMethodLocalNullCreateRequest(subscriptionId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postMethodLocalNullHandleResponse handles the PostMethodLocalNull response.
func (client *subscriptionInMethodOperations) postMethodLocalNullHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *subscriptionInMethodOperations) PostMethodLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error) {
	req, err := client.postMethodLocalValidCreateRequest(subscriptionId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postMethodLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postMethodLocalValidCreateRequest creates the PostMethodLocalValid request.
func (client *subscriptionInMethodOperations) postMethodLocalValidCreateRequest(subscriptionId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postMethodLocalValidHandleResponse handles the PostMethodLocalValid response.
func (client *subscriptionInMethodOperations) postMethodLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *subscriptionInMethodOperations) PostPathLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error) {
	req, err := client.postPathLocalValidCreateRequest(subscriptionId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postPathLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postPathLocalValidCreateRequest creates the PostPathLocalValid request.
func (client *subscriptionInMethodOperations) postPathLocalValidCreateRequest(subscriptionId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postPathLocalValidHandleResponse handles the PostPathLocalValid response.
func (client *subscriptionInMethodOperations) postPathLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method.  pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *subscriptionInMethodOperations) PostSwaggerLocalValid(ctx context.Context, subscriptionId string) (*http.Response, error) {
	req, err := client.postSwaggerLocalValidCreateRequest(subscriptionId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postSwaggerLocalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// postSwaggerLocalValidCreateRequest creates the PostSwaggerLocalValid request.
func (client *subscriptionInMethodOperations) postSwaggerLocalValidCreateRequest(subscriptionId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, nil
}

// postSwaggerLocalValidHandleResponse handles the PostSwaggerLocalValid response.
func (client *subscriptionInMethodOperations) postSwaggerLocalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
