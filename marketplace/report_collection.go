// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ReportCollection A collection of reports that match the parameters of the request.
type ReportCollection struct {

	// An array of reports.
	Items []ReportSummary `mandatory:"true" json:"items"`
}

func (m ReportCollection) String() string {
	return common.PointerString(m)
}
