// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RunAsAccountType undocumented
type RunAsAccountType int

const (
	// RunAsAccountTypeVSystem undocumented
	RunAsAccountTypeVSystem RunAsAccountType = 0
	// RunAsAccountTypeVUser undocumented
	RunAsAccountTypeVUser RunAsAccountType = 1
)

// RunAsAccountTypePSystem returns a pointer to RunAsAccountTypeVSystem
func RunAsAccountTypePSystem() *RunAsAccountType {
	v := RunAsAccountTypeVSystem
	return &v
}

// RunAsAccountTypePUser returns a pointer to RunAsAccountTypeVUser
func RunAsAccountTypePUser() *RunAsAccountType {
	v := RunAsAccountTypeVUser
	return &v
}
