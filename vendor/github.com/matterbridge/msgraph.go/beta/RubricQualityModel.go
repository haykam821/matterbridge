// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RubricQuality undocumented
type RubricQuality struct {
	// Object is the base model of RubricQuality
	Object
	// QualityID undocumented
	QualityID *string `json:"qualityId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// Weight undocumented
	Weight *float64 `json:"weight,omitempty"`
	// Criteria undocumented
	Criteria []RubricCriterion `json:"criteria,omitempty"`
}
