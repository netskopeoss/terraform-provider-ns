// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

// UpdateNPAPublisherProfilesBulkQueryParamSilent - flag to skip output except status code:
//   - `1` - Skip response data
//   - `0` - Do not skip response data
type UpdateNPAPublisherProfilesBulkQueryParamSilent string

const (
	UpdateNPAPublisherProfilesBulkQueryParamSilentOne  UpdateNPAPublisherProfilesBulkQueryParamSilent = "1"
	UpdateNPAPublisherProfilesBulkQueryParamSilentZero UpdateNPAPublisherProfilesBulkQueryParamSilent = "0"
)

func (e UpdateNPAPublisherProfilesBulkQueryParamSilent) ToPointer() *UpdateNPAPublisherProfilesBulkQueryParamSilent {
	return &e
}

func (e *UpdateNPAPublisherProfilesBulkQueryParamSilent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		fallthrough
	case "0":
		*e = UpdateNPAPublisherProfilesBulkQueryParamSilent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateNPAPublisherProfilesBulkQueryParamSilent: %v", v)
	}
}

type UpdateNPAPublisherProfilesBulkRequest struct {
	// flag to skip output except status code:
	//  * `1` - Skip response data
	//  * `0` - Do not skip response data
	Silent                             *UpdateNPAPublisherProfilesBulkQueryParamSilent `queryParam:"style=form,explode=true,name=silent"`
	PublisherUpgradeProfileBulkRequest shared.PublisherUpgradeProfileBulkRequest       `request:"mediaType=application/json"`
}

func (o *UpdateNPAPublisherProfilesBulkRequest) GetSilent() *UpdateNPAPublisherProfilesBulkQueryParamSilent {
	if o == nil {
		return nil
	}
	return o.Silent
}

func (o *UpdateNPAPublisherProfilesBulkRequest) GetPublisherUpgradeProfileBulkRequest() shared.PublisherUpgradeProfileBulkRequest {
	if o == nil {
		return shared.PublisherUpgradeProfileBulkRequest{}
	}
	return o.PublisherUpgradeProfileBulkRequest
}

type UpdateNPAPublisherProfilesBulkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	PublisherUpgradeProfileBulkResponse *shared.PublisherUpgradeProfileBulkResponse
	// Invalid request
	FourHundred *shared.FourHundred
}

func (o *UpdateNPAPublisherProfilesBulkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNPAPublisherProfilesBulkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNPAPublisherProfilesBulkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNPAPublisherProfilesBulkResponse) GetPublisherUpgradeProfileBulkResponse() *shared.PublisherUpgradeProfileBulkResponse {
	if o == nil {
		return nil
	}
	return o.PublisherUpgradeProfileBulkResponse
}

func (o *UpdateNPAPublisherProfilesBulkResponse) GetFourHundred() *shared.FourHundred {
	if o == nil {
		return nil
	}
	return o.FourHundred
}