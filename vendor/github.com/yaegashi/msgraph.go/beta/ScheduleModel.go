// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Schedule undocumented
type Schedule struct {
	// Entity is the base model of Schedule
	Entity
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
	// ProvisionStatus undocumented
	ProvisionStatus *OperationStatus `json:"provisionStatus,omitempty"`
	// ProvisionStatusCode undocumented
	ProvisionStatusCode *string `json:"provisionStatusCode,omitempty"`
	// WorkforceIntegrationIDs undocumented
	WorkforceIntegrationIDs []string `json:"workforceIntegrationIds,omitempty"`
	// TimeClockEnabled undocumented
	TimeClockEnabled *bool `json:"timeClockEnabled,omitempty"`
	// OpenShiftsEnabled undocumented
	OpenShiftsEnabled *bool `json:"openShiftsEnabled,omitempty"`
	// SwapShiftsRequestsEnabled undocumented
	SwapShiftsRequestsEnabled *bool `json:"swapShiftsRequestsEnabled,omitempty"`
	// OfferShiftRequestsEnabled undocumented
	OfferShiftRequestsEnabled *bool `json:"offerShiftRequestsEnabled,omitempty"`
	// TimeOffRequestsEnabled undocumented
	TimeOffRequestsEnabled *bool `json:"timeOffRequestsEnabled,omitempty"`
	// Shifts undocumented
	Shifts []Shift `json:"shifts,omitempty"`
	// OpenShifts undocumented
	OpenShifts []OpenShift `json:"openShifts,omitempty"`
	// TimesOff undocumented
	TimesOff []TimeOff `json:"timesOff,omitempty"`
	// TimeOffReasons undocumented
	TimeOffReasons []TimeOffReason `json:"timeOffReasons,omitempty"`
	// SchedulingGroups undocumented
	SchedulingGroups []SchedulingGroup `json:"schedulingGroups,omitempty"`
	// SwapShiftsChangeRequests undocumented
	SwapShiftsChangeRequests []SwapShiftsChangeRequestObject `json:"swapShiftsChangeRequests,omitempty"`
	// OpenShiftChangeRequests undocumented
	OpenShiftChangeRequests []OpenShiftChangeRequestObject `json:"openShiftChangeRequests,omitempty"`
	// TimeOffRequests undocumented
	TimeOffRequests []TimeOffRequestObject `json:"timeOffRequests,omitempty"`
}
