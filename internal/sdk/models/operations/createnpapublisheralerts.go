// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
	"net/http"
)

type CreateNPAPublisherAlertsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	PublishersAlertPutResponse *shared.PublishersAlertPutResponse
	// Invalid request
	PublishersResponse400 *shared.PublishersResponse400
}

func (o *CreateNPAPublisherAlertsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateNPAPublisherAlertsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateNPAPublisherAlertsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateNPAPublisherAlertsResponse) GetPublishersAlertPutResponse() *shared.PublishersAlertPutResponse {
	if o == nil {
		return nil
	}
	return o.PublishersAlertPutResponse
}

func (o *CreateNPAPublisherAlertsResponse) GetPublishersResponse400() *shared.PublishersResponse400 {
	if o == nil {
		return nil
	}
	return o.PublishersResponse400
}