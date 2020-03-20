// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartFillClearRequestParameter undocumented
type WorkbookChartFillClearRequestParameter struct {
}

// WorkbookChartFillSetSolidColorRequestParameter undocumented
type WorkbookChartFillSetSolidColorRequestParameter struct {
	// Color undocumented
	Color *string `json:"color,omitempty"`
}

//
type WorkbookChartFillClearRequestBuilder struct{ BaseRequestBuilder }

// Clear action undocumented
func (b *WorkbookChartFillRequestBuilder) Clear(reqObj *WorkbookChartFillClearRequestParameter) *WorkbookChartFillClearRequestBuilder {
	bb := &WorkbookChartFillClearRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clear"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookChartFillClearRequest struct{ BaseRequest }

//
func (b *WorkbookChartFillClearRequestBuilder) Request() *WorkbookChartFillClearRequest {
	return &WorkbookChartFillClearRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookChartFillClearRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type WorkbookChartFillSetSolidColorRequestBuilder struct{ BaseRequestBuilder }

// SetSolidColor action undocumented
func (b *WorkbookChartFillRequestBuilder) SetSolidColor(reqObj *WorkbookChartFillSetSolidColorRequestParameter) *WorkbookChartFillSetSolidColorRequestBuilder {
	bb := &WorkbookChartFillSetSolidColorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setSolidColor"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookChartFillSetSolidColorRequest struct{ BaseRequest }

//
func (b *WorkbookChartFillSetSolidColorRequestBuilder) Request() *WorkbookChartFillSetSolidColorRequest {
	return &WorkbookChartFillSetSolidColorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookChartFillSetSolidColorRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
