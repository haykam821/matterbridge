// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// ManagementConditionStatementRequestBuilder is request builder for ManagementConditionStatement
type ManagementConditionStatementRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagementConditionStatementRequest
func (b *ManagementConditionStatementRequestBuilder) Request() *ManagementConditionStatementRequest {
	return &ManagementConditionStatementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagementConditionStatementRequest is request for ManagementConditionStatement
type ManagementConditionStatementRequest struct{ BaseRequest }

// Get performs GET request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Get(ctx context.Context) (resObj *ManagementConditionStatement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Update(ctx context.Context, reqObj *ManagementConditionStatement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagementConditionStatement
func (r *ManagementConditionStatementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagementConditions returns request builder for ManagementCondition collection
func (b *ManagementConditionStatementRequestBuilder) ManagementConditions() *ManagementConditionStatementManagementConditionsCollectionRequestBuilder {
	bb := &ManagementConditionStatementManagementConditionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managementConditions"
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequestBuilder is request builder for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagementCondition collection
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) Request() *ManagementConditionStatementManagementConditionsCollectionRequest {
	return &ManagementConditionStatementManagementConditionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagementCondition item
func (b *ManagementConditionStatementManagementConditionsCollectionRequestBuilder) ID(id string) *ManagementConditionRequestBuilder {
	bb := &ManagementConditionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagementConditionStatementManagementConditionsCollectionRequest is request for ManagementCondition collection
type ManagementConditionStatementManagementConditionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagementCondition, error) {
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
	var values []ManagementCondition
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
			value  []ManagementCondition
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

// Get performs GET request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Get(ctx context.Context) ([]ManagementCondition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagementCondition collection
func (r *ManagementConditionStatementManagementConditionsCollectionRequest) Add(ctx context.Context, reqObj *ManagementCondition) (resObj *ManagementCondition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
