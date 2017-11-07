// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// LoadBalancerHealth. The health status details for the specified load balancer.
// This object does not explicitly enumerate backend sets with a status of `OK`. However, they are included in the
// `totalBackendSetCount` sum.
type LoadBalancerHealth struct {

	// A list of backend sets that are currently in the `CRITICAL` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `My_backend_set`
	CriticalStateBackendSetNames *[]string `mandatory:"true" json:"criticalStateBackendSetNames,omitempty"`

	// The overall health status of the load balancer.
	// *  **OK:** All backend sets associated with the load balancer return a status of `OK`.
	// *  **WARNING:** At least one of the backend sets associated with the load balancer returns a status of `WARNING`,
	// no backend sets return a status of `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	// *  **CRITICAL:** One or more of the backend sets associated with the load balancer return a status of `CRITICAL`.
	// *  **UNKNOWN:** If any one of the following conditions is true:
	//     *  The load balancer life cycle state is not `ACTIVE`.
	//     *  No backend sets are defined for the load balancer.
	//     *  More than half of the backend sets associated with the load balancer return a status of `UNKNOWN`, none of the backend
	//        sets return a status of `WARNING` or `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	//     *  The system could not retrieve metrics for any reason.
	Status LoadBalancerHealthStatusEnum `mandatory:"true" json:"status,omitempty"`

	// The total number of backend sets associated with this load balancer.
	// Example: `4`
	TotalBackendSetCount *int `mandatory:"true" json:"totalBackendSetCount,omitempty"`

	// A list of backend sets that are currently in the `UNKNOWN` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `Backend_set2`
	UnknownStateBackendSetNames *[]string `mandatory:"true" json:"unknownStateBackendSetNames,omitempty"`

	// A list of backend sets that are currently in the `WARNING` health state. The list identifies each backend set by the
	// friendly name you assigned when you created it.
	// Example: `Backend_set3`
	WarningStateBackendSetNames *[]string `mandatory:"true" json:"warningStateBackendSetNames,omitempty"`
}

func (model LoadBalancerHealth) String() string {
	return common.PointerString(model)
}

type LoadBalancerHealthStatusEnum string

const (
	LOAD_BALANCER_HEALTH_STATUS_OK       LoadBalancerHealthStatusEnum = "OK"
	LOAD_BALANCER_HEALTH_STATUS_WARNING  LoadBalancerHealthStatusEnum = "WARNING"
	LOAD_BALANCER_HEALTH_STATUS_CRITICAL LoadBalancerHealthStatusEnum = "CRITICAL"
	LOAD_BALANCER_HEALTH_STATUS_UNKNOWN  LoadBalancerHealthStatusEnum = "UNKNOWN"
	LOAD_BALANCER_HEALTH_STATUS_UNKNOWN  LoadBalancerHealthStatusEnum = "UNKNOWN"
)

var mapping_loadbalancerhealth_status = map[string]LoadBalancerHealthStatusEnum{
	"OK":       LOAD_BALANCER_HEALTH_STATUS_OK,
	"WARNING":  LOAD_BALANCER_HEALTH_STATUS_WARNING,
	"CRITICAL": LOAD_BALANCER_HEALTH_STATUS_CRITICAL,
	"UNKNOWN":  LOAD_BALANCER_HEALTH_STATUS_UNKNOWN,
	"UNKNOWN":  LOAD_BALANCER_HEALTH_STATUS_UNKNOWN,
}

func GetLoadBalancerHealthStatusEnumValues() []LoadBalancerHealthStatusEnum {
	values := make([]LoadBalancerHealthStatusEnum, 0)
	for _, v := range mapping_loadbalancerhealth_status {
		if v != LOAD_BALANCER_HEALTH_STATUS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
