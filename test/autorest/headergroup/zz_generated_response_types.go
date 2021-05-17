// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headergroup

import (
	"net/http"
	"time"
)

// HeaderResponseBoolResponse contains the response from method Header.ResponseBool.
type HeaderResponseBoolResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *bool
}

// HeaderResponseByteResponse contains the response from method Header.ResponseByte.
type HeaderResponseByteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value []byte
}

// HeaderResponseDateResponse contains the response from method Header.ResponseDate.
type HeaderResponseDateResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderResponseDatetimeRFC1123Response contains the response from method Header.ResponseDatetimeRFC1123.
type HeaderResponseDatetimeRFC1123Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderResponseDatetimeResponse contains the response from method Header.ResponseDatetime.
type HeaderResponseDatetimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *time.Time
}

// HeaderResponseDoubleResponse contains the response from method Header.ResponseDouble.
type HeaderResponseDoubleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *float64
}

// HeaderResponseDurationResponse contains the response from method Header.ResponseDuration.
type HeaderResponseDurationResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *string
}

// HeaderResponseEnumResponse contains the response from method Header.ResponseEnum.
type HeaderResponseEnumResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *GreyscaleColors
}

// HeaderResponseExistingKeyResponse contains the response from method Header.ResponseExistingKey.
type HeaderResponseExistingKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// UserAgent contains the information returned from the User-Agent header response.
	UserAgent *string
}

// HeaderResponseFloatResponse contains the response from method Header.ResponseFloat.
type HeaderResponseFloatResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *float32
}

// HeaderResponseIntegerResponse contains the response from method Header.ResponseInteger.
type HeaderResponseIntegerResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *int32
}

// HeaderResponseLongResponse contains the response from method Header.ResponseLong.
type HeaderResponseLongResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *int64
}

// HeaderResponseProtectedKeyResponse contains the response from method Header.ResponseProtectedKey.
type HeaderResponseProtectedKeyResponse struct {
	// ContentType contains the information returned from the Content-Type header response.
	ContentType *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderResponseStringResponse contains the response from method Header.ResponseString.
type HeaderResponseStringResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Value contains the information returned from the value header response.
	Value *string
}
