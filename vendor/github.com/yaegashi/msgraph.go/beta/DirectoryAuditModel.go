// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DirectoryAudit undocumented
type DirectoryAudit struct {
	// Entity is the base model of DirectoryAudit
	Entity
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// CorrelationID undocumented
	CorrelationID *string `json:"correlationId,omitempty"`
	// Result undocumented
	Result *OperationResult `json:"result,omitempty"`
	// ResultReason undocumented
	ResultReason *string `json:"resultReason,omitempty"`
	// ActivityDisplayName undocumented
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// LoggedByService undocumented
	LoggedByService *string `json:"loggedByService,omitempty"`
	// OperationType undocumented
	OperationType *string `json:"operationType,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	// TargetResources undocumented
	TargetResources []TargetResource `json:"targetResources,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails []KeyValue `json:"additionalDetails,omitempty"`
}
