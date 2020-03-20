// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidCertificateProfileBase Android certificate profile base.
type AndroidCertificateProfileBase struct {
	// DeviceConfiguration is the base model of AndroidCertificateProfileBase
	DeviceConfiguration
	// RenewalThresholdPercentage Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int `json:"renewalThresholdPercentage,omitempty"`
	// SubjectNameFormat Certificate Subject Name Format.
	SubjectNameFormat *SubjectNameFormat `json:"subjectNameFormat,omitempty"`
	// SubjectAlternativeNameType Certificate Subject Alternative Name Type.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`
	// CertificateValidityPeriodValue Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int `json:"certificateValidityPeriodValue,omitempty"`
	// CertificateValidityPeriodScale Scale for the Certificate Validity Period.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	// ExtendedKeyUsages Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages []ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`
	// RootCertificate undocumented
	RootCertificate *AndroidTrustedRootCertificate `json:"rootCertificate,omitempty"`
}
