// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SkypeForBusinessOrganizerActivityCounts undocumented
type SkypeForBusinessOrganizerActivityCounts struct {
	// Entity is the base model of SkypeForBusinessOrganizerActivityCounts
	Entity
	// Im undocumented
	Im *int `json:"im,omitempty"`
	// AudioVideo undocumented
	AudioVideo *int `json:"audioVideo,omitempty"`
	// AppSharing undocumented
	AppSharing *int `json:"appSharing,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// DialInOut3rdParty undocumented
	DialInOut3rdParty *int `json:"dialInOut3rdParty,omitempty"`
	// DialInOutMicrosoft undocumented
	DialInOutMicrosoft *int `json:"dialInOutMicrosoft,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
