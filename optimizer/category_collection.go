// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CategoryCollection A list of categories that match filter criteria, if any. Results contain `CategorySummary` objects.
type CategoryCollection struct {

	// A collection of category summaries.
	Items []CategorySummary `mandatory:"true" json:"items"`
}

func (m CategoryCollection) String() string {
	return common.PointerString(m)
}
