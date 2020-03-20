// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSVpnSecurityAssociationParameters undocumented
type IOSVpnSecurityAssociationParameters struct {
	// Object is the base model of IOSVpnSecurityAssociationParameters
	Object
	// SecurityEncryptionAlgorithm Encryption algorithm
	SecurityEncryptionAlgorithm *VpnEncryptionAlgorithmType `json:"securityEncryptionAlgorithm,omitempty"`
	// SecurityIntegrityAlgorithm Integrity algorithm
	SecurityIntegrityAlgorithm *VpnIntegrityAlgorithmType `json:"securityIntegrityAlgorithm,omitempty"`
	// SecurityDiffieHellmanGroup Diffie-Hellman Group
	SecurityDiffieHellmanGroup *int `json:"securityDiffieHellmanGroup,omitempty"`
	// LifetimeInMinutes Lifetime (minutes)
	LifetimeInMinutes *int `json:"lifetimeInMinutes,omitempty"`
}
