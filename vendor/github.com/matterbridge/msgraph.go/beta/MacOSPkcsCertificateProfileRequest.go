// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// MacOSPkcsCertificateProfileRequestBuilder is request builder for MacOSPkcsCertificateProfile
type MacOSPkcsCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSPkcsCertificateProfileRequest
func (b *MacOSPkcsCertificateProfileRequestBuilder) Request() *MacOSPkcsCertificateProfileRequest {
	return &MacOSPkcsCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSPkcsCertificateProfileRequest is request for MacOSPkcsCertificateProfile
type MacOSPkcsCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Get(ctx context.Context) (resObj *MacOSPkcsCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Update(ctx context.Context, reqObj *MacOSPkcsCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSPkcsCertificateProfile
func (r *MacOSPkcsCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *MacOSPkcsCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
	var values []ManagedDeviceCertificateState
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
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *MacOSPkcsCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
