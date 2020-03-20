// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// JournalPostRequestParameter undocumented
type JournalPostRequestParameter struct {
}

//
type JournalPostRequestBuilder struct{ BaseRequestBuilder }

// Post action undocumented
func (b *JournalRequestBuilder) Post(reqObj *JournalPostRequestParameter) *JournalPostRequestBuilder {
	bb := &JournalPostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/post"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type JournalPostRequest struct{ BaseRequest }

//
func (b *JournalPostRequestBuilder) Request() *JournalPostRequest {
	return &JournalPostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *JournalPostRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
