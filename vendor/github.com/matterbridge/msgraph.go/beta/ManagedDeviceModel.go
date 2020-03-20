// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ManagedDevice Devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
	// Entity is the base model of ManagedDevice
	Entity
	// UserID Unique Identifier for the user associated with the device. This property is read-only.
	UserID *string `json:"userId,omitempty"`
	// DeviceName Name of the device. This property is read-only.
	DeviceName *string `json:"deviceName,omitempty"`
	// HardwareInformation The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
	HardwareInformation *HardwareInformation `json:"hardwareInformation,omitempty"`
	// OwnerType Ownership of the device. Can be 'company' or 'personal'
	OwnerType *OwnerType `json:"ownerType,omitempty"`
	// ManagedDeviceOwnerType Ownership of the device. Can be 'company' or 'personal'
	ManagedDeviceOwnerType *ManagedDeviceOwnerType `json:"managedDeviceOwnerType,omitempty"`
	// DeviceActionResults List of ComplexType deviceActionResult objects. This property is read-only.
	DeviceActionResults []DeviceActionResult `json:"deviceActionResults,omitempty"`
	// ManagementState Management state of the device. This property is read-only.
	ManagementState *ManagementState `json:"managementState,omitempty"`
	// EnrolledDateTime Enrollment time of the device. This property is read-only.
	EnrolledDateTime *time.Time `json:"enrolledDateTime,omitempty"`
	// LastSyncDateTime The date and time that the device last completed a successful sync with Intune. This property is read-only.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// ChassisType Chassis type of the device. This property is read-only.
	ChassisType *ChassisType `json:"chassisType,omitempty"`
	// OperatingSystem Operating system of the device. Windows, iOS, etc. This property is read-only.
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// DeviceType Platform of the device. This property is read-only.
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// ComplianceState Compliance state of the device. This property is read-only.
	ComplianceState *ComplianceState `json:"complianceState,omitempty"`
	// JailBroken whether the device is jail broken or rooted. This property is read-only.
	JailBroken *string `json:"jailBroken,omitempty"`
	// ManagementAgent Management channel of the device. Intune, EAS, etc. This property is read-only.
	ManagementAgent *ManagementAgentType `json:"managementAgent,omitempty"`
	// OsVersion Operating system version of the device. This property is read-only.
	OsVersion *string `json:"osVersion,omitempty"`
	// EasActivated Whether the device is Exchange ActiveSync activated. This property is read-only.
	EasActivated *bool `json:"easActivated,omitempty"`
	// EasDeviceID Exchange ActiveSync Id of the device. This property is read-only.
	EasDeviceID *string `json:"easDeviceId,omitempty"`
	// EasActivationDateTime Exchange ActivationSync activation time of the device. This property is read-only.
	EasActivationDateTime *time.Time `json:"easActivationDateTime,omitempty"`
	// AadRegistered Whether the device is Azure Active Directory registered. This property is read-only.
	AadRegistered *bool `json:"aadRegistered,omitempty"`
	// AzureADRegistered Whether the device is Azure Active Directory registered. This property is read-only.
	AzureADRegistered *bool `json:"azureADRegistered,omitempty"`
	// DeviceEnrollmentType Enrollment type of the device. This property is read-only.
	DeviceEnrollmentType *DeviceEnrollmentType `json:"deviceEnrollmentType,omitempty"`
	// LostModeState Indicates if Lost mode is enabled or disabled. This property is read-only.
	LostModeState *LostModeState `json:"lostModeState,omitempty"`
	// ActivationLockBypassCode Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
	ActivationLockBypassCode *string `json:"activationLockBypassCode,omitempty"`
	// EmailAddress Email(s) for the user associated with the device. This property is read-only.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// AzureActiveDirectoryDeviceID The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
	AzureActiveDirectoryDeviceID *string `json:"azureActiveDirectoryDeviceId,omitempty"`
	// AzureADDeviceID The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
	AzureADDeviceID *string `json:"azureADDeviceId,omitempty"`
	// DeviceRegistrationState Device registration state. This property is read-only.
	DeviceRegistrationState *DeviceRegistrationState `json:"deviceRegistrationState,omitempty"`
	// DeviceCategoryDisplayName Device category display name. This property is read-only.
	DeviceCategoryDisplayName *string `json:"deviceCategoryDisplayName,omitempty"`
	// IsSupervised Device supervised status. This property is read-only.
	IsSupervised *bool `json:"isSupervised,omitempty"`
	// ExchangeLastSuccessfulSyncDateTime Last time the device contacted Exchange. This property is read-only.
	ExchangeLastSuccessfulSyncDateTime *time.Time `json:"exchangeLastSuccessfulSyncDateTime,omitempty"`
	// ExchangeAccessState The Access State of the device in Exchange. This property is read-only.
	ExchangeAccessState *DeviceManagementExchangeAccessState `json:"exchangeAccessState,omitempty"`
	// ExchangeAccessStateReason The reason for the device's access state in Exchange. This property is read-only.
	ExchangeAccessStateReason *DeviceManagementExchangeAccessStateReason `json:"exchangeAccessStateReason,omitempty"`
	// RemoteAssistanceSessionURL Url that allows a Remote Assistance session to be established with the device. This property is read-only.
	RemoteAssistanceSessionURL *string `json:"remoteAssistanceSessionUrl,omitempty"`
	// RemoteAssistanceSessionErrorDetails An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
	RemoteAssistanceSessionErrorDetails *string `json:"remoteAssistanceSessionErrorDetails,omitempty"`
	// IsEncrypted Device encryption status. This property is read-only.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// UserPrincipalName Device user principal name. This property is read-only.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// Model Model of the device. This property is read-only.
	Model *string `json:"model,omitempty"`
	// Manufacturer Manufacturer of the device. This property is read-only.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Imei IMEI. This property is read-only.
	Imei *string `json:"imei,omitempty"`
	// ComplianceGracePeriodExpirationDateTime The DateTime when device compliance grace period expires. This property is read-only.
	ComplianceGracePeriodExpirationDateTime *time.Time `json:"complianceGracePeriodExpirationDateTime,omitempty"`
	// SerialNumber SerialNumber. This property is read-only.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// PhoneNumber Phone number of the device. This property is read-only.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// AndroidSecurityPatchLevel Android security patch level. This property is read-only.
	AndroidSecurityPatchLevel *string `json:"androidSecurityPatchLevel,omitempty"`
	// UserDisplayName User display name. This property is read-only.
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// ConfigurationManagerClientEnabledFeatures ConfigrMgr client enabled features. This property is read-only.
	ConfigurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures `json:"configurationManagerClientEnabledFeatures,omitempty"`
	// WiFiMacAddress Wi-Fi MAC. This property is read-only.
	WiFiMacAddress *string `json:"wiFiMacAddress,omitempty"`
	// DeviceHealthAttestationState The device health attestation state. This property is read-only.
	DeviceHealthAttestationState *DeviceHealthAttestationState `json:"deviceHealthAttestationState,omitempty"`
	// SubscriberCarrier Subscriber Carrier. This property is read-only.
	SubscriberCarrier *string `json:"subscriberCarrier,omitempty"`
	// Meid MEID. This property is read-only.
	Meid *string `json:"meid,omitempty"`
	// TotalStorageSpaceInBytes Total Storage in Bytes. This property is read-only.
	TotalStorageSpaceInBytes *int `json:"totalStorageSpaceInBytes,omitempty"`
	// FreeStorageSpaceInBytes Free Storage in Bytes. This property is read-only.
	FreeStorageSpaceInBytes *int `json:"freeStorageSpaceInBytes,omitempty"`
	// ManagedDeviceName Automatically generated name to identify a device. Can be overwritten to a user friendly name.
	ManagedDeviceName *string `json:"managedDeviceName,omitempty"`
	// PartnerReportedThreatState Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only.
	PartnerReportedThreatState *ManagedDevicePartnerReportedHealthState `json:"partnerReportedThreatState,omitempty"`
	// RetireAfterDateTime Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
	RetireAfterDateTime *time.Time `json:"retireAfterDateTime,omitempty"`
	// UsersLoggedOn Indicates the last logged on users of a device. This property is read-only.
	UsersLoggedOn []LoggedOnUser `json:"usersLoggedOn,omitempty"`
	// PreferMdmOverGroupPolicyAppliedDateTime Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
	PreferMdmOverGroupPolicyAppliedDateTime *time.Time `json:"preferMdmOverGroupPolicyAppliedDateTime,omitempty"`
	// AutopilotEnrolled Reports if the managed device is enrolled via auto-pilot. This property is read-only.
	AutopilotEnrolled *bool `json:"autopilotEnrolled,omitempty"`
	// RequireUserEnrollmentApproval Reports if the managed iOS device is user approval enrollment. This property is read-only.
	RequireUserEnrollmentApproval *bool `json:"requireUserEnrollmentApproval,omitempty"`
	// ManagementCertificateExpirationDate Reports device management certificate expiration date. This property is read-only.
	ManagementCertificateExpirationDate *time.Time `json:"managementCertificateExpirationDate,omitempty"`
	// Iccid Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
	Iccid *string `json:"iccid,omitempty"`
	// Udid Unique Device Identifier for iOS and macOS devices. This property is read-only.
	Udid *string `json:"udid,omitempty"`
	// RoleScopeTagIDs List of Scope Tag IDs for this Device instance.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// WindowsActiveMalwareCount Count of active malware for this windows device. This property is read-only.
	WindowsActiveMalwareCount *int `json:"windowsActiveMalwareCount,omitempty"`
	// WindowsRemediatedMalwareCount Count of remediated malware for this windows device. This property is read-only.
	WindowsRemediatedMalwareCount *int `json:"windowsRemediatedMalwareCount,omitempty"`
	// Notes Notes on the device created by IT Admin
	Notes *string `json:"notes,omitempty"`
	// ConfigurationManagerClientHealthState Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
	ConfigurationManagerClientHealthState *ConfigurationManagerClientHealthState `json:"configurationManagerClientHealthState,omitempty"`
	// ConfigurationManagerClientInformation Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
	ConfigurationManagerClientInformation *ConfigurationManagerClientInformation `json:"configurationManagerClientInformation,omitempty"`
	// EthernetMacAddress Ethernet MAC. This property is read-only.
	EthernetMacAddress *string `json:"ethernetMacAddress,omitempty"`
	// SecurityBaselineStates undocumented
	SecurityBaselineStates []SecurityBaselineState `json:"securityBaselineStates,omitempty"`
	// DeviceConfigurationStates undocumented
	DeviceConfigurationStates []DeviceConfigurationState `json:"deviceConfigurationStates,omitempty"`
	// DeviceCompliancePolicyStates undocumented
	DeviceCompliancePolicyStates []DeviceCompliancePolicyState `json:"deviceCompliancePolicyStates,omitempty"`
	// ManagedDeviceMobileAppConfigurationStates undocumented
	ManagedDeviceMobileAppConfigurationStates []ManagedDeviceMobileAppConfigurationState `json:"managedDeviceMobileAppConfigurationStates,omitempty"`
	// DetectedApps undocumented
	DetectedApps []DetectedApp `json:"detectedApps,omitempty"`
	// DeviceCategory undocumented
	DeviceCategory *DeviceCategory `json:"deviceCategory,omitempty"`
	// WindowsProtectionState undocumented
	WindowsProtectionState *WindowsProtectionState `json:"windowsProtectionState,omitempty"`
	// Users undocumented
	Users []User `json:"users,omitempty"`
}
