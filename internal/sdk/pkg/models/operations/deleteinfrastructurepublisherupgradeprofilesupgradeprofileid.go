// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDRequest struct {
	// publisher upgrade profile id
	UpgradeProfileID int `pathParam:"style=simple,explode=false,name=upgrade_profile_id"`
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDRequest) GetUpgradeProfileID() int {
	if o == nil {
		return 0
	}
	return o.UpgradeProfileID
}

type DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus string

const (
	DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatusSuccess DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus = "success"
	DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatusError   DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus = "error"
)

func (e DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus) ToPointer() *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus {
	return &e
}

func (e *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus: %v", v)
	}
}

// DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponseBody - successful operation
type DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponseBody struct {
	Status *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus `json:"status,omitempty"`
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponseBody) GetStatus() *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse struct {
	// Invalid request
	FourHundred *shared.FourHundred
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	Object *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponseBody
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse) GetFourHundred() *shared.FourHundred {
	if o == nil {
		return nil
	}
	return o.FourHundred
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponse) GetObject() *DeleteInfrastructurePublisherupgradeprofilesUpgradeProfileIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}