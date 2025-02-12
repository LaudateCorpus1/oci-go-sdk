// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ExportDetails Input arguments for running a query synchronosly and streaming the results as soon as they become available.
type ExportDetails struct {

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Query to perform.
	QueryString *string `mandatory:"true" json:"queryString"`

	// Default subsystem to qualify fields with in the queryString if not specified.
	SubSystem SubSystemNameEnum `mandatory:"true" json:"subSystem"`

	// Flag to search all child compartments of the compartment Id specified in the compartmentId query parameter.
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// List of filters to be applied when the query executes. More than one filter per field is not permitted.
	ScopeFilters []ScopeFilter `mandatory:"false" json:"scopeFilters"`

	// Maximum number of results retrieved from data source.  Note a maximum value will be enforced; if the export results can be streamed, the maximum will be 50000000, otherwise 10000; that is, if not streamed, actualMaxTotalCountUsed = Math.min(maxTotalCount, 10000).
	//
	// Export will incrementally stream results depending on the queryString.
	// Some commands including head/tail are not compatible with streaming result delivery and therefore enforce a reduced limit on overall maxtotalcount.
	//  no sort command or sort by id, e.g. ' | sort id ' - is streaming compatible
	//  sort by time and id, e.g. ' | sort -time, id ' - is streaming compatible
	// all other cases, e.g. ' | sort -time, id, mtgtguid ' - is not streaming compatible due to the additional sort field
	MaxTotalCount *int `mandatory:"false" json:"maxTotalCount"`

	TimeFilter *TimeRange `mandatory:"false" json:"timeFilter"`

	// Amount of time, in seconds, allowed for a query to execute. If this time expires before the query is complete, any partial results will be returned.
	QueryTimeoutInSeconds *int `mandatory:"false" json:"queryTimeoutInSeconds"`

	// Include columns in response
	ShouldIncludeColumns *bool `mandatory:"false" json:"shouldIncludeColumns"`

	// Specifies the format for the returned results.
	OutputFormat ExportDetailsOutputFormatEnum `mandatory:"false" json:"outputFormat,omitempty"`

	// Localize results, including header columns, List-Of-Values and timestamp values.
	ShouldLocalize *bool `mandatory:"false" json:"shouldLocalize"`

	// Controls if query should ignore pre-calculated results if available and only use raw data.
	ShouldUseAcceleration *bool `mandatory:"false" json:"shouldUseAcceleration"`
}

func (m ExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubSystemNameEnum(string(m.SubSystem)); !ok && m.SubSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubSystem: %s. Supported values are: %s.", m.SubSystem, strings.Join(GetSubSystemNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExportDetailsOutputFormatEnum(string(m.OutputFormat)); !ok && m.OutputFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputFormat: %s. Supported values are: %s.", m.OutputFormat, strings.Join(GetExportDetailsOutputFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExportDetailsOutputFormatEnum Enum with underlying type: string
type ExportDetailsOutputFormatEnum string

// Set of constants representing the allowable values for ExportDetailsOutputFormatEnum
const (
	ExportDetailsOutputFormatCsv  ExportDetailsOutputFormatEnum = "CSV"
	ExportDetailsOutputFormatJson ExportDetailsOutputFormatEnum = "JSON"
)

var mappingExportDetailsOutputFormatEnum = map[string]ExportDetailsOutputFormatEnum{
	"CSV":  ExportDetailsOutputFormatCsv,
	"JSON": ExportDetailsOutputFormatJson,
}

// GetExportDetailsOutputFormatEnumValues Enumerates the set of values for ExportDetailsOutputFormatEnum
func GetExportDetailsOutputFormatEnumValues() []ExportDetailsOutputFormatEnum {
	values := make([]ExportDetailsOutputFormatEnum, 0)
	for _, v := range mappingExportDetailsOutputFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetExportDetailsOutputFormatEnumStringValues Enumerates the set of values in String for ExportDetailsOutputFormatEnum
func GetExportDetailsOutputFormatEnumStringValues() []string {
	return []string{
		"CSV",
		"JSON",
	}
}

// GetMappingExportDetailsOutputFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExportDetailsOutputFormatEnum(val string) (ExportDetailsOutputFormatEnum, bool) {
	mappingExportDetailsOutputFormatEnumIgnoreCase := make(map[string]ExportDetailsOutputFormatEnum)
	for k, v := range mappingExportDetailsOutputFormatEnum {
		mappingExportDetailsOutputFormatEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExportDetailsOutputFormatEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
