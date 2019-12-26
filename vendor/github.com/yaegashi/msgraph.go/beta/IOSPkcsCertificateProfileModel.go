// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSPkcsCertificateProfile iOS PKCS certificate profile.
type IOSPkcsCertificateProfile struct {
	// IOSCertificateProfileBase is the base model of IOSPkcsCertificateProfile
	IOSCertificateProfileBase
	// CertificationAuthority PKCS Certification Authority.
	CertificationAuthority *string `json:"certificationAuthority,omitempty"`
	// CertificationAuthorityName PKCS Certification Authority Name.
	CertificationAuthorityName *string `json:"certificationAuthorityName,omitempty"`
	// CertificateTemplateName PKCS Certificate Template Name.
	CertificateTemplateName *string `json:"certificateTemplateName,omitempty"`
	// SubjectAlternativeNameFormatString Custom String that defines the AAD Attribute.
	SubjectAlternativeNameFormatString *string `json:"subjectAlternativeNameFormatString,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}
