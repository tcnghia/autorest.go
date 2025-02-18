//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paramgroupinggroup

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

// FirstParameterGroup contains a group of parameters for the ParameterGrouping.PostMultiParamGroups method.
type FirstParameterGroup struct {
	HeaderOne *string
	// Query parameter with default
	QueryOne *int32
}

// ParameterGroupingPostMultiParamGroupsOptions contains the optional parameters for the ParameterGrouping.PostMultiParamGroups method.
type ParameterGroupingPostMultiParamGroupsOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingPostMultiParamGroupsSecondParamGroup contains a group of parameters for the ParameterGrouping.PostMultiParamGroups method.
type ParameterGroupingPostMultiParamGroupsSecondParamGroup struct {
	HeaderTwo *string
	// Query parameter with default
	QueryTwo *int32
}

// ParameterGroupingPostOptionalOptions contains the optional parameters for the ParameterGrouping.PostOptional method.
type ParameterGroupingPostOptionalOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingPostOptionalParameters contains a group of parameters for the ParameterGrouping.PostOptional method.
type ParameterGroupingPostOptionalParameters struct {
	CustomHeader *string
	// Query parameter with default
	Query *int32
}

// ParameterGroupingPostRequiredOptions contains the optional parameters for the ParameterGrouping.PostRequired method.
type ParameterGroupingPostRequiredOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingPostRequiredParameters contains a group of parameters for the ParameterGrouping.PostRequired method.
type ParameterGroupingPostRequiredParameters struct {
	Body         int32
	CustomHeader *string
	// Path parameter
	Path string
	// Query parameter with default
	Query *int32
}

// ParameterGroupingPostReservedWordsOptions contains the optional parameters for the ParameterGrouping.PostReservedWords method.
type ParameterGroupingPostReservedWordsOptions struct {
	// placeholder for future optional parameters
}

// ParameterGroupingPostReservedWordsParameters contains a group of parameters for the ParameterGrouping.PostReservedWords method.
type ParameterGroupingPostReservedWordsParameters struct {
	// 'accept' is a reserved word. Pass in 'yes' to pass.
	Accept *string
	// 'from' is a reserved word. Pass in 'bob' to pass.
	From *string
}

// ParameterGroupingPostSharedParameterGroupObjectOptions contains the optional parameters for the ParameterGrouping.PostSharedParameterGroupObject method.
type ParameterGroupingPostSharedParameterGroupObjectOptions struct {
	// placeholder for future optional parameters
}
