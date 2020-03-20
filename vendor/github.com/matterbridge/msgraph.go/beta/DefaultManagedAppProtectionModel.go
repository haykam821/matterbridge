// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DefaultManagedAppProtection Policy used to configure detailed management settings for a specified set of apps for all users not targeted by a TargetedManagedAppProtection Policy
type DefaultManagedAppProtection struct {
	// ManagedAppProtection is the base model of DefaultManagedAppProtection
	ManagedAppProtection
	// AppDataEncryptionType Type of encryption which should be used for data in a managed app. (iOS Only)
	AppDataEncryptionType *ManagedAppDataEncryptionType `json:"appDataEncryptionType,omitempty"`
	// ScreenCaptureBlocked Indicates whether screen capture is blocked. (Android only)
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// EncryptAppData Indicates whether managed-app data should be encrypted. (Android only)
	EncryptAppData *bool `json:"encryptAppData,omitempty"`
	// DisableAppEncryptionIfDeviceEncryptionIsEnabled When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool `json:"disableAppEncryptionIfDeviceEncryptionIsEnabled,omitempty"`
	// MinimumRequiredSdkVersion Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
	MinimumRequiredSdkVersion *string `json:"minimumRequiredSdkVersion,omitempty"`
	// CustomSettings A set of string key and string value pairs to be sent to the affected users, unalterned by this service
	CustomSettings []KeyValuePair `json:"customSettings,omitempty"`
	// DeployedAppCount Count of apps to which the current policy is deployed.
	DeployedAppCount *int `json:"deployedAppCount,omitempty"`
	// MinimumRequiredPatchVersion Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
	MinimumRequiredPatchVersion *string `json:"minimumRequiredPatchVersion,omitempty"`
	// MinimumWarningPatchVersion Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
	MinimumWarningPatchVersion *string `json:"minimumWarningPatchVersion,omitempty"`
	// ExemptedAppProtocols iOS Apps in this list will be exempt from the policy and will be able to receive data from managed apps. (iOS Only)
	ExemptedAppProtocols []KeyValuePair `json:"exemptedAppProtocols,omitempty"`
	// ExemptedAppPackages Android App packages in this list will be exempt from the policy and will be able to receive data from managed apps. (Android only)
	ExemptedAppPackages []KeyValuePair `json:"exemptedAppPackages,omitempty"`
	// FaceIDBlocked Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
	FaceIDBlocked *bool `json:"faceIdBlocked,omitempty"`
	// MinimumWipeSdkVersion Versions less than the specified version will block the managed app from accessing company data.
	MinimumWipeSdkVersion *string `json:"minimumWipeSdkVersion,omitempty"`
	// MinimumWipePatchVersion Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data. (Android only)
	MinimumWipePatchVersion *string `json:"minimumWipePatchVersion,omitempty"`
	// AllowedIOSDeviceModels Semicolon seperated list of device models allowed, as a string, for the managed app to work. (iOS Only)
	AllowedIOSDeviceModels *string `json:"allowedIosDeviceModels,omitempty"`
	// AppActionIfIOSDeviceModelNotAllowed Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. (iOS Only)
	AppActionIfIOSDeviceModelNotAllowed *ManagedAppRemediationAction `json:"appActionIfIosDeviceModelNotAllowed,omitempty"`
	// AllowedAndroidDeviceManufacturers Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work. (Android only)
	AllowedAndroidDeviceManufacturers *string `json:"allowedAndroidDeviceManufacturers,omitempty"`
	// AppActionIfAndroidDeviceManufacturerNotAllowed Defines a managed app behavior, either block or wipe, if the specified device manufacturer is not allowed. (Android only)
	AppActionIfAndroidDeviceManufacturerNotAllowed *ManagedAppRemediationAction `json:"appActionIfAndroidDeviceManufacturerNotAllowed,omitempty"`
	// ThirdPartyKeyboardsBlocked Defines if third party keyboards are allowed while accessing a managed app. (iOS Only)
	ThirdPartyKeyboardsBlocked *bool `json:"thirdPartyKeyboardsBlocked,omitempty"`
	// FilterOpenInToOnlyManagedApps Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False. (iOS Only)
	FilterOpenInToOnlyManagedApps *bool `json:"filterOpenInToOnlyManagedApps,omitempty"`
	// DisableProtectionOfManagedOutboundOpenInData Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps. (iOS Only)
	DisableProtectionOfManagedOutboundOpenInData *bool `json:"disableProtectionOfManagedOutboundOpenInData,omitempty"`
	// ProtectInboundDataFromUnknownSources Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps. (iOS Only)
	ProtectInboundDataFromUnknownSources *bool `json:"protectInboundDataFromUnknownSources,omitempty"`
	// RequiredAndroidSafetyNetDeviceAttestationType Defines the Android SafetyNet Device Attestation requirement for a managed app to work.
	RequiredAndroidSafetyNetDeviceAttestationType *AndroidManagedAppSafetyNetDeviceAttestationType `json:"requiredAndroidSafetyNetDeviceAttestationType,omitempty"`
	// AppActionIfAndroidSafetyNetDeviceAttestationFailed Defines a managed app behavior, either warn or block, if the specified Android SafetyNet Attestation requirment fails.
	AppActionIfAndroidSafetyNetDeviceAttestationFailed *ManagedAppRemediationAction `json:"appActionIfAndroidSafetyNetDeviceAttestationFailed,omitempty"`
	// RequiredAndroidSafetyNetAppsVerificationType Defines the Android SafetyNet Apps Verification requirement for a managed app to work.
	RequiredAndroidSafetyNetAppsVerificationType *AndroidManagedAppSafetyNetAppsVerificationType `json:"requiredAndroidSafetyNetAppsVerificationType,omitempty"`
	// AppActionIfAndroidSafetyNetAppsVerificationFailed Defines a managed app behavior, either warn or block, if the specified Android App Verification requirment fails.
	AppActionIfAndroidSafetyNetAppsVerificationFailed *ManagedAppRemediationAction `json:"appActionIfAndroidSafetyNetAppsVerificationFailed,omitempty"`
	// CustomBrowserProtocol A custom browser protocol to open weblink on iOS. (iOS only)
	CustomBrowserProtocol *string `json:"customBrowserProtocol,omitempty"`
	// CustomBrowserPackageID Unique identifier of a custom browser to open weblink on Android. (Android only)
	CustomBrowserPackageID *string `json:"customBrowserPackageId,omitempty"`
	// CustomBrowserDisplayName Friendly name of the preferred custom browser to open weblink on Android. (Android only)
	CustomBrowserDisplayName *string `json:"customBrowserDisplayName,omitempty"`
	// MinimumRequiredCompanyPortalVersion Minimum version of the Company portal that must be installed on the device or app access will be blocked
	MinimumRequiredCompanyPortalVersion *string `json:"minimumRequiredCompanyPortalVersion,omitempty"`
	// MinimumWarningCompanyPortalVersion Minimum version of the Company portal that must be installed on the device or the user will receive a warning
	MinimumWarningCompanyPortalVersion *string `json:"minimumWarningCompanyPortalVersion,omitempty"`
	// MinimumWipeCompanyPortalVersion Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
	MinimumWipeCompanyPortalVersion *string `json:"minimumWipeCompanyPortalVersion,omitempty"`
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
	// DeploymentSummary undocumented
	DeploymentSummary *ManagedAppPolicyDeploymentSummary `json:"deploymentSummary,omitempty"`
}
