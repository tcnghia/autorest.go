// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package validationgroup

import "net/http"

// ProductResponse is the response envelope for operations that return a Product type.
type ProductResponse struct {
	// The product documentation.
	Product *Product

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
