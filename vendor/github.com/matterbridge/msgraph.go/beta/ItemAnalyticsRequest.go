// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ItemAnalyticsRequestBuilder is request builder for ItemAnalytics
type ItemAnalyticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemAnalyticsRequest
func (b *ItemAnalyticsRequestBuilder) Request() *ItemAnalyticsRequest {
	return &ItemAnalyticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemAnalyticsRequest is request for ItemAnalytics
type ItemAnalyticsRequest struct{ BaseRequest }

// Get performs GET request for ItemAnalytics
func (r *ItemAnalyticsRequest) Get(ctx context.Context) (resObj *ItemAnalytics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ItemAnalytics
func (r *ItemAnalyticsRequest) Update(ctx context.Context, reqObj *ItemAnalytics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ItemAnalytics
func (r *ItemAnalyticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AllTime is navigation property
func (b *ItemAnalyticsRequestBuilder) AllTime() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/allTime"
	return bb
}

// ItemActivityStats returns request builder for ItemActivityStat collection
func (b *ItemAnalyticsRequestBuilder) ItemActivityStats() *ItemAnalyticsItemActivityStatsCollectionRequestBuilder {
	bb := &ItemAnalyticsItemActivityStatsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/itemActivityStats"
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequestBuilder is request builder for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityStat collection
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) Request() *ItemAnalyticsItemActivityStatsCollectionRequest {
	return &ItemAnalyticsItemActivityStatsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityStat item
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) ID(id string) *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequest is request for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ItemActivityStat, error) {
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
	var values []ItemActivityStat
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
			value  []ItemActivityStat
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

// Get performs GET request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Get(ctx context.Context) ([]ItemActivityStat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Add(ctx context.Context, reqObj *ItemActivityStat) (resObj *ItemActivityStat, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LastSevenDays is navigation property
func (b *ItemAnalyticsRequestBuilder) LastSevenDays() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastSevenDays"
	return bb
}
