// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/models/shared"
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/utils"
	"net/http"
)

type GetSteeringGreTunnelsRequest struct {
	// Provide comma separated list of fields to be displayed
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Max number of records to retrieve
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// Offset used to shift the window
	Offset *int64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// GRE tunnel POP name
	Pop *string `queryParam:"style=form,explode=true,name=pop"`
	// GRE tunnel site
	Site *string `queryParam:"style=form,explode=true,name=site"`
	// GRE tunnel sortby column
	Sortby *string `queryParam:"style=form,explode=true,name=sortby"`
	// GRE tunnel sortorder asc or desc. Default is asc
	Sortorder *string `queryParam:"style=form,explode=true,name=sortorder"`
	// GRE tunnel sourcetype
	Sourcetype *string `queryParam:"style=form,explode=true,name=sourcetype"`
	// GRE tunnel source ip identity
	Srcipidentity *string `queryParam:"style=form,explode=true,name=srcipidentity"`
	// GRE tunnel status
	Status *string `queryParam:"style=form,explode=true,name=status"`
}

func (g GetSteeringGreTunnelsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetSteeringGreTunnelsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetSteeringGreTunnelsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetSteeringGreTunnelsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetSteeringGreTunnelsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetSteeringGreTunnelsRequest) GetPop() *string {
	if o == nil {
		return nil
	}
	return o.Pop
}

func (o *GetSteeringGreTunnelsRequest) GetSite() *string {
	if o == nil {
		return nil
	}
	return o.Site
}

func (o *GetSteeringGreTunnelsRequest) GetSortby() *string {
	if o == nil {
		return nil
	}
	return o.Sortby
}

func (o *GetSteeringGreTunnelsRequest) GetSortorder() *string {
	if o == nil {
		return nil
	}
	return o.Sortorder
}

func (o *GetSteeringGreTunnelsRequest) GetSourcetype() *string {
	if o == nil {
		return nil
	}
	return o.Sourcetype
}

func (o *GetSteeringGreTunnelsRequest) GetSrcipidentity() *string {
	if o == nil {
		return nil
	}
	return o.Srcipidentity
}

func (o *GetSteeringGreTunnelsRequest) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetSteeringGreTunnelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
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
	// Successful operation
	GreTunnelResponse200 *shared.GreTunnelResponse200
}

func (o *GetSteeringGreTunnelsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSteeringGreTunnelsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSteeringGreTunnelsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSteeringGreTunnelsResponse) GetGreResponse400() *shared.GreResponse400 {
	if o == nil {
		return nil
	}
	return o.GreResponse400
}

func (o *GetSteeringGreTunnelsResponse) GetGreResponse403() *shared.GreResponse403 {
	if o == nil {
		return nil
	}
	return o.GreResponse403
}

func (o *GetSteeringGreTunnelsResponse) GetGreResponse405() *shared.GreResponse405 {
	if o == nil {
		return nil
	}
	return o.GreResponse405
}

func (o *GetSteeringGreTunnelsResponse) GetGreResponse429() *shared.GreResponse429 {
	if o == nil {
		return nil
	}
	return o.GreResponse429
}

func (o *GetSteeringGreTunnelsResponse) GetGreResponse500() *shared.GreResponse500 {
	if o == nil {
		return nil
	}
	return o.GreResponse500
}

func (o *GetSteeringGreTunnelsResponse) GetGreTunnelResponse200() *shared.GreTunnelResponse200 {
	if o == nil {
		return nil
	}
	return o.GreTunnelResponse200
}