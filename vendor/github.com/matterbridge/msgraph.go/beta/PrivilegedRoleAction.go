// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PrivilegedRoleSelfActivateRequestParameter undocumented
type PrivilegedRoleSelfActivateRequestParameter struct {
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// Duration undocumented
	Duration *string `json:"duration,omitempty"`
	// TicketNumber undocumented
	TicketNumber *string `json:"ticketNumber,omitempty"`
	// TicketSystem undocumented
	TicketSystem *string `json:"ticketSystem,omitempty"`
}

// PrivilegedRoleSelfDeactivateRequestParameter undocumented
type PrivilegedRoleSelfDeactivateRequestParameter struct {
}

//
type PrivilegedRoleSelfActivateRequestBuilder struct{ BaseRequestBuilder }

// SelfActivate action undocumented
func (b *PrivilegedRoleRequestBuilder) SelfActivate(reqObj *PrivilegedRoleSelfActivateRequestParameter) *PrivilegedRoleSelfActivateRequestBuilder {
	bb := &PrivilegedRoleSelfActivateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/selfActivate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrivilegedRoleSelfActivateRequest struct{ BaseRequest }

//
func (b *PrivilegedRoleSelfActivateRequestBuilder) Request() *PrivilegedRoleSelfActivateRequest {
	return &PrivilegedRoleSelfActivateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrivilegedRoleSelfActivateRequest) Post(ctx context.Context) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type PrivilegedRoleSelfDeactivateRequestBuilder struct{ BaseRequestBuilder }

// SelfDeactivate action undocumented
func (b *PrivilegedRoleRequestBuilder) SelfDeactivate(reqObj *PrivilegedRoleSelfDeactivateRequestParameter) *PrivilegedRoleSelfDeactivateRequestBuilder {
	bb := &PrivilegedRoleSelfDeactivateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/selfDeactivate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrivilegedRoleSelfDeactivateRequest struct{ BaseRequest }

//
func (b *PrivilegedRoleSelfDeactivateRequestBuilder) Request() *PrivilegedRoleSelfDeactivateRequest {
	return &PrivilegedRoleSelfDeactivateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrivilegedRoleSelfDeactivateRequest) Post(ctx context.Context) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
