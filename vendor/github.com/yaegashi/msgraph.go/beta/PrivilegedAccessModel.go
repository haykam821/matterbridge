// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PrivilegedAccess undocumented
type PrivilegedAccess struct {
	// Entity is the base model of PrivilegedAccess
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Resources undocumented
	Resources []GovernanceResource `json:"resources,omitempty"`
	// RoleDefinitions undocumented
	RoleDefinitions []GovernanceRoleDefinition `json:"roleDefinitions,omitempty"`
	// RoleAssignments undocumented
	RoleAssignments []GovernanceRoleAssignment `json:"roleAssignments,omitempty"`
	// RoleAssignmentRequests undocumented
	RoleAssignmentRequests []GovernanceRoleAssignmentRequestObject `json:"roleAssignmentRequests,omitempty"`
	// RoleSettings undocumented
	RoleSettings []GovernanceRoleSetting `json:"roleSettings,omitempty"`
}
