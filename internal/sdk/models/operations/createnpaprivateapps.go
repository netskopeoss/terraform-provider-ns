// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
	"net/http"
)

// Silent - flag to skip output except status code
type Silent string

const (
	SilentOne  Silent = "1"
	SilentZero Silent = "0"
)

func (e Silent) ToPointer() *Silent {
	return &e
}
func (e *Silent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		fallthrough
	case "0":
		*e = Silent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Silent: %v", v)
	}
}

type CreateNPAPrivateAppsRequest struct {
	// flag to skip output except status code
	Silent             *Silent                   `queryParam:"style=form,explode=true,name=silent"`
	PrivateAppsRequest shared.PrivateAppsRequest `request:"mediaType=application/json"`
}

func (o *CreateNPAPrivateAppsRequest) GetSilent() *Silent {
	if o == nil {
		return nil
	}
	return o.Silent
}

func (o *CreateNPAPrivateAppsRequest) GetPrivateAppsRequest() shared.PrivateAppsRequest {
	if o == nil {
		return shared.PrivateAppsRequest{}
	}
	return o.PrivateAppsRequest
}

type CreateNPAPrivateAppsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	PrivateAppsResponse []shared.PrivateAppsResponse
	// Invalid request
	PrivateAppsResponse400 *shared.PrivateAppsResponse400
}

func (o *CreateNPAPrivateAppsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateNPAPrivateAppsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateNPAPrivateAppsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateNPAPrivateAppsResponse) GetPrivateAppsResponse() []shared.PrivateAppsResponse {
	if o == nil {
		return nil
	}
	return o.PrivateAppsResponse
}

func (o *CreateNPAPrivateAppsResponse) GetPrivateAppsResponse400() *shared.PrivateAppsResponse400 {
	if o == nil {
		return nil
	}
	return o.PrivateAppsResponse400
}