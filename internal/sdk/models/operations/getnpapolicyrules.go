// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetNPAPolicyRulesRequest struct {
	// Return values only from specified fields
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Query string based on query operaters
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
	// Max number of policies to retrieve. Default will be all policies.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// The offset of the first policy in the list to retrieve.
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// Sort retrieved policies by specified field. Default is policy id
	Sortby *string `queryParam:"style=form,explode=true,name=sortby"`
	// Sort in either asc or desc order. The default is asc order
	Sortorder *string `queryParam:"style=form,explode=true,name=sortorder"`
}

func (o *GetNPAPolicyRulesRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetNPAPolicyRulesRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetNPAPolicyRulesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetNPAPolicyRulesRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetNPAPolicyRulesRequest) GetSortby() *string {
	if o == nil {
		return nil
	}
	return o.Sortby
}

func (o *GetNPAPolicyRulesRequest) GetSortorder() *string {
	if o == nil {
		return nil
	}
	return o.Sortorder
}

type GetNPAPolicyRulesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	NpaPolicyListResponse *shared.NpaPolicyListResponse
	// Invalid request
	NpaPolicyResponse400 *shared.NpaPolicyResponse400
}

func (o *GetNPAPolicyRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNPAPolicyRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNPAPolicyRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetNPAPolicyRulesResponse) GetNpaPolicyListResponse() *shared.NpaPolicyListResponse {
	if o == nil {
		return nil
	}
	return o.NpaPolicyListResponse
}

func (o *GetNPAPolicyRulesResponse) GetNpaPolicyResponse400() *shared.NpaPolicyResponse400 {
	if o == nil {
		return nil
	}
	return o.NpaPolicyResponse400
}