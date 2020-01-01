// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsFeatureUpdateProfileRequestBuilder is request builder for WindowsFeatureUpdateProfile
type WindowsFeatureUpdateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsFeatureUpdateProfileRequest
func (b *WindowsFeatureUpdateProfileRequestBuilder) Request() *WindowsFeatureUpdateProfileRequest {
	return &WindowsFeatureUpdateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsFeatureUpdateProfileRequest is request for WindowsFeatureUpdateProfile
type WindowsFeatureUpdateProfileRequest struct{ BaseRequest }

// Get performs GET request for WindowsFeatureUpdateProfile
func (r *WindowsFeatureUpdateProfileRequest) Get(ctx context.Context) (resObj *WindowsFeatureUpdateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsFeatureUpdateProfile
func (r *WindowsFeatureUpdateProfileRequest) Update(ctx context.Context, reqObj *WindowsFeatureUpdateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsFeatureUpdateProfile
func (r *WindowsFeatureUpdateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for WindowsFeatureUpdateProfileAssignment collection
func (b *WindowsFeatureUpdateProfileRequestBuilder) Assignments() *WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder {
	bb := &WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder is request builder for WindowsFeatureUpdateProfileAssignment collection
type WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsFeatureUpdateProfileAssignment collection
func (b *WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder) Request() *WindowsFeatureUpdateProfileAssignmentsCollectionRequest {
	return &WindowsFeatureUpdateProfileAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsFeatureUpdateProfileAssignment item
func (b *WindowsFeatureUpdateProfileAssignmentsCollectionRequestBuilder) ID(id string) *WindowsFeatureUpdateProfileAssignmentRequestBuilder {
	bb := &WindowsFeatureUpdateProfileAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsFeatureUpdateProfileAssignmentsCollectionRequest is request for WindowsFeatureUpdateProfileAssignment collection
type WindowsFeatureUpdateProfileAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsFeatureUpdateProfileAssignment collection
func (r *WindowsFeatureUpdateProfileAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WindowsFeatureUpdateProfileAssignment, error) {
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
	var values []WindowsFeatureUpdateProfileAssignment
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
			value  []WindowsFeatureUpdateProfileAssignment
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

// Get performs GET request for WindowsFeatureUpdateProfileAssignment collection
func (r *WindowsFeatureUpdateProfileAssignmentsCollectionRequest) Get(ctx context.Context) ([]WindowsFeatureUpdateProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WindowsFeatureUpdateProfileAssignment collection
func (r *WindowsFeatureUpdateProfileAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *WindowsFeatureUpdateProfileAssignment) (resObj *WindowsFeatureUpdateProfileAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceUpdateStates returns request builder for WindowsUpdateState collection
func (b *WindowsFeatureUpdateProfileRequestBuilder) DeviceUpdateStates() *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder {
	bb := &WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceUpdateStates"
	return bb
}

// WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder is request builder for WindowsUpdateState collection
type WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsUpdateState collection
func (b *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder) Request() *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest {
	return &WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsUpdateState item
func (b *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequestBuilder) ID(id string) *WindowsUpdateStateRequestBuilder {
	bb := &WindowsUpdateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest is request for WindowsUpdateState collection
type WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsUpdateState collection
func (r *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WindowsUpdateState, error) {
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
	var values []WindowsUpdateState
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
			value  []WindowsUpdateState
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

// Get performs GET request for WindowsUpdateState collection
func (r *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest) Get(ctx context.Context) ([]WindowsUpdateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WindowsUpdateState collection
func (r *WindowsFeatureUpdateProfileDeviceUpdateStatesCollectionRequest) Add(ctx context.Context, reqObj *WindowsUpdateState) (resObj *WindowsUpdateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}