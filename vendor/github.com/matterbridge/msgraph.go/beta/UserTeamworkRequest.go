// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// UserTeamworkRequestBuilder is request builder for UserTeamwork
type UserTeamworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTeamworkRequest
func (b *UserTeamworkRequestBuilder) Request() *UserTeamworkRequest {
	return &UserTeamworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTeamworkRequest is request for UserTeamwork
type UserTeamworkRequest struct{ BaseRequest }

// Get performs GET request for UserTeamwork
func (r *UserTeamworkRequest) Get(ctx context.Context) (resObj *UserTeamwork, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTeamwork
func (r *UserTeamworkRequest) Update(ctx context.Context, reqObj *UserTeamwork) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTeamwork
func (r *UserTeamworkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *UserTeamworkRequestBuilder) InstalledApps() *UserTeamworkInstalledAppsCollectionRequestBuilder {
	bb := &UserTeamworkInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// UserTeamworkInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
type UserTeamworkInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *UserTeamworkInstalledAppsCollectionRequestBuilder) Request() *UserTeamworkInstalledAppsCollectionRequest {
	return &UserTeamworkInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *UserTeamworkInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserTeamworkInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type UserTeamworkInstalledAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TeamsAppInstallation, error) {
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
	var values []TeamsAppInstallation
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
			value  []TeamsAppInstallation
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

// Get performs GET request for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Get(ctx context.Context) ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppInstallation) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
