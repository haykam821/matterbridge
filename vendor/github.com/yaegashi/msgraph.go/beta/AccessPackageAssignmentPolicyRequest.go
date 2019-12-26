// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageAssignmentPolicyRequestBuilder is request builder for AccessPackageAssignmentPolicy
type AccessPackageAssignmentPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageAssignmentPolicyRequest
func (b *AccessPackageAssignmentPolicyRequestBuilder) Request() *AccessPackageAssignmentPolicyRequest {
	return &AccessPackageAssignmentPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageAssignmentPolicyRequest is request for AccessPackageAssignmentPolicy
type AccessPackageAssignmentPolicyRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageAssignmentPolicy
func (r *AccessPackageAssignmentPolicyRequest) Get(ctx context.Context) (resObj *AccessPackageAssignmentPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageAssignmentPolicy
func (r *AccessPackageAssignmentPolicyRequest) Update(ctx context.Context, reqObj *AccessPackageAssignmentPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageAssignmentPolicy
func (r *AccessPackageAssignmentPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackage is navigation property
func (b *AccessPackageAssignmentPolicyRequestBuilder) AccessPackage() *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackage"
	return bb
}

// AccessPackageCatalog is navigation property
func (b *AccessPackageAssignmentPolicyRequestBuilder) AccessPackageCatalog() *AccessPackageCatalogRequestBuilder {
	bb := &AccessPackageCatalogRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageCatalog"
	return bb
}
