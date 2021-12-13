// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
)

// ListProblemHistoriesRequest wrapper for the ListProblemHistories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListProblemHistories.go.html to see an example of how to use ListProblemHistoriesRequest.
type ListProblemHistoriesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCId of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListProblemHistoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListProblemHistoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProblemHistoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProblemHistoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProblemHistoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProblemHistoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProblemHistoriesResponse wrapper for the ListProblemHistories operation
type ListProblemHistoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProblemHistoryCollection instances
	ProblemHistoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProblemHistoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProblemHistoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProblemHistoriesSortOrderEnum Enum with underlying type: string
type ListProblemHistoriesSortOrderEnum string

// Set of constants representing the allowable values for ListProblemHistoriesSortOrderEnum
const (
	ListProblemHistoriesSortOrderAsc  ListProblemHistoriesSortOrderEnum = "ASC"
	ListProblemHistoriesSortOrderDesc ListProblemHistoriesSortOrderEnum = "DESC"
)

var mappingListProblemHistoriesSortOrder = map[string]ListProblemHistoriesSortOrderEnum{
	"ASC":  ListProblemHistoriesSortOrderAsc,
	"DESC": ListProblemHistoriesSortOrderDesc,
}

// GetListProblemHistoriesSortOrderEnumValues Enumerates the set of values for ListProblemHistoriesSortOrderEnum
func GetListProblemHistoriesSortOrderEnumValues() []ListProblemHistoriesSortOrderEnum {
	values := make([]ListProblemHistoriesSortOrderEnum, 0)
	for _, v := range mappingListProblemHistoriesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProblemHistoriesSortByEnum Enum with underlying type: string
type ListProblemHistoriesSortByEnum string

// Set of constants representing the allowable values for ListProblemHistoriesSortByEnum
const (
	ListProblemHistoriesSortByTimecreated ListProblemHistoriesSortByEnum = "timeCreated"
)

var mappingListProblemHistoriesSortBy = map[string]ListProblemHistoriesSortByEnum{
	"timeCreated": ListProblemHistoriesSortByTimecreated,
}

// GetListProblemHistoriesSortByEnumValues Enumerates the set of values for ListProblemHistoriesSortByEnum
func GetListProblemHistoriesSortByEnumValues() []ListProblemHistoriesSortByEnum {
	values := make([]ListProblemHistoriesSortByEnum, 0)
	for _, v := range mappingListProblemHistoriesSortBy {
		values = append(values, v)
	}
	return values
}
