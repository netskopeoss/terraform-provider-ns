// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
	"net/http"
)

type GetNPAAppsRequest struct {
	// Return values only from specified fields
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *GetNPAAppsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type GetNPAAppsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	PrivateAppsGetResponse *shared.PrivateAppsGetResponse
	// Invalid request
	PrivateAppsResponse400 *shared.PrivateAppsResponse400
}

func (o *GetNPAAppsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNPAAppsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNPAAppsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetNPAAppsResponse) GetPrivateAppsGetResponse() *shared.PrivateAppsGetResponse {
	if o == nil {
		return nil
	}
	return o.PrivateAppsGetResponse
}

func (o *GetNPAAppsResponse) GetPrivateAppsResponse400() *shared.PrivateAppsResponse400 {
	if o == nil {
		return nil
	}
	return o.PrivateAppsResponse400
}
