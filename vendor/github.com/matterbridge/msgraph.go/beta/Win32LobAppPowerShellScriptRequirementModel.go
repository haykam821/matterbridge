// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Win32LobAppPowerShellScriptRequirement undocumented
type Win32LobAppPowerShellScriptRequirement struct {
	// Win32LobAppRequirement is the base model of Win32LobAppPowerShellScriptRequirement
	Win32LobAppRequirement
	// DisplayName The unique display name for this rule
	DisplayName *string `json:"displayName,omitempty"`
	// EnforceSignatureCheck A value indicating whether signature check is enforced
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// RunAs32Bit A value indicating whether this script should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// RunAsAccount Indicates the type of execution context the script runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
	// ScriptContent The base64 encoded script content to detect Win32 Line of Business (LoB) app
	ScriptContent *string `json:"scriptContent,omitempty"`
	// DetectionType The detection type for script output
	DetectionType *Win32LobAppPowerShellScriptDetectionType `json:"detectionType,omitempty"`
}
