// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceEnrollmentWindowsHelloForBusinessConfiguration Windows Hello for Business settings lets users access their devices using a gesture, such as biometric authentication, or a PIN. Configure settings for enrolled Windows 10, Windows 10 Mobile and later.
type DeviceEnrollmentWindowsHelloForBusinessConfiguration struct {
	// DeviceEnrollmentConfiguration is the base model of DeviceEnrollmentWindowsHelloForBusinessConfiguration
	DeviceEnrollmentConfiguration
	// PinMinimumLength Controls the minimum number of characters required for the Windows Hello for Business PIN.  This value must be between 4 and 127, inclusive, and less than or equal to the value set for the maximum PIN.
	PinMinimumLength *int `json:"pinMinimumLength,omitempty"`
	// PinMaximumLength Controls the maximum number of characters allowed for the Windows Hello for Business PIN. This value must be between 4 and 127, inclusive. This value must be greater than or equal to the value set for the minimum PIN.
	PinMaximumLength *int `json:"pinMaximumLength,omitempty"`
	// PinUppercaseCharactersUsage Controls the ability to use uppercase letters in the Windows Hello for Business PIN.  Allowed permits the use of uppercase letter(s), whereas Required ensures they are present. If set to Not Allowed, uppercase letters will not be permitted.
	PinUppercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinUppercaseCharactersUsage,omitempty"`
	// PinLowercaseCharactersUsage Controls the ability to use lowercase letters in the Windows Hello for Business PIN.  Allowed permits the use of lowercase letter(s), whereas Required ensures they are present. If set to Not Allowed, lowercase letters will not be permitted.
	PinLowercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinLowercaseCharactersUsage,omitempty"`
	// PinSpecialCharactersUsage Controls the ability to use special characters in the Windows Hello for Business PIN.  Allowed permits the use of special character(s), whereas Required ensures they are present. If set to Not Allowed, special character(s) will not be permitted.
	PinSpecialCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinSpecialCharactersUsage,omitempty"`
	// State Controls whether to allow the device to be configured for Windows Hello for Business. If set to disabled, the user cannot provision Windows Hello for Business except on Azure Active Directory joined mobile phones if otherwise required. If set to Not Configured, Intune will not override client defaults.
	State *Enablement `json:"state,omitempty"`
	// SecurityDeviceRequired Controls whether to require a Trusted Platform Module (TPM) for provisioning Windows Hello for Business. A TPM provides an additional security benefit in that data stored on it cannot be used on other devices. If set to False, all devices can provision Windows Hello for Business even if there is not a usable TPM.
	SecurityDeviceRequired *bool `json:"securityDeviceRequired,omitempty"`
	// UnlockWithBiometricsEnabled Controls the use of biometric gestures, such as face and fingerprint, as an alternative to the Windows Hello for Business PIN.  If set to False, biometric gestures are not allowed. Users must still configure a PIN as a backup in case of failures.
	UnlockWithBiometricsEnabled *bool `json:"unlockWithBiometricsEnabled,omitempty"`
	// RemotePassportEnabled Controls the use of Remote Windows Hello for Business. Remote Windows Hello for Business provides the ability for a portable, registered device to be usable as a companion for desktop authentication. The desktop must be Azure AD joined and the companion device must have a Windows Hello for Business PIN.
	RemotePassportEnabled *bool `json:"remotePassportEnabled,omitempty"`
	// PinPreviousBlockCount Controls the ability to prevent users from using past PINs. This must be set between 0 and 50, inclusive, and the current PIN of the user is included in that count. If set to 0, previous PINs are not stored. PIN history is not preserved through a PIN reset.
	PinPreviousBlockCount *int `json:"pinPreviousBlockCount,omitempty"`
	// PinExpirationInDays Controls the period of time (in days) that a PIN can be used before the system requires the user to change it. This must be set between 0 and 730, inclusive. If set to 0, the user's PIN will never expire
	PinExpirationInDays *int `json:"pinExpirationInDays,omitempty"`
	// EnhancedBiometricsState Controls the ability to use the anti-spoofing features for facial recognition on devices which support it. If set to disabled, anti-spoofing features are not allowed. If set to Not Configured, the user can choose whether they want to use anti-spoofing.
	EnhancedBiometricsState *Enablement `json:"enhancedBiometricsState,omitempty"`
	// SecurityKeyForSignIn Security key for Sign In provides the capacity for remotely turning ON/OFF Windows Hello Sercurity Keyl Not configured will honor configurations done on the clinet.
	SecurityKeyForSignIn *Enablement `json:"securityKeyForSignIn,omitempty"`
}
