// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ItemPhone undocumented
type ItemPhone struct {
	// ItemFacet is the base model of ItemPhone
	ItemFacet
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *PhoneType `json:"type,omitempty"`
	// Number undocumented
	Number *string `json:"number,omitempty"`
}
