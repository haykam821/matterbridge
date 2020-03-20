// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SwapShiftsChangeRequestObjectRequestBuilder is request builder for SwapShiftsChangeRequestObject
type SwapShiftsChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns SwapShiftsChangeRequestObjectRequest
func (b *SwapShiftsChangeRequestObjectRequestBuilder) Request() *SwapShiftsChangeRequestObjectRequest {
	return &SwapShiftsChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SwapShiftsChangeRequestObjectRequest is request for SwapShiftsChangeRequestObject
type SwapShiftsChangeRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for SwapShiftsChangeRequestObject
func (r *SwapShiftsChangeRequestObjectRequest) Get(ctx context.Context) (resObj *SwapShiftsChangeRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SwapShiftsChangeRequestObject
func (r *SwapShiftsChangeRequestObjectRequest) Update(ctx context.Context, reqObj *SwapShiftsChangeRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SwapShiftsChangeRequestObject
func (r *SwapShiftsChangeRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
