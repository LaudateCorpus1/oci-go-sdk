// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreditCardPaymentOption Credit card Payment related details
type CreditCardPaymentOption struct {

	// Wallet instrument internal id.
	WalletInstrumentId *string `mandatory:"false" json:"walletInstrumentId"`

	// Wallet transaction id.
	WalletTransactionId *string `mandatory:"false" json:"walletTransactionId"`
}

//GetWalletInstrumentId returns WalletInstrumentId
func (m CreditCardPaymentOption) GetWalletInstrumentId() *string {
	return m.WalletInstrumentId
}

//GetWalletTransactionId returns WalletTransactionId
func (m CreditCardPaymentOption) GetWalletTransactionId() *string {
	return m.WalletTransactionId
}

func (m CreditCardPaymentOption) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreditCardPaymentOption) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreditCardPaymentOption) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreditCardPaymentOption CreditCardPaymentOption
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypeCreditCardPaymentOption
	}{
		"CREDIT_CARD",
		(MarshalTypeCreditCardPaymentOption)(m),
	}

	return json.Marshal(&s)
}
