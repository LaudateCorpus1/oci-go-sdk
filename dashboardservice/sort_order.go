// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dashboards API
//
// Use the Oracle Cloud Infrastructure Dashboards service API to manage dashboards in the Console.
// Dashboards provide an organized and customizable view of resources and their metrics in the Console.
// For more information, see Dashboards (https://docs.cloud.oracle.com/Content/Dashboards/home.htm).
// **Important:** Resources for the Dashboards service are created in the tenacy's home region.
// Although it is possible to create dashboard and dashboard group resources in regions other than the home region,
// you won't be able to view those resources in the Console.
// Therefore, creating resources outside of the home region is not recommended.
//

package dashboardservice

import (
	"strings"
)

// SortOrderEnum Enum with underlying type: string
type SortOrderEnum string

// Set of constants representing the allowable values for SortOrderEnum
const (
	SortOrderAsc  SortOrderEnum = "ASC"
	SortOrderDesc SortOrderEnum = "DESC"
)

var mappingSortOrderEnum = map[string]SortOrderEnum{
	"ASC":  SortOrderAsc,
	"DESC": SortOrderDesc,
}

// GetSortOrderEnumValues Enumerates the set of values for SortOrderEnum
func GetSortOrderEnumValues() []SortOrderEnum {
	values := make([]SortOrderEnum, 0)
	for _, v := range mappingSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSortOrderEnumStringValues Enumerates the set of values in String for SortOrderEnum
func GetSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortOrderEnum(val string) (SortOrderEnum, bool) {
	mappingSortOrderEnumIgnoreCase := make(map[string]SortOrderEnum)
	for k, v := range mappingSortOrderEnum {
		mappingSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}