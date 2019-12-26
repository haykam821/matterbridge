// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSStoreAppAssignmentSettings undocumented
type IOSStoreAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of IOSStoreAppAssignmentSettings
	MobileAppAssignmentSettings
	// VpnConfigurationID The VPN Configuration Id to apply for this app.
	VpnConfigurationID *string `json:"vpnConfigurationId,omitempty"`
	// UninstallOnDeviceRemoval Whether or not to uninstall the app when device is removed from Intune.
	UninstallOnDeviceRemoval *bool `json:"uninstallOnDeviceRemoval,omitempty"`
}
