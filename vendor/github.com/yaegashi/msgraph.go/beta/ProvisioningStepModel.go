// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProvisioningStep undocumented
type ProvisioningStep struct {
	// Object is the base model of ProvisioningStep
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Status undocumented
	Status *ProvisioningResult `json:"status,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Details undocumented
	Details *DetailsInfo `json:"details,omitempty"`
	// ProvisioningStepType undocumented
	ProvisioningStepType *ProvisioningStepType `json:"provisioningStepType,omitempty"`
}
