// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteGreTunnelsRequest struct {
	// GRE tunnel id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteGreTunnelsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type DeleteGreTunnelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	GreResponse200 *shared.GreResponse200
	// Access forbidden
	GreResponse403 *shared.GreResponse403
	// Not found
	GreResponse404 *shared.GreResponse404
	// Method not allowed
	GreResponse405 *shared.GreResponse405
	// Too many requests
	GreResponse429 *shared.GreResponse429
	// Internal server error
	GreResponse500 *shared.GreResponse500
}

func (o *DeleteGreTunnelsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteGreTunnelsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteGreTunnelsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteGreTunnelsResponse) GetGreResponse200() *shared.GreResponse200 {
	if o == nil {
		return nil
	}
	return o.GreResponse200
}

func (o *DeleteGreTunnelsResponse) GetGreResponse403() *shared.GreResponse403 {
	if o == nil {
		return nil
	}
	return o.GreResponse403
}

func (o *DeleteGreTunnelsResponse) GetGreResponse404() *shared.GreResponse404 {
	if o == nil {
		return nil
	}
	return o.GreResponse404
}

func (o *DeleteGreTunnelsResponse) GetGreResponse405() *shared.GreResponse405 {
	if o == nil {
		return nil
	}
	return o.GreResponse405
}

func (o *DeleteGreTunnelsResponse) GetGreResponse429() *shared.GreResponse429 {
	if o == nil {
		return nil
	}
	return o.GreResponse429
}

func (o *DeleteGreTunnelsResponse) GetGreResponse500() *shared.GreResponse500 {
	if o == nil {
		return nil
	}
	return o.GreResponse500
}