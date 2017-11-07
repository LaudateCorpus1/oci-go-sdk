// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// ListCertificatesRequest wrapper for the ListCertificates operation
type ListCertificatesRequest struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer associated with the certificates to be listed.
	LoadBalancerID *string `mandatory:"true" contributesTo:"path" name:"loadBalancerId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`
}

func (request ListCertificatesRequest) String() string {
	return common.PointerString(request)
}

// ListCertificatesResponse wrapper for the ListCertificates operation
type ListCertificatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []Certificate instance
	Items []Certificate `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCertificatesResponse) String() string {
	return common.PointerString(response)
}
