// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AssignedPlan undocumented
type AssignedPlan struct {
	// Object is the base model of AssignedPlan
	Object
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// Service undocumented
	Service *string `json:"service,omitempty"`
	// ServicePlanID undocumented
	ServicePlanID *UUID `json:"servicePlanId,omitempty"`
}
