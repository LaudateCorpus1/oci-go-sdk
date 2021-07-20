// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// UpdateTaskFromSqlTask The information about the SQL task.
type UpdateTaskFromSqlTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Script *Script `mandatory:"false" json:"script"`

	// Describes the shape of the execution result
	Operation *interface{} `mandatory:"false" json:"operation"`

	// Indicates whether the task is invoking a custom SQL script or stored procedure.
	SqlScriptType UpdateTaskFromSqlTaskSqlScriptTypeEnum `mandatory:"false" json:"sqlScriptType,omitempty"`
}

//GetKey returns Key
func (m UpdateTaskFromSqlTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateTaskFromSqlTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateTaskFromSqlTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateTaskFromSqlTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateTaskFromSqlTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromSqlTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromSqlTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateTaskFromSqlTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m UpdateTaskFromSqlTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m UpdateTaskFromSqlTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m UpdateTaskFromSqlTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromSqlTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromSqlTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromSqlTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromSqlTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateTaskFromSqlTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTaskFromSqlTask UpdateTaskFromSqlTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateTaskFromSqlTask
	}{
		"SQL_TASK",
		(MarshalTypeUpdateTaskFromSqlTask)(m),
	}

	return json.Marshal(&s)
}

// UpdateTaskFromSqlTaskSqlScriptTypeEnum Enum with underlying type: string
type UpdateTaskFromSqlTaskSqlScriptTypeEnum string

// Set of constants representing the allowable values for UpdateTaskFromSqlTaskSqlScriptTypeEnum
const (
	UpdateTaskFromSqlTaskSqlScriptTypeStoredProcedure UpdateTaskFromSqlTaskSqlScriptTypeEnum = "STORED_PROCEDURE"
	UpdateTaskFromSqlTaskSqlScriptTypeSqlCode         UpdateTaskFromSqlTaskSqlScriptTypeEnum = "SQL_CODE"
)

var mappingUpdateTaskFromSqlTaskSqlScriptType = map[string]UpdateTaskFromSqlTaskSqlScriptTypeEnum{
	"STORED_PROCEDURE": UpdateTaskFromSqlTaskSqlScriptTypeStoredProcedure,
	"SQL_CODE":         UpdateTaskFromSqlTaskSqlScriptTypeSqlCode,
}

// GetUpdateTaskFromSqlTaskSqlScriptTypeEnumValues Enumerates the set of values for UpdateTaskFromSqlTaskSqlScriptTypeEnum
func GetUpdateTaskFromSqlTaskSqlScriptTypeEnumValues() []UpdateTaskFromSqlTaskSqlScriptTypeEnum {
	values := make([]UpdateTaskFromSqlTaskSqlScriptTypeEnum, 0)
	for _, v := range mappingUpdateTaskFromSqlTaskSqlScriptType {
		values = append(values, v)
	}
	return values
}
