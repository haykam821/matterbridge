// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationUser undocumented
type EducationUser struct {
	// Entity is the base model of EducationUser
	Entity
	// RelatedContacts undocumented
	RelatedContacts []RelatedContact `json:"relatedContacts,omitempty"`
	// PrimaryRole undocumented
	PrimaryRole *EducationUserRole `json:"primaryRole,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ResidenceAddress undocumented
	ResidenceAddress *PhysicalAddress `json:"residenceAddress,omitempty"`
	// MailingAddress undocumented
	MailingAddress *PhysicalAddress `json:"mailingAddress,omitempty"`
	// Student undocumented
	Student *EducationStudent `json:"student,omitempty"`
	// Teacher undocumented
	Teacher *EducationTeacher `json:"teacher,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// RefreshTokensValidFromDateTime undocumented
	RefreshTokensValidFromDateTime *time.Time `json:"refreshTokensValidFromDateTime,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// OnPremisesInfo undocumented
	OnPremisesInfo *EducationOnPremisesInfo `json:"onPremisesInfo,omitempty"`
	// Assignments undocumented
	Assignments []EducationAssignment `json:"assignments,omitempty"`
	// Rubrics undocumented
	Rubrics []EducationRubric `json:"rubrics,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// TaughtClasses undocumented
	TaughtClasses []EducationClass `json:"taughtClasses,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}
