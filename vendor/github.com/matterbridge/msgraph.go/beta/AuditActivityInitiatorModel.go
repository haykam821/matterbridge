// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AuditActivityInitiator undocumented
type AuditActivityInitiator struct {
	// Object is the base model of AuditActivityInitiator
	Object
	// User undocumented
	User *UserIdentity `json:"user,omitempty"`
	// App undocumented
	App *AppIdentity `json:"app,omitempty"`
}
