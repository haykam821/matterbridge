// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// OnPremisesAgentRequestBuilder is request builder for OnPremisesAgent
type OnPremisesAgentRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesAgentRequest
func (b *OnPremisesAgentRequestBuilder) Request() *OnPremisesAgentRequest {
	return &OnPremisesAgentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesAgentRequest is request for OnPremisesAgent
type OnPremisesAgentRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Get(ctx context.Context) (resObj *OnPremisesAgent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Update(ctx context.Context, reqObj *OnPremisesAgent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesAgent
func (r *OnPremisesAgentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *OnPremisesAgentRequestBuilder) AgentGroups() *OnPremisesAgentAgentGroupsCollectionRequestBuilder {
	bb := &OnPremisesAgentAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection
type OnPremisesAgentAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) Request() *OnPremisesAgentAgentGroupsCollectionRequest {
	return &OnPremisesAgentAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *OnPremisesAgentAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnPremisesAgentAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type OnPremisesAgentAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnPremisesAgentGroup, error) {
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
	var values []OnPremisesAgentGroup
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
			value  []OnPremisesAgentGroup
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

// Get performs GET request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Get(ctx context.Context) ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *OnPremisesAgentAgentGroupsCollectionRequest) Add(ctx context.Context, reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
