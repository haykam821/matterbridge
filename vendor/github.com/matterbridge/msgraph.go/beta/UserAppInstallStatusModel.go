// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserAppInstallStatus Contains properties for the installation status for a user.
type UserAppInstallStatus struct {
	// Entity is the base model of UserAppInstallStatus
	Entity
	// UserName User name.
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName User Principal Name.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// InstalledDeviceCount Installed Device Count.
	InstalledDeviceCount *int `json:"installedDeviceCount,omitempty"`
	// FailedDeviceCount Failed Device Count.
	FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
	// NotInstalledDeviceCount Not installed device count.
	NotInstalledDeviceCount *int `json:"notInstalledDeviceCount,omitempty"`
	// App undocumented
	App *MobileApp `json:"app,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []MobileAppInstallStatus `json:"deviceStatuses,omitempty"`
}
