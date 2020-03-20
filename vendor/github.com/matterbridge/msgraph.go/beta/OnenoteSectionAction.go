// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// OnenoteSectionCopyToNotebookRequestParameter undocumented
type OnenoteSectionCopyToNotebookRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// RenameAs undocumented
	RenameAs *string `json:"renameAs,omitempty"`
	// SiteCollectionID undocumented
	SiteCollectionID *string `json:"siteCollectionId,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

// OnenoteSectionCopyToSectionGroupRequestParameter undocumented
type OnenoteSectionCopyToSectionGroupRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// RenameAs undocumented
	RenameAs *string `json:"renameAs,omitempty"`
	// SiteCollectionID undocumented
	SiteCollectionID *string `json:"siteCollectionId,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}

//
type OnenoteSectionCopyToNotebookRequestBuilder struct{ BaseRequestBuilder }

// CopyToNotebook action undocumented
func (b *OnenoteSectionRequestBuilder) CopyToNotebook(reqObj *OnenoteSectionCopyToNotebookRequestParameter) *OnenoteSectionCopyToNotebookRequestBuilder {
	bb := &OnenoteSectionCopyToNotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copyToNotebook"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenoteSectionCopyToNotebookRequest struct{ BaseRequest }

//
func (b *OnenoteSectionCopyToNotebookRequestBuilder) Request() *OnenoteSectionCopyToNotebookRequest {
	return &OnenoteSectionCopyToNotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenoteSectionCopyToNotebookRequest) Post(ctx context.Context) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type OnenoteSectionCopyToSectionGroupRequestBuilder struct{ BaseRequestBuilder }

// CopyToSectionGroup action undocumented
func (b *OnenoteSectionRequestBuilder) CopyToSectionGroup(reqObj *OnenoteSectionCopyToSectionGroupRequestParameter) *OnenoteSectionCopyToSectionGroupRequestBuilder {
	bb := &OnenoteSectionCopyToSectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/copyToSectionGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OnenoteSectionCopyToSectionGroupRequest struct{ BaseRequest }

//
func (b *OnenoteSectionCopyToSectionGroupRequestBuilder) Request() *OnenoteSectionCopyToSectionGroupRequest {
	return &OnenoteSectionCopyToSectionGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OnenoteSectionCopyToSectionGroupRequest) Post(ctx context.Context) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
