// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ApplicationIdentityType undocumented
type ApplicationIdentityType int

const (
	// ApplicationIdentityTypeVAadApplication undocumented
	ApplicationIdentityTypeVAadApplication ApplicationIdentityType = 0
	// ApplicationIdentityTypeVBot undocumented
	ApplicationIdentityTypeVBot ApplicationIdentityType = 1
	// ApplicationIdentityTypeVTenantBot undocumented
	ApplicationIdentityTypeVTenantBot ApplicationIdentityType = 2
	// ApplicationIdentityTypeVOffice365Connector undocumented
	ApplicationIdentityTypeVOffice365Connector ApplicationIdentityType = 3
	// ApplicationIdentityTypeVOutgoingWebhook undocumented
	ApplicationIdentityTypeVOutgoingWebhook ApplicationIdentityType = 4
)

// ApplicationIdentityTypePAadApplication returns a pointer to ApplicationIdentityTypeVAadApplication
func ApplicationIdentityTypePAadApplication() *ApplicationIdentityType {
	v := ApplicationIdentityTypeVAadApplication
	return &v
}

// ApplicationIdentityTypePBot returns a pointer to ApplicationIdentityTypeVBot
func ApplicationIdentityTypePBot() *ApplicationIdentityType {
	v := ApplicationIdentityTypeVBot
	return &v
}

// ApplicationIdentityTypePTenantBot returns a pointer to ApplicationIdentityTypeVTenantBot
func ApplicationIdentityTypePTenantBot() *ApplicationIdentityType {
	v := ApplicationIdentityTypeVTenantBot
	return &v
}

// ApplicationIdentityTypePOffice365Connector returns a pointer to ApplicationIdentityTypeVOffice365Connector
func ApplicationIdentityTypePOffice365Connector() *ApplicationIdentityType {
	v := ApplicationIdentityTypeVOffice365Connector
	return &v
}

// ApplicationIdentityTypePOutgoingWebhook returns a pointer to ApplicationIdentityTypeVOutgoingWebhook
func ApplicationIdentityTypePOutgoingWebhook() *ApplicationIdentityType {
	v := ApplicationIdentityTypeVOutgoingWebhook
	return &v
}
