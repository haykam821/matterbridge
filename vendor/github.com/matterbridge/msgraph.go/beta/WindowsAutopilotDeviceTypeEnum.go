// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsAutopilotDeviceType undocumented
type WindowsAutopilotDeviceType int

const (
	// WindowsAutopilotDeviceTypeVWindowsPc undocumented
	WindowsAutopilotDeviceTypeVWindowsPc WindowsAutopilotDeviceType = 0
	// WindowsAutopilotDeviceTypeVSurfaceHub2 undocumented
	WindowsAutopilotDeviceTypeVSurfaceHub2 WindowsAutopilotDeviceType = 1
)

// WindowsAutopilotDeviceTypePWindowsPc returns a pointer to WindowsAutopilotDeviceTypeVWindowsPc
func WindowsAutopilotDeviceTypePWindowsPc() *WindowsAutopilotDeviceType {
	v := WindowsAutopilotDeviceTypeVWindowsPc
	return &v
}

// WindowsAutopilotDeviceTypePSurfaceHub2 returns a pointer to WindowsAutopilotDeviceTypeVSurfaceHub2
func WindowsAutopilotDeviceTypePSurfaceHub2() *WindowsAutopilotDeviceType {
	v := WindowsAutopilotDeviceTypeVSurfaceHub2
	return &v
}
