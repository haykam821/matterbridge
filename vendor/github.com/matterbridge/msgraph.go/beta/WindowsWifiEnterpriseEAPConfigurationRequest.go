// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matterbridge/msgraph.go/jsonx"
)

// WindowsWifiEnterpriseEAPConfigurationRequestBuilder is request builder for WindowsWifiEnterpriseEAPConfiguration
type WindowsWifiEnterpriseEAPConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsWifiEnterpriseEAPConfigurationRequest
func (b *WindowsWifiEnterpriseEAPConfigurationRequestBuilder) Request() *WindowsWifiEnterpriseEAPConfigurationRequest {
	return &WindowsWifiEnterpriseEAPConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsWifiEnterpriseEAPConfigurationRequest is request for WindowsWifiEnterpriseEAPConfiguration
type WindowsWifiEnterpriseEAPConfigurationRequest struct{ BaseRequest }

// Get performs GET request for WindowsWifiEnterpriseEAPConfiguration
func (r *WindowsWifiEnterpriseEAPConfigurationRequest) Get(ctx context.Context) (resObj *WindowsWifiEnterpriseEAPConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsWifiEnterpriseEAPConfiguration
func (r *WindowsWifiEnterpriseEAPConfigurationRequest) Update(ctx context.Context, reqObj *WindowsWifiEnterpriseEAPConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsWifiEnterpriseEAPConfiguration
func (r *WindowsWifiEnterpriseEAPConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *WindowsWifiEnterpriseEAPConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *WindowsCertificateProfileBaseRequestBuilder {
	bb := &WindowsCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificatesForServerValidation returns request builder for Windows81TrustedRootCertificate collection
func (b *WindowsWifiEnterpriseEAPConfigurationRequestBuilder) RootCertificatesForServerValidation() *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder {
	bb := &WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificatesForServerValidation"
	return bb
}

// WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder is request builder for Windows81TrustedRootCertificate collection
type WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Windows81TrustedRootCertificate collection
func (b *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder) Request() *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest {
	return &WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Windows81TrustedRootCertificate item
func (b *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequestBuilder) ID(id string) *Windows81TrustedRootCertificateRequestBuilder {
	bb := &Windows81TrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest is request for Windows81TrustedRootCertificate collection
type WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Windows81TrustedRootCertificate collection
func (r *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Windows81TrustedRootCertificate, error) {
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
	var values []Windows81TrustedRootCertificate
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
			value  []Windows81TrustedRootCertificate
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

// Get performs GET request for Windows81TrustedRootCertificate collection
func (r *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest) Get(ctx context.Context) ([]Windows81TrustedRootCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Windows81TrustedRootCertificate collection
func (r *WindowsWifiEnterpriseEAPConfigurationRootCertificatesForServerValidationCollectionRequest) Add(ctx context.Context, reqObj *Windows81TrustedRootCertificate) (resObj *Windows81TrustedRootCertificate, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
