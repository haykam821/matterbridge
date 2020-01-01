// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ManagedDeviceOverview Summary data for managed devices
type ManagedDeviceOverview struct {
	// Entity is the base model of ManagedDeviceOverview
	Entity
	// EnrolledDeviceCount Total enrolled device count. Does not include PC devices managed via Intune PC Agent
	EnrolledDeviceCount *int `json:"enrolledDeviceCount,omitempty"`
	// MdmEnrolledCount The number of devices enrolled in MDM
	MdmEnrolledCount *int `json:"mdmEnrolledCount,omitempty"`
	// DualEnrolledDeviceCount The number of devices enrolled in both MDM and EAS
	DualEnrolledDeviceCount *int `json:"dualEnrolledDeviceCount,omitempty"`
	// DeviceOperatingSystemSummary Device operating system summary.
	DeviceOperatingSystemSummary *DeviceOperatingSystemSummary `json:"deviceOperatingSystemSummary,omitempty"`
	// DeviceExchangeAccessStateSummary Distribution of Exchange Access State in Intune
	DeviceExchangeAccessStateSummary *DeviceExchangeAccessStateSummary `json:"deviceExchangeAccessStateSummary,omitempty"`
	// ManagedDeviceModelsAndManufacturers Models and Manufactures meatadata for managed devices in the account
	ManagedDeviceModelsAndManufacturers *ManagedDeviceModelsAndManufacturers `json:"managedDeviceModelsAndManufacturers,omitempty"`
	// LastModifiedDateTime Last modified date time of device overview
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}