// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceAppManagementTaskCategory undocumented
type DeviceAppManagementTaskCategory int

const (
	// DeviceAppManagementTaskCategoryVUnknown undocumented
	DeviceAppManagementTaskCategoryVUnknown DeviceAppManagementTaskCategory = 0
	// DeviceAppManagementTaskCategoryVAdvancedThreatProtection undocumented
	DeviceAppManagementTaskCategoryVAdvancedThreatProtection DeviceAppManagementTaskCategory = 1
)

// DeviceAppManagementTaskCategoryPUnknown returns a pointer to DeviceAppManagementTaskCategoryVUnknown
func DeviceAppManagementTaskCategoryPUnknown() *DeviceAppManagementTaskCategory {
	v := DeviceAppManagementTaskCategoryVUnknown
	return &v
}

// DeviceAppManagementTaskCategoryPAdvancedThreatProtection returns a pointer to DeviceAppManagementTaskCategoryVAdvancedThreatProtection
func DeviceAppManagementTaskCategoryPAdvancedThreatProtection() *DeviceAppManagementTaskCategory {
	v := DeviceAppManagementTaskCategoryVAdvancedThreatProtection
	return &v
}
