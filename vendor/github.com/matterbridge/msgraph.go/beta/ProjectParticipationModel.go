// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProjectParticipation undocumented
type ProjectParticipation struct {
	// ItemFacet is the base model of ProjectParticipation
	ItemFacet
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Client undocumented
	Client *CompanyDetail `json:"client,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Detail undocumented
	Detail *PositionDetail `json:"detail,omitempty"`
	// Colleagues undocumented
	Colleagues []RelatedPerson `json:"colleagues,omitempty"`
	// Sponsors undocumented
	Sponsors []RelatedPerson `json:"sponsors,omitempty"`
}
