// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// DeliveredArtifactCollection Specifies the list of artifacts delivered through the Deliver Artifacts stage.
type DeliveredArtifactCollection struct {

	// List of artifacts delivered through the Deliver Artifacts stage.
	Items []DeliveredArtifact `mandatory:"true" json:"items"`
}

func (m DeliveredArtifactCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DeliveredArtifactCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []deliveredartifact `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DeliveredArtifact, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DeliveredArtifact)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
