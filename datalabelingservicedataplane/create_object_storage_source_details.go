// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateObjectStorageSourceDetails Object Storage Source Details.
type CreateObjectStorageSourceDetails struct {

	// The path relative to the prefix specified in the dataset source details (file name).
	RelativePath *string `mandatory:"true" json:"relativePath"`

	// The offset into the file containing the content.
	Offset *float32 `mandatory:"false" json:"offset"`

	// The length from offset into the file containing the content.
	Length *float32 `mandatory:"false" json:"length"`
}

func (m CreateObjectStorageSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateObjectStorageSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateObjectStorageSourceDetails CreateObjectStorageSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateObjectStorageSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeCreateObjectStorageSourceDetails)(m),
	}

	return json.Marshal(&s)
}
