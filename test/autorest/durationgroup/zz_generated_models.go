// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package durationgroup

// DurationGetInvalidOptions contains the optional parameters for the Duration.GetInvalid method.
type DurationGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// DurationGetNullOptions contains the optional parameters for the Duration.GetNull method.
type DurationGetNullOptions struct {
	// placeholder for future optional parameters
}

// DurationGetPositiveDurationOptions contains the optional parameters for the Duration.GetPositiveDuration method.
type DurationGetPositiveDurationOptions struct {
	// placeholder for future optional parameters
}

// DurationPutPositiveDurationOptions contains the optional parameters for the Duration.PutPositiveDuration method.
type DurationPutPositiveDurationOptions struct {
	// placeholder for future optional parameters
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}
