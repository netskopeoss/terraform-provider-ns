// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PutPublisherupgradeprofilesUpgradeProfileIDRequestBody struct {
	DockerTag   *string `json:"docker_tag,omitempty"`
	Enabled     *int    `json:"enabled,omitempty"`
	Frequency   *string `json:"frequency,omitempty"`
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ReleaseType *string `json:"release_type,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetDockerTag() *string {
	if o == nil {
		return nil
	}
	return o.DockerTag
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetEnabled() *int {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetFrequency() *string {
	if o == nil {
		return nil
	}
	return o.Frequency
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetReleaseType() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseType
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequestBody) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

// PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent - flag to skip output except status code
type PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent string

const (
	PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilentOne  PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent = "1"
	PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilentZero PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent = "0"
)

func (e PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent) ToPointer() *PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent {
	return &e
}

func (e *PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1":
		fallthrough
	case "0":
		*e = PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent: %v", v)
	}
}

type PutPublisherupgradeprofilesUpgradeProfileIDRequest struct {
	RequestBody PutPublisherupgradeprofilesUpgradeProfileIDRequestBody `request:"mediaType=application/json"`
	// flag to skip output except status code
	Silent *PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent `queryParam:"style=form,explode=true,name=silent"`
	// publisher upgrade profile id
	UpgradeProfileID int `pathParam:"style=simple,explode=false,name=upgrade_profile_id"`
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequest) GetRequestBody() PutPublisherupgradeprofilesUpgradeProfileIDRequestBody {
	if o == nil {
		return PutPublisherupgradeprofilesUpgradeProfileIDRequestBody{}
	}
	return o.RequestBody
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequest) GetSilent() *PutPublisherupgradeprofilesUpgradeProfileIDQueryParamSilent {
	if o == nil {
		return nil
	}
	return o.Silent
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDRequest) GetUpgradeProfileID() int {
	if o == nil {
		return 0
	}
	return o.UpgradeProfileID
}

// PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody - Invalid request
type PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody struct {
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type PutPublisherupgradeprofilesUpgradeProfileIDResult struct {
}

// PutPublisherupgradeprofilesUpgradeProfileIDResponseBody - successful operation
type PutPublisherupgradeprofilesUpgradeProfileIDResponseBody struct {
	Result []PutPublisherupgradeprofilesUpgradeProfileIDResult `json:"result,omitempty"`
	Status *int64                                              `json:"status,omitempty"`
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponseBody) GetResult() []PutPublisherupgradeprofilesUpgradeProfileIDResult {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type PutPublisherupgradeprofilesUpgradeProfileIDResponse struct {
	// successful operation
	TwoHundredApplicationJSONObject *PutPublisherupgradeprofilesUpgradeProfileIDResponseBody
	// Invalid request
	FourHundredApplicationJSONObject *PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponse) GetTwoHundredApplicationJSONObject() *PutPublisherupgradeprofilesUpgradeProfileIDResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponse) GetFourHundredApplicationJSONObject() *PutPublisherupgradeprofilesUpgradeProfileIDNPAPublisherUpgradeProfilesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutPublisherupgradeprofilesUpgradeProfileIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}