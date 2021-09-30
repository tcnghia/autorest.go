//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

var scopes = []string{"https://dev.azuresynapse.net/.default"}

// connectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type connectionOptions struct {
	// Transport sets the transport for making HTTP requests.
	Transport policy.Transporter
	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []policy.Policy
}

type connection struct {
	u string
	p runtime.Pipeline
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newConnection(endpoint string, cred azcore.Credential, options *connectionOptions) *connection {
	if options == nil {
		options = &connectionOptions{}
	}
	policies := []policy.Policy{}
	if !options.Telemetry.Disabled {
		policies = append(policies, runtime.NewTelemetryPolicy(module, version, &options.Telemetry))
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, runtime.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, cred.NewAuthenticationPolicy(runtime.AuthenticationOptions{TokenRequest: policy.TokenRequestOptions{Scopes: scopes}}))
	policies = append(policies, runtime.NewLogPolicy(&options.Logging))
	return &connection{u: endpoint, p: runtime.NewPipeline(options.Transport, policies...)}
}

// Endpoint returns the connection's endpoint.
func (c *connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *connection) Pipeline() runtime.Pipeline {
	return c.p
}
