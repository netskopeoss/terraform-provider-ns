// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetIpsecPopsRequest struct {
	// POP name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// POP region, requires ISO-3166 region code and return POPs closest to value
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// POP country, requires ISO-3166 country code and return POPs closest to value
	Country *string `queryParam:"style=form,explode=true,name=country"`
	// POP latitude, lat requires long in combination and return POPs closest to value
	Lat *string `queryParam:"style=form,explode=true,name=lat"`
	// POP longitude, long requires lat in combination and return POPs closest to value
	Long *string `queryParam:"style=form,explode=true,name=long"`
	// POP ipaddress, return POPs closest to value
	IP *string `queryParam:"style=form,explode=true,name=ip"`
	// POP capabilities, return POPs capabilities
	Capabilities *string `queryParam:"style=form,explode=true,name=capabilities"`
	// Provide comma separated list of fields to be displayed
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// Offset used to shift the output window
	Offset *int64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// Max number of POPs to retrieve
	Limit *int64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
}

func (g GetIpsecPopsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetIpsecPopsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetIpsecPopsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetIpsecPopsRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetIpsecPopsRequest) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetIpsecPopsRequest) GetLat() *string {
	if o == nil {
		return nil
	}
	return o.Lat
}

func (o *GetIpsecPopsRequest) GetLong() *string {
	if o == nil {
		return nil
	}
	return o.Long
}

func (o *GetIpsecPopsRequest) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *GetIpsecPopsRequest) GetCapabilities() *string {
	if o == nil {
		return nil
	}
	return o.Capabilities
}

func (o *GetIpsecPopsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetIpsecPopsRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetIpsecPopsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

type GetIpsecPopsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	IpsecPopResponse200 *shared.IpsecPopResponse200
	// Invalid request
	IpsecResponse400 *shared.IpsecResponse400
	// Access forbidden
	IpsecResponse403 *shared.IpsecResponse403
	// Method not allowed
	IpsecResponse405 *shared.IpsecResponse405
	// Too many requests
	IpsecResponse429 *shared.IpsecResponse429
	// Internal server error
	IpsecResponse500 *shared.IpsecResponse500
}

func (o *GetIpsecPopsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIpsecPopsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIpsecPopsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetIpsecPopsResponse) GetIpsecPopResponse200() *shared.IpsecPopResponse200 {
	if o == nil {
		return nil
	}
	return o.IpsecPopResponse200
}

func (o *GetIpsecPopsResponse) GetIpsecResponse400() *shared.IpsecResponse400 {
	if o == nil {
		return nil
	}
	return o.IpsecResponse400
}

func (o *GetIpsecPopsResponse) GetIpsecResponse403() *shared.IpsecResponse403 {
	if o == nil {
		return nil
	}
	return o.IpsecResponse403
}

func (o *GetIpsecPopsResponse) GetIpsecResponse405() *shared.IpsecResponse405 {
	if o == nil {
		return nil
	}
	return o.IpsecResponse405
}

func (o *GetIpsecPopsResponse) GetIpsecResponse429() *shared.IpsecResponse429 {
	if o == nil {
		return nil
	}
	return o.IpsecResponse429
}

func (o *GetIpsecPopsResponse) GetIpsecResponse500() *shared.IpsecResponse500 {
	if o == nil {
		return nil
	}
	return o.IpsecResponse500
}