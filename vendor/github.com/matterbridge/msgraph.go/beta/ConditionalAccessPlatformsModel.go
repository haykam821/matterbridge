// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessPlatforms undocumented
type ConditionalAccessPlatforms struct {
	// Object is the base model of ConditionalAccessPlatforms
	Object
	// IncludePlatforms undocumented
	IncludePlatforms []ConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`
	// ExcludePlatforms undocumented
	ExcludePlatforms []ConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`
}
