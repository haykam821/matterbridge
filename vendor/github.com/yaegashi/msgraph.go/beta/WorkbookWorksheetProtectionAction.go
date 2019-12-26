// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookWorksheetProtectionProtectRequestParameter undocumented
type WorkbookWorksheetProtectionProtectRequestParameter struct {
	// Options undocumented
	Options *WorkbookWorksheetProtectionOptions `json:"options,omitempty"`
}

// WorkbookWorksheetProtectionUnprotectRequestParameter undocumented
type WorkbookWorksheetProtectionUnprotectRequestParameter struct {
}

//
type WorkbookWorksheetProtectionProtectRequestBuilder struct{ BaseRequestBuilder }

// Protect action undocumented
func (b *WorkbookWorksheetProtectionRequestBuilder) Protect(reqObj *WorkbookWorksheetProtectionProtectRequestParameter) *WorkbookWorksheetProtectionProtectRequestBuilder {
	bb := &WorkbookWorksheetProtectionProtectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/protect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookWorksheetProtectionProtectRequest struct{ BaseRequest }

//
func (b *WorkbookWorksheetProtectionProtectRequestBuilder) Request() *WorkbookWorksheetProtectionProtectRequest {
	return &WorkbookWorksheetProtectionProtectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookWorksheetProtectionProtectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type WorkbookWorksheetProtectionUnprotectRequestBuilder struct{ BaseRequestBuilder }

// Unprotect action undocumented
func (b *WorkbookWorksheetProtectionRequestBuilder) Unprotect(reqObj *WorkbookWorksheetProtectionUnprotectRequestParameter) *WorkbookWorksheetProtectionUnprotectRequestBuilder {
	bb := &WorkbookWorksheetProtectionUnprotectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unprotect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookWorksheetProtectionUnprotectRequest struct{ BaseRequest }

//
func (b *WorkbookWorksheetProtectionUnprotectRequestBuilder) Request() *WorkbookWorksheetProtectionUnprotectRequest {
	return &WorkbookWorksheetProtectionUnprotectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookWorksheetProtectionUnprotectRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
