// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeliveryOptimizationGroupIDSourceOptions undocumented
type DeliveryOptimizationGroupIDSourceOptions struct {
	// DeliveryOptimizationGroupIDSource is the base model of DeliveryOptimizationGroupIDSourceOptions
	DeliveryOptimizationGroupIDSource
	// GroupIDSourceOption Set this policy to restrict peer selection to a specific source.
	GroupIDSourceOption *DeliveryOptimizationGroupIDOptionsType `json:"groupIdSourceOption,omitempty"`
}
