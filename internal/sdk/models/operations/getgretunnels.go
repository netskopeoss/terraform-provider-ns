// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetGreTunnelsRequest struct {
	// GRE tunnel site
	Site *string `queryParam:"style=form,explode=true,name=site"`
	// GRE tunnel source ip identity
	Srcipidentity *string `queryParam:"style=form,explode=true,name=srcipidentity"`
	// GRE tunnel POP name
	Pop *string `queryParam:"style=form,explode=true,name=pop"`
	// GRE tunnel status
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// GRE tunnel sourcetype
	Sourcetype *string `queryParam:"style=form,explode=true,name=sourcetype"`
	// GRE tunnel sortby column
	Sortby *string `queryParam:"style=form,explode=true,name=sortby"`
	// GRE tunnel sortorder asc or desc. Default is asc
	Sortorder *string `queryParam:"style=form,explode=true,name=sortorder"`
	// Provide comma separated list of fields to be displayed
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Offset used to shift the window
	Offset *int64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// Max number of records to retrieve
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
}

func (g GetGreTunnelsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetGreTunnelsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetGreTunnelsRequest) GetSite() *string {
	if o == nil {
		return nil
	}
	return o.Site
}

func (o *GetGreTunnelsRequest) GetSrcipidentity() *string {
	if o == nil {
		return nil
	}
	return o.Srcipidentity
}

func (o *GetGreTunnelsRequest) GetPop() *string {
	if o == nil {
		return nil
	}
	return o.Pop
}

func (o *GetGreTunnelsRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetGreTunnelsRequest) GetSourcetype() *string {
	if o == nil {
		return nil
	}
	return o.Sourcetype
}

func (o *GetGreTunnelsRequest) GetSortby() *string {
	if o == nil {
		return nil
	}
	return o.Sortby
}

func (o *GetGreTunnelsRequest) GetSortorder() *string {
	if o == nil {
		return nil
	}
	return o.Sortorder
}

func (o *GetGreTunnelsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetGreTunnelsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetGreTunnelsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetGreTunnelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	GreTunnelResponse200 *shared.GreTunnelResponse200
	// Invalid request
	GreResponse400 *shared.GreResponse400
	// Access forbidden
	GreResponse403 *shared.GreResponse403
	// Method not allowed
	GreResponse405 *shared.GreResponse405
	// Too many requests
	GreResponse429 *shared.GreResponse429
	// Internal server error
	GreResponse500 *shared.GreResponse500
}

func (o *GetGreTunnelsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGreTunnelsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGreTunnelsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGreTunnelsResponse) GetGreTunnelResponse200() *shared.GreTunnelResponse200 {
	if o == nil {
		return nil
	}
	return o.GreTunnelResponse200
}

func (o *GetGreTunnelsResponse) GetGreResponse400() *shared.GreResponse400 {
	if o == nil {
		return nil
	}
	return o.GreResponse400
}

func (o *GetGreTunnelsResponse) GetGreResponse403() *shared.GreResponse403 {
	if o == nil {
		return nil
	}
	return o.GreResponse403
}

func (o *GetGreTunnelsResponse) GetGreResponse405() *shared.GreResponse405 {
	if o == nil {
		return nil
	}
	return o.GreResponse405
}

func (o *GetGreTunnelsResponse) GetGreResponse429() *shared.GreResponse429 {
	if o == nil {
		return nil
	}
	return o.GreResponse429
}

func (o *GetGreTunnelsResponse) GetGreResponse500() *shared.GreResponse500 {
	if o == nil {
		return nil
	}
	return o.GreResponse500
}