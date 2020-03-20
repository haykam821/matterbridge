// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// PrivilegedRoleRequestBuilder is request builder for PrivilegedRole
type PrivilegedRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleRequest
func (b *PrivilegedRoleRequestBuilder) Request() *PrivilegedRoleRequest {
	return &PrivilegedRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleRequest is request for PrivilegedRole
type PrivilegedRoleRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRole
func (r *PrivilegedRoleRequest) Get(ctx context.Context) (resObj *PrivilegedRole, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRole
func (r *PrivilegedRoleRequest) Update(ctx context.Context, reqObj *PrivilegedRole) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRole
func (r *PrivilegedRoleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for PrivilegedRoleAssignment collection
func (b *PrivilegedRoleRequestBuilder) Assignments() *PrivilegedRoleAssignmentsCollectionRequestBuilder {
	bb := &PrivilegedRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// PrivilegedRoleAssignmentsCollectionRequestBuilder is request builder for PrivilegedRoleAssignment collection
type PrivilegedRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PrivilegedRoleAssignment collection
func (b *PrivilegedRoleAssignmentsCollectionRequestBuilder) Request() *PrivilegedRoleAssignmentsCollectionRequest {
	return &PrivilegedRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PrivilegedRoleAssignment item
func (b *PrivilegedRoleAssignmentsCollectionRequestBuilder) ID(id string) *PrivilegedRoleAssignmentRequestBuilder {
	bb := &PrivilegedRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedRoleAssignmentsCollectionRequest is request for PrivilegedRoleAssignment collection
type PrivilegedRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]PrivilegedRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PrivilegedRoleAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PrivilegedRoleAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]PrivilegedRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *PrivilegedRoleAssignment) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Settings is navigation property
func (b *PrivilegedRoleRequestBuilder) Settings() *PrivilegedRoleSettingsRequestBuilder {
	bb := &PrivilegedRoleSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/settings"
	return bb
}

// Summary is navigation property
func (b *PrivilegedRoleRequestBuilder) Summary() *PrivilegedRoleSummaryRequestBuilder {
	bb := &PrivilegedRoleSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/summary"
	return bb
}
