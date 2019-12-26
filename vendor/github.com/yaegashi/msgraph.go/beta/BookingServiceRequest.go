// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// BookingServiceRequestBuilder is request builder for BookingService
type BookingServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingServiceRequest
func (b *BookingServiceRequestBuilder) Request() *BookingServiceRequest {
	return &BookingServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingServiceRequest is request for BookingService
type BookingServiceRequest struct{ BaseRequest }

// Get performs GET request for BookingService
func (r *BookingServiceRequest) Get(ctx context.Context) (resObj *BookingService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingService
func (r *BookingServiceRequest) Update(ctx context.Context, reqObj *BookingService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingService
func (r *BookingServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
