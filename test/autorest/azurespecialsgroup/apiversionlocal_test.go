// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getAPIVersionLocalOperations(t *testing.T) azurespecialsgroup.APIVersionLocalOperations {
	client, err := azurespecialsgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.APIVersionLocalOperations()
}

// GetMethodLocalNull - Get method with api-version modeled in the method.  pass in api-version = null to succeed
func TestGetMethodLocalNull(t *testing.T) {
	client := getAPIVersionLocalOperations(t)
	result, err := client.GetMethodLocalNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetMethodLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetMethodLocalValid(t *testing.T) {
	client := getAPIVersionLocalOperations(t)
	result, err := client.GetMethodLocalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetPathLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetPathLocalValid(t *testing.T) {
	client := getAPIVersionLocalOperations(t)
	result, err := client.GetPathLocalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method.  pass in api-version = '2.0' to succeed
func TestGetSwaggerLocalValid(t *testing.T) {
	client := getAPIVersionLocalOperations(t)
	result, err := client.GetSwaggerLocalValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
