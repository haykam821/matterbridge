// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DepTokenType undocumented
type DepTokenType int

const (
	// DepTokenTypeVNone undocumented
	DepTokenTypeVNone DepTokenType = 0
	// DepTokenTypeVDep undocumented
	DepTokenTypeVDep DepTokenType = 1
	// DepTokenTypeVAppleSchoolManager undocumented
	DepTokenTypeVAppleSchoolManager DepTokenType = 2
)

// DepTokenTypePNone returns a pointer to DepTokenTypeVNone
func DepTokenTypePNone() *DepTokenType {
	v := DepTokenTypeVNone
	return &v
}

// DepTokenTypePDep returns a pointer to DepTokenTypeVDep
func DepTokenTypePDep() *DepTokenType {
	v := DepTokenTypeVDep
	return &v
}

// DepTokenTypePAppleSchoolManager returns a pointer to DepTokenTypeVAppleSchoolManager
func DepTokenTypePAppleSchoolManager() *DepTokenType {
	v := DepTokenTypeVAppleSchoolManager
	return &v
}