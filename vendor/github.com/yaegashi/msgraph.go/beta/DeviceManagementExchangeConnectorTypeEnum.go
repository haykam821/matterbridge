// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeConnectorType undocumented
type DeviceManagementExchangeConnectorType int

const (
	// DeviceManagementExchangeConnectorTypeVOnPremises undocumented
	DeviceManagementExchangeConnectorTypeVOnPremises DeviceManagementExchangeConnectorType = 0
	// DeviceManagementExchangeConnectorTypeVHosted undocumented
	DeviceManagementExchangeConnectorTypeVHosted DeviceManagementExchangeConnectorType = 1
	// DeviceManagementExchangeConnectorTypeVServiceToService undocumented
	DeviceManagementExchangeConnectorTypeVServiceToService DeviceManagementExchangeConnectorType = 2
	// DeviceManagementExchangeConnectorTypeVDedicated undocumented
	DeviceManagementExchangeConnectorTypeVDedicated DeviceManagementExchangeConnectorType = 3
)

// DeviceManagementExchangeConnectorTypePOnPremises returns a pointer to DeviceManagementExchangeConnectorTypeVOnPremises
func DeviceManagementExchangeConnectorTypePOnPremises() *DeviceManagementExchangeConnectorType {
	v := DeviceManagementExchangeConnectorTypeVOnPremises
	return &v
}

// DeviceManagementExchangeConnectorTypePHosted returns a pointer to DeviceManagementExchangeConnectorTypeVHosted
func DeviceManagementExchangeConnectorTypePHosted() *DeviceManagementExchangeConnectorType {
	v := DeviceManagementExchangeConnectorTypeVHosted
	return &v
}

// DeviceManagementExchangeConnectorTypePServiceToService returns a pointer to DeviceManagementExchangeConnectorTypeVServiceToService
func DeviceManagementExchangeConnectorTypePServiceToService() *DeviceManagementExchangeConnectorType {
	v := DeviceManagementExchangeConnectorTypeVServiceToService
	return &v
}

// DeviceManagementExchangeConnectorTypePDedicated returns a pointer to DeviceManagementExchangeConnectorTypeVDedicated
func DeviceManagementExchangeConnectorTypePDedicated() *DeviceManagementExchangeConnectorType {
	v := DeviceManagementExchangeConnectorTypeVDedicated
	return &v
}
