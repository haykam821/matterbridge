// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceAndAppManagementAssignmentSource undocumented
type DeviceAndAppManagementAssignmentSource int

const (
	// DeviceAndAppManagementAssignmentSourceVDirect undocumented
	DeviceAndAppManagementAssignmentSourceVDirect DeviceAndAppManagementAssignmentSource = 0
	// DeviceAndAppManagementAssignmentSourceVPolicySets undocumented
	DeviceAndAppManagementAssignmentSourceVPolicySets DeviceAndAppManagementAssignmentSource = 1
)

// DeviceAndAppManagementAssignmentSourcePDirect returns a pointer to DeviceAndAppManagementAssignmentSourceVDirect
func DeviceAndAppManagementAssignmentSourcePDirect() *DeviceAndAppManagementAssignmentSource {
	v := DeviceAndAppManagementAssignmentSourceVDirect
	return &v
}

// DeviceAndAppManagementAssignmentSourcePPolicySets returns a pointer to DeviceAndAppManagementAssignmentSourceVPolicySets
func DeviceAndAppManagementAssignmentSourcePPolicySets() *DeviceAndAppManagementAssignmentSource {
	v := DeviceAndAppManagementAssignmentSourceVPolicySets
	return &v
}
