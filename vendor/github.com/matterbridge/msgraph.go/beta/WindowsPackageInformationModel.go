// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPackageInformation undocumented
type WindowsPackageInformation struct {
	// Object is the base model of WindowsPackageInformation
	Object
	// ApplicableArchitecture The Windows architecture for which this app can run on.
	ApplicableArchitecture *WindowsArchitecture `json:"applicableArchitecture,omitempty"`
	// DisplayName The Display Name.
	DisplayName *string `json:"displayName,omitempty"`
	// IdentityName The Identity Name.
	IdentityName *string `json:"identityName,omitempty"`
	// IdentityPublisher The Identity Publisher.
	IdentityPublisher *string `json:"identityPublisher,omitempty"`
	// IdentityResourceIdentifier The Identity Resource Identifier.
	IdentityResourceIdentifier *string `json:"identityResourceIdentifier,omitempty"`
	// IdentityVersion The Identity Version.
	IdentityVersion *string `json:"identityVersion,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}
