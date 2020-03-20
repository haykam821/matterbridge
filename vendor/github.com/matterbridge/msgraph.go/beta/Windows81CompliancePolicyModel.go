// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows81CompliancePolicy This class contains compliance settings for Windows 8.1.
type Windows81CompliancePolicy struct {
	// DeviceCompliancePolicy is the base model of Windows81CompliancePolicy
	DeviceCompliancePolicy
	// PasswordRequired Require a password to unlock Windows device.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// PasswordBlockSimple Indicates whether or not to block simple password.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays Password expiration in days.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength The minimum password length.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock Minutes of inactivity before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordMinimumCharacterSetCount The number of character sets required in the password.
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordRequiredType The required password type.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordPreviousPasswordBlockCount The number of previous passwords to prevent re-use of. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// OsMinimumVersion Minimum Windows 8.1 version.
	OsMinimumVersion *string `json:"osMinimumVersion,omitempty"`
	// OsMaximumVersion Maximum Windows 8.1 version.
	OsMaximumVersion *string `json:"osMaximumVersion,omitempty"`
	// StorageRequireEncryption Indicates whether or not to require encryption on a windows 8.1 device.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
}
