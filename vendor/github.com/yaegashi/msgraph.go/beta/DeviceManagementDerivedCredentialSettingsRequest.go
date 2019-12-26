// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceManagementDerivedCredentialSettingsRequestBuilder is request builder for DeviceManagementDerivedCredentialSettings
type DeviceManagementDerivedCredentialSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementDerivedCredentialSettingsRequest
func (b *DeviceManagementDerivedCredentialSettingsRequestBuilder) Request() *DeviceManagementDerivedCredentialSettingsRequest {
	return &DeviceManagementDerivedCredentialSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementDerivedCredentialSettingsRequest is request for DeviceManagementDerivedCredentialSettings
type DeviceManagementDerivedCredentialSettingsRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementDerivedCredentialSettings
func (r *DeviceManagementDerivedCredentialSettingsRequest) Get(ctx context.Context) (resObj *DeviceManagementDerivedCredentialSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementDerivedCredentialSettings
func (r *DeviceManagementDerivedCredentialSettingsRequest) Update(ctx context.Context, reqObj *DeviceManagementDerivedCredentialSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementDerivedCredentialSettings
func (r *DeviceManagementDerivedCredentialSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
