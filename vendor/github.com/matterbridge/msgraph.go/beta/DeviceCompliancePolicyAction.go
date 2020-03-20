// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// DeviceCompliancePolicyCollectionHasPayloadLinksRequestParameter undocumented
type DeviceCompliancePolicyCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestParameter undocumented
type DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestParameter struct {
}

// DeviceCompliancePolicyAssignRequestParameter undocumented
type DeviceCompliancePolicyAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []DeviceCompliancePolicyAssignment `json:"assignments,omitempty"`
}

// DeviceCompliancePolicyScheduleActionsForRulesRequestParameter undocumented
type DeviceCompliancePolicyScheduleActionsForRulesRequestParameter struct {
	// DeviceComplianceScheduledActionForRules undocumented
	DeviceComplianceScheduledActionForRules []DeviceComplianceScheduledActionForRule `json:"deviceComplianceScheduledActionForRules,omitempty"`
}

//
type DeviceCompliancePolicyCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceManagementDeviceCompliancePoliciesCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceCompliancePolicyCollectionHasPayloadLinksRequestParameter) *DeviceCompliancePolicyCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceCompliancePolicyCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyCollectionHasPayloadLinksRequestBuilder) Request() *DeviceCompliancePolicyCollectionHasPayloadLinksRequest {
	return &DeviceCompliancePolicyCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
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
	var values [][]HasPayloadLinkResultItem
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
			value  [][]HasPayloadLinkResultItem
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

//
func (r *DeviceCompliancePolicyCollectionHasPayloadLinksRequest) Get(ctx context.Context) ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestBuilder struct{ BaseRequestBuilder }

// RefreshDeviceComplianceReportSummarization action undocumented
func (b *DeviceManagementDeviceCompliancePoliciesCollectionRequestBuilder) RefreshDeviceComplianceReportSummarization(reqObj *DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestParameter) *DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestBuilder {
	bb := &DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/refreshDeviceComplianceReportSummarization"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequestBuilder) Request() *DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequest {
	return &DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyCollectionRefreshDeviceComplianceReportSummarizationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type DeviceCompliancePolicyAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceCompliancePolicyRequestBuilder) Assign(reqObj *DeviceCompliancePolicyAssignRequestParameter) *DeviceCompliancePolicyAssignRequestBuilder {
	bb := &DeviceCompliancePolicyAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyAssignRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyAssignRequestBuilder) Request() *DeviceCompliancePolicyAssignRequest {
	return &DeviceCompliancePolicyAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]DeviceCompliancePolicyAssignment, error) {
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
	var values [][]DeviceCompliancePolicyAssignment
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
			value  [][]DeviceCompliancePolicyAssignment
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

//
func (r *DeviceCompliancePolicyAssignRequest) Get(ctx context.Context) ([][]DeviceCompliancePolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder struct{ BaseRequestBuilder }

// ScheduleActionsForRules action undocumented
func (b *DeviceCompliancePolicyRequestBuilder) ScheduleActionsForRules(reqObj *DeviceCompliancePolicyScheduleActionsForRulesRequestParameter) *DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder {
	bb := &DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/scheduleActionsForRules"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceCompliancePolicyScheduleActionsForRulesRequest struct{ BaseRequest }

//
func (b *DeviceCompliancePolicyScheduleActionsForRulesRequestBuilder) Request() *DeviceCompliancePolicyScheduleActionsForRulesRequest {
	return &DeviceCompliancePolicyScheduleActionsForRulesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceCompliancePolicyScheduleActionsForRulesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
