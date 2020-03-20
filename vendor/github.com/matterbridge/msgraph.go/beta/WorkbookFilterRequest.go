// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookFilterRequestBuilder is request builder for WorkbookFilter
type WorkbookFilterRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookFilterRequest
func (b *WorkbookFilterRequestBuilder) Request() *WorkbookFilterRequest {
	return &WorkbookFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookFilterRequest is request for WorkbookFilter
type WorkbookFilterRequest struct{ BaseRequest }

// Get performs GET request for WorkbookFilter
func (r *WorkbookFilterRequest) Get(ctx context.Context) (resObj *WorkbookFilter, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookFilter
func (r *WorkbookFilterRequest) Update(ctx context.Context, reqObj *WorkbookFilter) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookFilter
func (r *WorkbookFilterRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
