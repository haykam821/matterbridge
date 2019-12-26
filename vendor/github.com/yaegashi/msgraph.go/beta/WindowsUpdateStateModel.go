// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsUpdateState undocumented
type WindowsUpdateState struct {
	// Entity is the base model of WindowsUpdateState
	Entity
	// DeviceID The id of the device.
	DeviceID *string `json:"deviceId,omitempty"`
	// UserID The id of the user.
	UserID *string `json:"userId,omitempty"`
	// DeviceDisplayName Device display name.
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	// UserPrincipalName User principal name.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// Status Windows udpate status.
	Status *WindowsUpdateStatus `json:"status,omitempty"`
	// QualityUpdateVersion The Quality Update Version of the device.
	QualityUpdateVersion *string `json:"qualityUpdateVersion,omitempty"`
	// FeatureUpdateVersion The current feature update version of the device.
	FeatureUpdateVersion *string `json:"featureUpdateVersion,omitempty"`
	// LastScanDateTime The date time that the Windows Update Agent did a successful scan.
	LastScanDateTime *time.Time `json:"lastScanDateTime,omitempty"`
	// LastSyncDateTime Last date time that the device sync with with Microsoft Intune.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
}
