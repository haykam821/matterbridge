// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// AccessPackageRequestBuilder is request builder for AccessPackage
type AccessPackageRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageRequest
func (b *AccessPackageRequestBuilder) Request() *AccessPackageRequest {
	return &AccessPackageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageRequest is request for AccessPackage
type AccessPackageRequest struct{ BaseRequest }

// Get performs GET request for AccessPackage
func (r *AccessPackageRequest) Get(ctx context.Context) (resObj *AccessPackage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackage
func (r *AccessPackageRequest) Update(ctx context.Context, reqObj *AccessPackage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackage
func (r *AccessPackageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageAssignmentPolicies returns request builder for AccessPackageAssignmentPolicy collection
func (b *AccessPackageRequestBuilder) AccessPackageAssignmentPolicies() *AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder {
	bb := &AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentPolicies"
	return bb
}

// AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder is request builder for AccessPackageAssignmentPolicy collection
type AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentPolicy collection
func (b *AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder) Request() *AccessPackageAccessPackageAssignmentPoliciesCollectionRequest {
	return &AccessPackageAccessPackageAssignmentPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentPolicy item
func (b *AccessPackageAccessPackageAssignmentPoliciesCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentPolicyRequestBuilder {
	bb := &AccessPackageAssignmentPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageAccessPackageAssignmentPoliciesCollectionRequest is request for AccessPackageAssignmentPolicy collection
type AccessPackageAccessPackageAssignmentPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentPolicy collection
func (r *AccessPackageAccessPackageAssignmentPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignmentPolicy, error) {
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
	var values []AccessPackageAssignmentPolicy
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
			value  []AccessPackageAssignmentPolicy
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

// Get performs GET request for AccessPackageAssignmentPolicy collection
func (r *AccessPackageAccessPackageAssignmentPoliciesCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentPolicy collection
func (r *AccessPackageAccessPackageAssignmentPoliciesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentPolicy) (resObj *AccessPackageAssignmentPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageCatalog is navigation property
func (b *AccessPackageRequestBuilder) AccessPackageCatalog() *AccessPackageCatalogRequestBuilder {
	bb := &AccessPackageCatalogRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageCatalog"
	return bb
}

// AccessPackageResourceRoleScopes returns request builder for AccessPackageResourceRoleScope collection
func (b *AccessPackageRequestBuilder) AccessPackageResourceRoleScopes() *AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder {
	bb := &AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRoleScopes"
	return bb
}

// AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder is request builder for AccessPackageResourceRoleScope collection
type AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceRoleScope collection
func (b *AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder) Request() *AccessPackageAccessPackageResourceRoleScopesCollectionRequest {
	return &AccessPackageAccessPackageResourceRoleScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceRoleScope item
func (b *AccessPackageAccessPackageResourceRoleScopesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRoleScopeRequestBuilder {
	bb := &AccessPackageResourceRoleScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageAccessPackageResourceRoleScopesCollectionRequest is request for AccessPackageResourceRoleScope collection
type AccessPackageAccessPackageResourceRoleScopesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResourceRoleScope collection
func (r *AccessPackageAccessPackageResourceRoleScopesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResourceRoleScope, error) {
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
	var values []AccessPackageResourceRoleScope
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
			value  []AccessPackageResourceRoleScope
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

// Get performs GET request for AccessPackageResourceRoleScope collection
func (r *AccessPackageAccessPackageResourceRoleScopesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResourceRoleScope, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResourceRoleScope collection
func (r *AccessPackageAccessPackageResourceRoleScopesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResourceRoleScope) (resObj *AccessPackageResourceRoleScope, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
