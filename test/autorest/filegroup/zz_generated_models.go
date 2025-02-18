//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package filegroup

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

// FilesGetEmptyFileOptions contains the optional parameters for the Files.GetEmptyFile method.
type FilesGetEmptyFileOptions struct {
	// placeholder for future optional parameters
}

// FilesGetFileLargeOptions contains the optional parameters for the Files.GetFileLarge method.
type FilesGetFileLargeOptions struct {
	// placeholder for future optional parameters
}

// FilesGetFileOptions contains the optional parameters for the Files.GetFile method.
type FilesGetFileOptions struct {
	// placeholder for future optional parameters
}
