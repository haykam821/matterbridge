// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SignInActivity undocumented
type SignInActivity struct {
	// Object is the base model of SignInActivity
	Object
	// LastSignInDateTime undocumented
	LastSignInDateTime *time.Time `json:"lastSignInDateTime,omitempty"`
	// LastSignInRequestID undocumented
	LastSignInRequestID *string `json:"lastSignInRequestId,omitempty"`
}
