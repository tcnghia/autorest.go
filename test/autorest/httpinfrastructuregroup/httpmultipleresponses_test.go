// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func newMultipleResponsesClient() *MultipleResponsesClient {
	return NewMultipleResponsesClient(NewDefaultConnection(nil))
}

// Get200Model201ModelDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200Model201ModelDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	switch x := result.Value.(type) {
	case MyException:
		if r := cmp.Diff(x.StatusCode, to.StringPtr("200")); r != "" {
			t.Fatal(r)
		}
	case B:
		if s := result.RawResponse.StatusCode; s != http.StatusCreated {
			t.Fatalf("unexpected status code %d", s)
		}
	default:
		t.Fatalf("unhandled response type %T", x)
	}
}

// Get200Model201ModelDefaultError201Valid - Send a 201 response with valid payload: {'statusCode': '201', 'textStatusCode': 'Created'}
func TestGet200Model201ModelDefaultError201Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError201Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.Value.(B)
	if !ok {
		t.Fatalf("unexpected response type %T", result)
	}
	if r := cmp.Diff(r, B{
		MyException: MyException{
			StatusCode: to.StringPtr("201"),
		},
		TextStatusCode: to.StringPtr("Created"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// Get200Model201ModelDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet200Model201ModelDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError400Valid(context.Background(), nil)
	var r *Error
	if !errors.As(err, &r) {
		t.Fatalf("unexpected error type %T", err)
	}
	if r := cmp.Diff(r, &Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	}, cmpopts.IgnoreUnexported(Error{})); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected nil result")
	}
}

// Get200Model204NoModelDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200Model204NoModelDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.MyException, MyException{
		StatusCode: to.StringPtr("200"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// Get200Model204NoModelDefaultError201Invalid - Send a 201 response with valid payload: {'statusCode': '201'}
func TestGet200Model204NoModelDefaultError201Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError201Invalid(context.Background(), nil)
	var r *Error
	if !errors.As(err, &r) {
		t.Fatal("unexpected error type")
	}
	if r := cmp.Diff(r, &Error{}, cmpopts.IgnoreUnexported(Error{})); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// Get200Model204NoModelDefaultError202None - Send a 202 response with no payload:
func TestGet200Model204NoModelDefaultError202None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError202None(context.Background(), nil)
	var r *Error
	if !errors.As(err, &r) {
		t.Fatal("unexpected error type")
	}
	if r := cmp.Diff(r, &Error{}, cmpopts.IgnoreUnexported(Error{})); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// Get200Model204NoModelDefaultError204Valid - Send a 204 response with no payload
func TestGet200Model204NoModelDefaultError204Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError204Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
	if !reflect.ValueOf(result.MyException).IsZero() {
		t.Fatal("expected zero-value MyException")
	}
}

// Get200Model204NoModelDefaultError400Valid - Send a 400 response with valid error payload: {'status': 400, 'message': 'client error'}
func TestGet200Model204NoModelDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError400Valid(context.Background(), nil)
	var r *Error
	if !errors.As(err, &r) {
		t.Fatal("unexpected error type")
	}
	if r := cmp.Diff(r, &Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	}, cmpopts.IgnoreUnexported(Error{})); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// Get200ModelA200Invalid - Send a 200 response with invalid payload {'statusCodeInvalid': '200'}
func TestGet200ModelA200Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// Get200ModelA200None - Send a 200 response with no payload, when a payload is expected - client should return a null object of thde type for model A
func TestGet200ModelA200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result.MyException).IsZero() {
		t.Fatal("expected zero-value MyException")
	}
}

// Get200ModelA200Valid - Send a 200 response with payload {'statusCode': '200'}
func TestGet200ModelA200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.MyException, MyException{
		StatusCode: to.StringPtr("200"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// Get200ModelA201ModelC404ModelDDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200ModelA201ModelC404ModelDDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.Value.(MyException)
	if !ok {
		t.Fatalf("unexpected result type %T", result)
	}
	if r := cmp.Diff(r, MyException{
		StatusCode: to.StringPtr("200"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// Get200ModelA201ModelC404ModelDDefaultError201Valid - Send a 200 response with valid payload: {'httpCode': '201'}
func TestGet200ModelA201ModelC404ModelDDefaultError201Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError201Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.Value.(C)
	if !ok {
		t.Fatalf("unexpected result type %T", result)
	}
	if r := cmp.Diff(r, C{
		HTTPCode: to.StringPtr("201"),
	}); r != "" {
		t.Fatal(r)
	}
}

// Get200ModelA201ModelC404ModelDDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet200ModelA201ModelC404ModelDDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError400Valid(context.Background(), nil)
	var r *Error
	if !errors.As(err, &r) {
		t.Fatal("unexpected error type")
	}
	if r := cmp.Diff(r, &Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	}, cmpopts.IgnoreUnexported(Error{})); r != "" {
		t.Fatal(r)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected nil result")
	}
}

// Get200ModelA201ModelC404ModelDDefaultError404Valid - Send a 200 response with valid payload: {'httpStatusCode': '404'}
func TestGet200ModelA201ModelC404ModelDDefaultError404Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError404Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.Value.(D)
	if !ok {
		t.Fatalf("unexpected result type %T", result)
	}
	if r := cmp.Diff(r, D{
		HTTPStatusCode: to.StringPtr("404"),
	}); r != "" {
		t.Fatal(r)
	}
}

// Get200ModelA202Valid - Send a 202 response with payload {'statusCode': '202'}
func TestGet200ModelA202Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.MyException, MyException{
		StatusCode: to.StringPtr("200"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// Get200ModelA400Invalid - Send a 200 response with invalid payload {'statusCodeInvalid': '400'}
func TestGet200ModelA400Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// Get200ModelA400None - Send a 400 response with no payload client should treat as an http error with no error model
func TestGet200ModelA400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA400None(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// Get200ModelA400Valid - Send a 200 response with payload {'statusCode': '400'}
func TestGet200ModelA400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA400Valid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected empty response")
	}
}

// Get202None204NoneDefaultError202None - Send a 202 response with no payload
func TestGet202None204NoneDefaultError202None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError202None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

// Get202None204NoneDefaultError204None - Send a 204 response with no payload
func TestGet202None204NoneDefaultError204None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError204None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

// Get202None204NoneDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet202None204NoneDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError400Valid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// Get202None204NoneDefaultNone202Invalid - Send a 202 response with an unexpected payload {'property': 'value'}
func TestGet202None204NoneDefaultNone202Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone202Invalid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusAccepted {
		t.Fatalf("unexpected status code %d", s)
	}
}

// Get202None204NoneDefaultNone204None - Send a 204 response with no payload
func TestGet202None204NoneDefaultNone204None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone204None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

// Get202None204NoneDefaultNone400Invalid - Send a 400 response with an unexpected payload {'property': 'value'}
func TestGet202None204NoneDefaultNone400Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone400Invalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// Get202None204NoneDefaultNone400None - Send a 400 response with no payload
func TestGet202None204NoneDefaultNone400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone400None(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultModelA200None - Send a 200 response with no payload
func TestGetDefaultModelA200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA200None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result.MyException).IsZero() {
		t.Fatal("expected zero-value MyException")
	}
}

// GetDefaultModelA200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGetDefaultModelA200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA200Valid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.MyException, MyException{
		StatusCode: to.StringPtr("200"),
	}, cmpopts.IgnoreUnexported(MyException{})); r != "" {
		t.Fatal(r)
	}
}

// GetDefaultModelA400None - Send a 400 response with no payload
func TestGetDefaultModelA400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA400None(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultModelA400Valid - Send a 400 response with valid payload: {'statusCode': '400'}
func TestGetDefaultModelA400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA400Valid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultNone200Invalid - Send a 200 response with invalid payload: {'statusCode': '200'}
func TestGetDefaultNone200Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// GetDefaultNone200None - Send a 200 response with no payload
func TestGetDefaultNone200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone200None(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// GetDefaultNone400Invalid - Send a 400 response with valid payload: {'statusCode': '400'}
func TestGetDefaultNone400Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone400Invalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultNone400None - Send a 400 response with no payload
func TestGetDefaultNone400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone400None(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("unexpected nil response")
	}
}
