// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OrganizationalBranding undocumented
type OrganizationalBranding struct {
	// Entity is the base model of OrganizationalBranding
	Entity
	// BackgroundColor undocumented
	BackgroundColor *string `json:"backgroundColor,omitempty"`
	// BackgroundImage undocumented
	BackgroundImage *Stream `json:"backgroundImage,omitempty"`
	// BannerLogo undocumented
	BannerLogo *Stream `json:"bannerLogo,omitempty"`
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
	// SignInPageText undocumented
	SignInPageText *string `json:"signInPageText,omitempty"`
	// SquareLogo undocumented
	SquareLogo *Stream `json:"squareLogo,omitempty"`
	// UsernameHintText undocumented
	UsernameHintText *string `json:"usernameHintText,omitempty"`
}
