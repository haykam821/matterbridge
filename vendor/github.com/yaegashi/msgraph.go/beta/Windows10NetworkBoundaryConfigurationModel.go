// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10NetworkBoundaryConfiguration Windows10 Network Boundary Configuration
type Windows10NetworkBoundaryConfiguration struct {
	// DeviceConfiguration is the base model of Windows10NetworkBoundaryConfiguration
	DeviceConfiguration
	// WindowsNetworkIsolationPolicy Windows Network Isolation Policy
	WindowsNetworkIsolationPolicy *WindowsNetworkIsolationPolicy `json:"windowsNetworkIsolationPolicy,omitempty"`
}
