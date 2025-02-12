// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// RecommendationTypeEnum Enum with underlying type: string
type RecommendationTypeEnum string

// Set of constants representing the allowable values for RecommendationTypeEnum
const (
	RecommendationTypeDetectorProblems RecommendationTypeEnum = "DETECTOR_PROBLEMS"
	RecommendationTypeResolvedProblems RecommendationTypeEnum = "RESOLVED_PROBLEMS"
)

var mappingRecommendationTypeEnum = map[string]RecommendationTypeEnum{
	"DETECTOR_PROBLEMS": RecommendationTypeDetectorProblems,
	"RESOLVED_PROBLEMS": RecommendationTypeResolvedProblems,
}

// GetRecommendationTypeEnumValues Enumerates the set of values for RecommendationTypeEnum
func GetRecommendationTypeEnumValues() []RecommendationTypeEnum {
	values := make([]RecommendationTypeEnum, 0)
	for _, v := range mappingRecommendationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRecommendationTypeEnumStringValues Enumerates the set of values in String for RecommendationTypeEnum
func GetRecommendationTypeEnumStringValues() []string {
	return []string{
		"DETECTOR_PROBLEMS",
		"RESOLVED_PROBLEMS",
	}
}

// GetMappingRecommendationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecommendationTypeEnum(val string) (RecommendationTypeEnum, bool) {
	mappingRecommendationTypeEnumIgnoreCase := make(map[string]RecommendationTypeEnum)
	for k, v := range mappingRecommendationTypeEnum {
		mappingRecommendationTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingRecommendationTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
