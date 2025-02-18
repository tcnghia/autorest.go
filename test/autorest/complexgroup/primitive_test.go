// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
)

func newPrimitiveClient() *PrimitiveClient {
	return NewPrimitiveClient(NewDefaultConnection(nil))
}

func TestPrimitiveGetInt(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetInt(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetInt: %v", err)
	}
	if r := cmp.Diff(result.IntWrapper, IntWrapper{Field1: to.Int32Ptr(-1), Field2: to.Int32Ptr(2)}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutInt(t *testing.T) {
	client := newPrimitiveClient()
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), IntWrapper{Field1: &a, Field2: &b}, nil)
	if err != nil {
		t.Fatalf("PutInt: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetLong(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetLong(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetLong: %v", err)
	}
	if r := cmp.Diff(result.LongWrapper, LongWrapper{
		Field1: to.Int64Ptr(1099511627775),
		Field2: to.Int64Ptr(-999511627788),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutLong(t *testing.T) {
	client := newPrimitiveClient()
	a, b := int64(1099511627775), int64(-999511627788)
	result, err := client.PutLong(context.Background(), LongWrapper{Field1: &a, Field2: &b}, nil)
	if err != nil {
		t.Fatalf("PutLong: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetFloat(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetFloat(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetFloat: %v", err)
	}
	if r := cmp.Diff(result.FloatWrapper, FloatWrapper{
		Field1: to.Float32Ptr(1.05),
		Field2: to.Float32Ptr(-0.003),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutFloat(t *testing.T) {
	client := newPrimitiveClient()
	a, b := float32(1.05), float32(-0.003)
	result, err := client.PutFloat(context.Background(), FloatWrapper{Field1: &a, Field2: &b}, nil)
	if err != nil {
		t.Fatalf("PutFloat: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetDouble(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDouble(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetDouble: %v", err)
	}
	if r := cmp.Diff(result.DoubleWrapper, DoubleWrapper{
		Field1: to.Float64Ptr(3e-100),
		Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: to.Float64Ptr(-0.000000000000000000000000000000000000000000000000000000005),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDouble(t *testing.T) {
	client := newPrimitiveClient()
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	result, err := client.PutDouble(context.Background(), DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b}, nil)
	if err != nil {
		t.Fatalf("PutDouble: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetBool(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetBool(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetBool: %v", err)
	}
	if r := cmp.Diff(result.BooleanWrapper, BooleanWrapper{
		FieldFalse: to.BoolPtr(false),
		FieldTrue:  to.BoolPtr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitiveGetByte(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetByte(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetByte: %v", err)
	}
	if r := cmp.Diff(result.ByteWrapper, ByteWrapper{Field: []byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutBool(t *testing.T) {
	client := newPrimitiveClient()
	a, b := true, false
	result, err := client.PutBool(context.Background(), BooleanWrapper{FieldTrue: &a, FieldFalse: &b}, nil)
	if err != nil {
		t.Fatalf("PutBool: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestByteWrapperJSONNull(t *testing.T) {
	bw := ByteWrapper{}
	b, err := json.Marshal(bw)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "{}" {
		t.Fatalf("unexpected value %s", string(b))
	}
	bw.Field = azcore.NullValue([]byte{}).([]byte)
	b, err = json.Marshal(bw)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != `{"field":null}` {
		t.Fatalf("unexpected value %s", string(b))
	}
}

func TestPrimitivePutByte(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.PutByte(context.Background(), ByteWrapper{Field: []byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}}, nil)
	if err != nil {
		t.Fatalf("PutByte: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetString(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetString(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetString: %v", err)
	}
	if r := cmp.Diff(result.StringWrapper, StringWrapper{
		Empty: to.StringPtr(""),
		Field: to.StringPtr("goodrequest"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutString(t *testing.T) {
	client := newPrimitiveClient()
	var c *string
	a, b, c := "goodrequest", "", nil
	result, err := client.PutString(context.Background(), StringWrapper{Field: &a, Empty: &b, Null: c}, nil)
	if err != nil {
		t.Fatalf("PutString: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetDate(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDate(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetDate: %v", err)
	}
	a, err := time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		t.Fatalf("Unable to parse date string: %v", err)
	}
	b, err := time.Parse("2006-01-02", "2016-02-29")
	if err != nil {
		t.Fatalf("Unable to parse leap year date string: %v", err)
	}
	dw := DateWrapper{Field: &a, Leap: &b}
	if r := cmp.Diff(result.DateWrapper, dw); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDate(t *testing.T) {
	client := newPrimitiveClient()
	a, err := time.Parse("2006-01-02", "0001-01-01")
	if err != nil {
		t.Fatalf("Unable to parse date string: %v", err)
	}
	b, err := time.Parse("2006-01-02", "2016-02-29")
	if err != nil {
		t.Fatalf("Unable to parse leap year date string: %v", err)
	}
	resp, err := client.PutDate(context.Background(), DateWrapper{Field: &a, Leap: &b}, nil)
	if err != nil {
		t.Fatalf("PutDate: %v", err)
	}
	if s := resp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetDuration(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDuration(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetDuration: %v", err)
	}
	if r := cmp.Diff(result.DurationWrapper, DurationWrapper{
		Field: to.StringPtr("P123DT22H14M12.011S"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDuration(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.PutDuration(context.Background(), DurationWrapper{Field: to.StringPtr("P123DT22H14M12.011S")}, nil)
	if err != nil {
		t.Fatalf("PutDuration: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitiveGetDateTime(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDateTime(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetDateTime: %v", err)
	}
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	if r := cmp.Diff(result.DatetimeWrapper, DatetimeWrapper{
		Field: &f,
		Now:   &n,
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitiveGetDateTimeRFC1123(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDateTimeRFC1123(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetDateTimeRFC1123: %v", err)
	}
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	if r := cmp.Diff(result.Datetimerfc1123Wrapper, Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDateTime(t *testing.T) {
	client := newPrimitiveClient()
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	result, err := client.PutDateTime(context.Background(), DatetimeWrapper{
		Field: &f,
		Now:   &n,
	}, nil)
	if err != nil {
		t.Fatalf("PutDateTime: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestPrimitivePutDateTimeRFC1123(t *testing.T) {
	client := newPrimitiveClient()
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	result, err := client.PutDateTimeRFC1123(context.Background(), Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	}, nil)
	if err != nil {
		t.Fatalf("PutDateTimeRFC1123: %v", err)
	}
	if s := result.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
