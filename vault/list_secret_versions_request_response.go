// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package vault

import (
	"github.com/oracle/oci-go-sdk/v36/common"
	"net/http"
)

// ListSecretVersionsRequest wrapper for the ListSecretVersions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/ListSecretVersions.go.html to see an example of how to use ListSecretVersionsRequest.
type ListSecretVersionsRequest struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided. Time created is default ordered as descending. Display name is default ordered as ascending.
	SortBy ListSecretVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSecretVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecretVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecretVersionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecretVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSecretVersionsResponse wrapper for the ListSecretVersions operation
type ListSecretVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecretVersionSummary instances
	Items []SecretVersionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSecretVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecretVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecretVersionsSortByEnum Enum with underlying type: string
type ListSecretVersionsSortByEnum string

// Set of constants representing the allowable values for ListSecretVersionsSortByEnum
const (
	ListSecretVersionsSortByVersionNumber ListSecretVersionsSortByEnum = "VERSION_NUMBER"
)

var mappingListSecretVersionsSortBy = map[string]ListSecretVersionsSortByEnum{
	"VERSION_NUMBER": ListSecretVersionsSortByVersionNumber,
}

// GetListSecretVersionsSortByEnumValues Enumerates the set of values for ListSecretVersionsSortByEnum
func GetListSecretVersionsSortByEnumValues() []ListSecretVersionsSortByEnum {
	values := make([]ListSecretVersionsSortByEnum, 0)
	for _, v := range mappingListSecretVersionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSecretVersionsSortOrderEnum Enum with underlying type: string
type ListSecretVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListSecretVersionsSortOrderEnum
const (
	ListSecretVersionsSortOrderAsc  ListSecretVersionsSortOrderEnum = "ASC"
	ListSecretVersionsSortOrderDesc ListSecretVersionsSortOrderEnum = "DESC"
)

var mappingListSecretVersionsSortOrder = map[string]ListSecretVersionsSortOrderEnum{
	"ASC":  ListSecretVersionsSortOrderAsc,
	"DESC": ListSecretVersionsSortOrderDesc,
}

// GetListSecretVersionsSortOrderEnumValues Enumerates the set of values for ListSecretVersionsSortOrderEnum
func GetListSecretVersionsSortOrderEnumValues() []ListSecretVersionsSortOrderEnum {
	values := make([]ListSecretVersionsSortOrderEnum, 0)
	for _, v := range mappingListSecretVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
