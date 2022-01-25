// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// DetachDevicesFromTransferPackageRequest wrapper for the DetachDevicesFromTransferPackage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/DetachDevicesFromTransferPackage.go.html to see an example of how to use DetachDevicesFromTransferPackageRequest.
type DetachDevicesFromTransferPackageRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// Label of the Transfer Package
	TransferPackageLabel *string `mandatory:"true" contributesTo:"path" name:"transferPackageLabel"`

	// Labels of Transfer Devices to detach
	DetachDevicesDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DetachDevicesFromTransferPackageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DetachDevicesFromTransferPackageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DetachDevicesFromTransferPackageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DetachDevicesFromTransferPackageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DetachDevicesFromTransferPackageResponse wrapper for the DetachDevicesFromTransferPackage operation
type DetachDevicesFromTransferPackageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response DetachDevicesFromTransferPackageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DetachDevicesFromTransferPackageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
