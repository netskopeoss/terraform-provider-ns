// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeletePolicyNpaRulesIDRequest struct {
	// npa policy id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeletePolicyNpaRulesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeletePolicyNpaRulesIDResponseBody - successful operation
type DeletePolicyNpaRulesIDResponseBody struct {
	Data *shared.NpaPolicyResponseItem `json:"data,omitempty"`
}

func (o *DeletePolicyNpaRulesIDResponseBody) GetData() *shared.NpaPolicyResponseItem {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeletePolicyNpaRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Invalid request
	NpaPolicyResponse400 *shared.NpaPolicyResponse400
	// successful operation
	Object *DeletePolicyNpaRulesIDResponseBody
}

func (o *DeletePolicyNpaRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePolicyNpaRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePolicyNpaRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePolicyNpaRulesIDResponse) GetNpaPolicyResponse400() *shared.NpaPolicyResponse400 {
	if o == nil {
		return nil
	}
	return o.NpaPolicyResponse400
}

func (o *DeletePolicyNpaRulesIDResponse) GetObject() *DeletePolicyNpaRulesIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}