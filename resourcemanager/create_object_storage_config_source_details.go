// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateObjectStorageConfigSourceDetails Details for the Object Storage bucket that contains Terraform configuration files.
type CreateObjectStorageConfigSourceDetails struct {

	// The name of the bucket's region.
	// Example: `PHX`
	Region *string `mandatory:"true" json:"region"`

	// The Object Storage namespace that contains the bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket that contains the Terraform configuration files.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CreateObjectStorageConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CreateObjectStorageConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateObjectStorageConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateObjectStorageConfigSourceDetails CreateObjectStorageConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCreateObjectStorageConfigSourceDetails
	}{
		"OBJECT_STORAGE_CONFIG_SOURCE",
		(MarshalTypeCreateObjectStorageConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
