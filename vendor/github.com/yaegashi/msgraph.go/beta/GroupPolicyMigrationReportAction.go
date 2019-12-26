// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// GroupPolicyMigrationReportCollectionCreateMigrationReportRequestParameter undocumented
type GroupPolicyMigrationReportCollectionCreateMigrationReportRequestParameter struct {
	// GroupPolicyObjectFile undocumented
	GroupPolicyObjectFile *GroupPolicyObjectFile `json:"groupPolicyObjectFile,omitempty"`
}

//
type GroupPolicyMigrationReportCollectionCreateMigrationReportRequestBuilder struct{ BaseRequestBuilder }

// CreateMigrationReport action undocumented
func (b *DeviceManagementGroupPolicyMigrationReportsCollectionRequestBuilder) CreateMigrationReport(reqObj *GroupPolicyMigrationReportCollectionCreateMigrationReportRequestParameter) *GroupPolicyMigrationReportCollectionCreateMigrationReportRequestBuilder {
	bb := &GroupPolicyMigrationReportCollectionCreateMigrationReportRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createMigrationReport"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupPolicyMigrationReportCollectionCreateMigrationReportRequest struct{ BaseRequest }

//
func (b *GroupPolicyMigrationReportCollectionCreateMigrationReportRequestBuilder) Request() *GroupPolicyMigrationReportCollectionCreateMigrationReportRequest {
	return &GroupPolicyMigrationReportCollectionCreateMigrationReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupPolicyMigrationReportCollectionCreateMigrationReportRequest) Post(ctx context.Context) (resObj *string, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
