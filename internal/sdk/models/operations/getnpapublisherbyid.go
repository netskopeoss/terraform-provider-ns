// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
	"net/http"
)

type GetNPAPublisherByIDRequest struct {
	// publisher id
	PublisherID int `pathParam:"style=simple,explode=false,name=publisher_id"`
}

func (o *GetNPAPublisherByIDRequest) GetPublisherID() int {
	if o == nil {
		return 0
	}
	return o.PublisherID
}

type GetNPAPublisherByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	PublisherResponse *shared.PublisherResponse
	// Invalid request
	PublishersResponse400 *shared.PublishersResponse400
}

func (o *GetNPAPublisherByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNPAPublisherByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNPAPublisherByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetNPAPublisherByIDResponse) GetPublisherResponse() *shared.PublisherResponse {
	if o == nil {
		return nil
	}
	return o.PublisherResponse
}

func (o *GetNPAPublisherByIDResponse) GetPublishersResponse400() *shared.PublishersResponse400 {
	if o == nil {
		return nil
	}
	return o.PublishersResponse400
}
