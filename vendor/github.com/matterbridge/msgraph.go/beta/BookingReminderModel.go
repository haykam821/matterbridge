// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// BookingReminder undocumented
type BookingReminder struct {
	// Object is the base model of BookingReminder
	Object
	// Offset How much time before an appointment the reminder should be sent.
	Offset *time.Duration `json:"offset,omitempty"`
	// Recipients Who should receive the reminder.
	Recipients *BookingReminderRecipients `json:"recipients,omitempty"`
	// Message Message to send.
	Message *string `json:"message,omitempty"`
}
