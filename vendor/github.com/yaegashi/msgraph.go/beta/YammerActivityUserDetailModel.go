// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// YammerActivityUserDetail undocumented
type YammerActivityUserDetail struct {
	// Entity is the base model of YammerActivityUserDetail
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UserState undocumented
	UserState *string `json:"userState,omitempty"`
	// StateChangeDate undocumented
	StateChangeDate *time.Time `json:"stateChangeDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *time.Time `json:"lastActivityDate,omitempty"`
	// PostedCount undocumented
	PostedCount *int `json:"postedCount,omitempty"`
	// ReadCount undocumented
	ReadCount *int `json:"readCount,omitempty"`
	// LikedCount undocumented
	LikedCount *int `json:"likedCount,omitempty"`
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
