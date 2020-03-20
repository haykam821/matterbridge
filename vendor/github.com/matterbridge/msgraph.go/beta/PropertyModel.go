// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Property undocumented
type Property struct {
	// Object is the base model of Property
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *PropertyType `json:"type,omitempty"`
	// IsSearchable undocumented
	IsSearchable *bool `json:"isSearchable,omitempty"`
	// IsRetrievable undocumented
	IsRetrievable *bool `json:"isRetrievable,omitempty"`
	// IsQueryable undocumented
	IsQueryable *bool `json:"isQueryable,omitempty"`
}
