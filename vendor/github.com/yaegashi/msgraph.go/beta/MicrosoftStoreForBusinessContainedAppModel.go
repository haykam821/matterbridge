// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MicrosoftStoreForBusinessContainedApp A class that represents a contained app of a MicrosoftStoreForBusinessApp.
type MicrosoftStoreForBusinessContainedApp struct {
	// MobileContainedApp is the base model of MicrosoftStoreForBusinessContainedApp
	MobileContainedApp
	// AppUserModelID The app user model ID of the contained app of a MicrosoftStoreForBusinessApp.
	AppUserModelID *string `json:"appUserModelId,omitempty"`
}
