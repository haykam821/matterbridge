// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FeatureRolloutPolicy undocumented
type FeatureRolloutPolicy struct {
	// Entity is the base model of FeatureRolloutPolicy
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Feature undocumented
	Feature *StagedFeatureName `json:"feature,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// IsAppliedToOrganization undocumented
	IsAppliedToOrganization *bool `json:"isAppliedToOrganization,omitempty"`
	// AppliesTo undocumented
	AppliesTo []DirectoryObject `json:"appliesTo,omitempty"`
}
