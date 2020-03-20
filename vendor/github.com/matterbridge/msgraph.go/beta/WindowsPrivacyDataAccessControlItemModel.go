// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPrivacyDataAccessControlItem Specify access control level per privacy data category
type WindowsPrivacyDataAccessControlItem struct {
	// Entity is the base model of WindowsPrivacyDataAccessControlItem
	Entity
	// AccessLevel This indicates an access level for the privacy data category to which the specified application will be given to.
	AccessLevel *WindowsPrivacyDataAccessLevel `json:"accessLevel,omitempty"`
	// DataCategory This indicates a privacy data category to which the specific access control will apply.
	DataCategory *WindowsPrivacyDataCategory `json:"dataCategory,omitempty"`
	// AppPackageFamilyName The Package Family Name of a Windows app. When set, the access level applies to the specified application.
	AppPackageFamilyName *string `json:"appPackageFamilyName,omitempty"`
	// AppDisplayName The Package Family Name of a Windows app. When set, the access level applies to the specified application.
	AppDisplayName *string `json:"appDisplayName,omitempty"`
}
