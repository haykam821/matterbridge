// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EvaluateLabelJobResultGroup undocumented
type EvaluateLabelJobResultGroup struct {
	// Object is the base model of EvaluateLabelJobResultGroup
	Object
	// Automatic undocumented
	Automatic *EvaluateLabelJobResult `json:"automatic,omitempty"`
	// Recommended undocumented
	Recommended *EvaluateLabelJobResult `json:"recommended,omitempty"`
}
