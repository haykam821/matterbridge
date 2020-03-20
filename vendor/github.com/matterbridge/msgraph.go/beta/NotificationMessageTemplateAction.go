// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// NotificationMessageTemplateSendTestMessageRequestParameter undocumented
type NotificationMessageTemplateSendTestMessageRequestParameter struct {
}

//
type NotificationMessageTemplateSendTestMessageRequestBuilder struct{ BaseRequestBuilder }

// SendTestMessage action undocumented
func (b *NotificationMessageTemplateRequestBuilder) SendTestMessage(reqObj *NotificationMessageTemplateSendTestMessageRequestParameter) *NotificationMessageTemplateSendTestMessageRequestBuilder {
	bb := &NotificationMessageTemplateSendTestMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendTestMessage"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type NotificationMessageTemplateSendTestMessageRequest struct{ BaseRequest }

//
func (b *NotificationMessageTemplateSendTestMessageRequestBuilder) Request() *NotificationMessageTemplateSendTestMessageRequest {
	return &NotificationMessageTemplateSendTestMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *NotificationMessageTemplateSendTestMessageRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
