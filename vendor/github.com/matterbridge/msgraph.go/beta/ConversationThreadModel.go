// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ConversationThread undocumented
type ConversationThread struct {
	// Entity is the base model of ConversationThread
	Entity
	// ToRecipients undocumented
	ToRecipients []Recipient `json:"toRecipients,omitempty"`
	// Topic undocumented
	Topic *string `json:"topic,omitempty"`
	// HasAttachments undocumented
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// LastDeliveredDateTime undocumented
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// UniqueSenders undocumented
	UniqueSenders []string `json:"uniqueSenders,omitempty"`
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// Preview undocumented
	Preview *string `json:"preview,omitempty"`
	// IsLocked undocumented
	IsLocked *bool `json:"isLocked,omitempty"`
	// Posts undocumented
	Posts []Post `json:"posts,omitempty"`
}
