// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceManagementReportsGetDeviceNonComplianceReportRequestParameter undocumented
type DeviceManagementReportsGetDeviceNonComplianceReportRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Select undocumented
	Select []string `json:"select,omitempty"`
	// Search undocumented
	Search *string `json:"search,omitempty"`
	// GroupBy undocumented
	GroupBy []string `json:"groupBy,omitempty"`
	// OrderBy undocumented
	OrderBy []string `json:"orderBy,omitempty"`
	// Skip undocumented
	Skip *int `json:"skip,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
	// SessionID undocumented
	SessionID *string `json:"sessionId,omitempty"`
	// Filter undocumented
	Filter *string `json:"filter,omitempty"`
}

// DeviceManagementReportsGetPolicyNonComplianceReportRequestParameter undocumented
type DeviceManagementReportsGetPolicyNonComplianceReportRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Select undocumented
	Select []string `json:"select,omitempty"`
	// Search undocumented
	Search *string `json:"search,omitempty"`
	// GroupBy undocumented
	GroupBy []string `json:"groupBy,omitempty"`
	// OrderBy undocumented
	OrderBy []string `json:"orderBy,omitempty"`
	// Skip undocumented
	Skip *int `json:"skip,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
	// SessionID undocumented
	SessionID *string `json:"sessionId,omitempty"`
	// Filter undocumented
	Filter *string `json:"filter,omitempty"`
}

// DeviceManagementReportsGetPolicyNonComplianceMetadataRequestParameter undocumented
type DeviceManagementReportsGetPolicyNonComplianceMetadataRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Select undocumented
	Select []string `json:"select,omitempty"`
	// Search undocumented
	Search *string `json:"search,omitempty"`
	// GroupBy undocumented
	GroupBy []string `json:"groupBy,omitempty"`
	// OrderBy undocumented
	OrderBy []string `json:"orderBy,omitempty"`
	// Skip undocumented
	Skip *int `json:"skip,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
	// SessionID undocumented
	SessionID *string `json:"sessionId,omitempty"`
	// Filter undocumented
	Filter *string `json:"filter,omitempty"`
}

// DeviceManagementReportsGetHistoricalReportRequestParameter undocumented
type DeviceManagementReportsGetHistoricalReportRequestParameter struct {
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Select undocumented
	Select []string `json:"select,omitempty"`
	// Search undocumented
	Search *string `json:"search,omitempty"`
	// GroupBy undocumented
	GroupBy []string `json:"groupBy,omitempty"`
	// OrderBy undocumented
	OrderBy []string `json:"orderBy,omitempty"`
	// Skip undocumented
	Skip *int `json:"skip,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
	// Filter undocumented
	Filter *string `json:"filter,omitempty"`
}

// DeviceManagementReportsGetCachedReportRequestParameter undocumented
type DeviceManagementReportsGetCachedReportRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Select undocumented
	Select []string `json:"select,omitempty"`
	// Search undocumented
	Search *string `json:"search,omitempty"`
	// GroupBy undocumented
	GroupBy []string `json:"groupBy,omitempty"`
	// OrderBy undocumented
	OrderBy []string `json:"orderBy,omitempty"`
	// Skip undocumented
	Skip *int `json:"skip,omitempty"`
	// Top undocumented
	Top *int `json:"top,omitempty"`
}

//
type DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder struct{ BaseRequestBuilder }

// GetDeviceNonComplianceReport action undocumented
func (b *DeviceManagementReportsRequestBuilder) GetDeviceNonComplianceReport(reqObj *DeviceManagementReportsGetDeviceNonComplianceReportRequestParameter) *DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder {
	bb := &DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getDeviceNonComplianceReport"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementReportsGetDeviceNonComplianceReportRequest struct{ BaseRequest }

//
func (b *DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder) Request() *DeviceManagementReportsGetDeviceNonComplianceReportRequest {
	return &DeviceManagementReportsGetDeviceNonComplianceReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementReportsGetDeviceNonComplianceReportRequest) Post(ctx context.Context) (resObj *Stream, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder struct{ BaseRequestBuilder }

// GetPolicyNonComplianceReport action undocumented
func (b *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceReport(reqObj *DeviceManagementReportsGetPolicyNonComplianceReportRequestParameter) *DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder {
	bb := &DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPolicyNonComplianceReport"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementReportsGetPolicyNonComplianceReportRequest struct{ BaseRequest }

//
func (b *DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder) Request() *DeviceManagementReportsGetPolicyNonComplianceReportRequest {
	return &DeviceManagementReportsGetPolicyNonComplianceReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementReportsGetPolicyNonComplianceReportRequest) Post(ctx context.Context) (resObj *Stream, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder struct{ BaseRequestBuilder }

// GetPolicyNonComplianceMetadata action undocumented
func (b *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceMetadata(reqObj *DeviceManagementReportsGetPolicyNonComplianceMetadataRequestParameter) *DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder {
	bb := &DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPolicyNonComplianceMetadata"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementReportsGetPolicyNonComplianceMetadataRequest struct{ BaseRequest }

//
func (b *DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder) Request() *DeviceManagementReportsGetPolicyNonComplianceMetadataRequest {
	return &DeviceManagementReportsGetPolicyNonComplianceMetadataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementReportsGetPolicyNonComplianceMetadataRequest) Post(ctx context.Context) (resObj *Stream, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DeviceManagementReportsGetHistoricalReportRequestBuilder struct{ BaseRequestBuilder }

// GetHistoricalReport action undocumented
func (b *DeviceManagementReportsRequestBuilder) GetHistoricalReport(reqObj *DeviceManagementReportsGetHistoricalReportRequestParameter) *DeviceManagementReportsGetHistoricalReportRequestBuilder {
	bb := &DeviceManagementReportsGetHistoricalReportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getHistoricalReport"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementReportsGetHistoricalReportRequest struct{ BaseRequest }

//
func (b *DeviceManagementReportsGetHistoricalReportRequestBuilder) Request() *DeviceManagementReportsGetHistoricalReportRequest {
	return &DeviceManagementReportsGetHistoricalReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementReportsGetHistoricalReportRequest) Post(ctx context.Context) (resObj *Stream, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DeviceManagementReportsGetCachedReportRequestBuilder struct{ BaseRequestBuilder }

// GetCachedReport action undocumented
func (b *DeviceManagementReportsRequestBuilder) GetCachedReport(reqObj *DeviceManagementReportsGetCachedReportRequestParameter) *DeviceManagementReportsGetCachedReportRequestBuilder {
	bb := &DeviceManagementReportsGetCachedReportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getCachedReport"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementReportsGetCachedReportRequest struct{ BaseRequest }

//
func (b *DeviceManagementReportsGetCachedReportRequestBuilder) Request() *DeviceManagementReportsGetCachedReportRequest {
	return &DeviceManagementReportsGetCachedReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementReportsGetCachedReportRequest) Post(ctx context.Context) (resObj *Stream, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
