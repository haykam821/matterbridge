// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BitLockerRecoveryPasswordRotationType undocumented
type BitLockerRecoveryPasswordRotationType int

const (
	// BitLockerRecoveryPasswordRotationTypeVNotConfigured undocumented
	BitLockerRecoveryPasswordRotationTypeVNotConfigured BitLockerRecoveryPasswordRotationType = 0
	// BitLockerRecoveryPasswordRotationTypeVDisabled undocumented
	BitLockerRecoveryPasswordRotationTypeVDisabled BitLockerRecoveryPasswordRotationType = 1
	// BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAd undocumented
	BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAd BitLockerRecoveryPasswordRotationType = 2
	// BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAdAndHybrid undocumented
	BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAdAndHybrid BitLockerRecoveryPasswordRotationType = 3
)

// BitLockerRecoveryPasswordRotationTypePNotConfigured returns a pointer to BitLockerRecoveryPasswordRotationTypeVNotConfigured
func BitLockerRecoveryPasswordRotationTypePNotConfigured() *BitLockerRecoveryPasswordRotationType {
	v := BitLockerRecoveryPasswordRotationTypeVNotConfigured
	return &v
}

// BitLockerRecoveryPasswordRotationTypePDisabled returns a pointer to BitLockerRecoveryPasswordRotationTypeVDisabled
func BitLockerRecoveryPasswordRotationTypePDisabled() *BitLockerRecoveryPasswordRotationType {
	v := BitLockerRecoveryPasswordRotationTypeVDisabled
	return &v
}

// BitLockerRecoveryPasswordRotationTypePEnabledForAzureAd returns a pointer to BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAd
func BitLockerRecoveryPasswordRotationTypePEnabledForAzureAd() *BitLockerRecoveryPasswordRotationType {
	v := BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAd
	return &v
}

// BitLockerRecoveryPasswordRotationTypePEnabledForAzureAdAndHybrid returns a pointer to BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAdAndHybrid
func BitLockerRecoveryPasswordRotationTypePEnabledForAzureAdAndHybrid() *BitLockerRecoveryPasswordRotationType {
	v := BitLockerRecoveryPasswordRotationTypeVEnabledForAzureAdAndHybrid
	return &v
}
