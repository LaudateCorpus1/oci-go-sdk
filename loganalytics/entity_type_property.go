// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// EntityTypeProperty Properties used in file patterns specified in log sources.
type EntityTypeProperty struct {

	// Log analytics entity type property name.
	Name *string `mandatory:"true" json:"name"`

	// Description for the log analytics entity type property.
	Description *string `mandatory:"false" json:"description"`
}

func (m EntityTypeProperty) String() string {
	return common.PointerString(m)
}
