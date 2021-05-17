// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurereportgroup

import "net/http"

// MapOfInt32Response is the response envelope for operations that return a map[string]*int32 type.
type MapOfInt32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Dictionary of <integer>
	Value map[string]*int32
}
