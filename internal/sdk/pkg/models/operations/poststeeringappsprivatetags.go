// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostSteeringAppsPrivateTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	TagResponse *shared.TagResponse
	// Invalid request
	TagResponse400 *shared.TagResponse400
}

func (o *PostSteeringAppsPrivateTagsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostSteeringAppsPrivateTagsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostSteeringAppsPrivateTagsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostSteeringAppsPrivateTagsResponse) GetTagResponse() *shared.TagResponse {
	if o == nil {
		return nil
	}
	return o.TagResponse
}

func (o *PostSteeringAppsPrivateTagsResponse) GetTagResponse400() *shared.TagResponse400 {
	if o == nil {
		return nil
	}
	return o.TagResponse400
}