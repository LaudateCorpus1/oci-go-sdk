// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AnnotationCollection The results of an annotations search. It contains AnnotationSummary items and other information, such as metadata.
type AnnotationCollection struct {

	// The list of annotations.
	Items []AnnotationSummary `mandatory:"true" json:"items"`
}

func (m AnnotationCollection) String() string {
	return common.PointerString(m)
}
