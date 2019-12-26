// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationComboBox Represents an ADMX comboBox element and an ADMX text element.
type GroupPolicyPresentationComboBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationComboBox
	GroupPolicyPresentation
	// DefaultValue Localized default string displayed in the combo box. The default value is empty.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Suggestions Localized strings listed in the drop-down list of the combo box. The default value is empty.
	Suggestions []string `json:"suggestions,omitempty"`
	// Required Specifies whether a value must be specified for the parameter. The default value is false.
	Required *bool `json:"required,omitempty"`
	// MaxLength An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
	MaxLength *int `json:"maxLength,omitempty"`
}
