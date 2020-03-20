// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// CustomerPaymentJournal undocumented
type CustomerPaymentJournal struct {
	// Entity is the base model of CustomerPaymentJournal
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// BalancingAccountID undocumented
	BalancingAccountID *UUID `json:"balancingAccountId,omitempty"`
	// BalancingAccountNumber undocumented
	BalancingAccountNumber *string `json:"balancingAccountNumber,omitempty"`
	// CustomerPayments undocumented
	CustomerPayments []CustomerPayment `json:"customerPayments,omitempty"`
	// Account undocumented
	Account *Account `json:"account,omitempty"`
}
