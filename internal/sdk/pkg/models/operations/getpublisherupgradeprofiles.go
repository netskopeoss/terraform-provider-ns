// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetPublisherupgradeprofilesRequest struct {
	// Return values only from specified fields
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *GetPublisherupgradeprofilesRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

// GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody - Invalid request
type GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody struct {
	Message *string `json:"message,omitempty"`
	Status  *int64  `json:"status,omitempty"`
}

func (o *GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetPublisherupgradeprofilesData struct {
	DockerTag   *string `json:"docker_tag,omitempty"`
	Enabled     *int    `json:"enabled,omitempty"`
	Frequency   *string `json:"frequency,omitempty"`
	ID          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ReleaseType *string `json:"release_type,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
}

func (o *GetPublisherupgradeprofilesData) GetDockerTag() *string {
	if o == nil {
		return nil
	}
	return o.DockerTag
}

func (o *GetPublisherupgradeprofilesData) GetEnabled() *int {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *GetPublisherupgradeprofilesData) GetFrequency() *string {
	if o == nil {
		return nil
	}
	return o.Frequency
}

func (o *GetPublisherupgradeprofilesData) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPublisherupgradeprofilesData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetPublisherupgradeprofilesData) GetReleaseType() *string {
	if o == nil {
		return nil
	}
	return o.ReleaseType
}

func (o *GetPublisherupgradeprofilesData) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

type GetPublisherupgradeprofilesStatus string

const (
	GetPublisherupgradeprofilesStatusSuccess  GetPublisherupgradeprofilesStatus = "success"
	GetPublisherupgradeprofilesStatusNotFound GetPublisherupgradeprofilesStatus = "not found"
)

func (e GetPublisherupgradeprofilesStatus) ToPointer() *GetPublisherupgradeprofilesStatus {
	return &e
}

func (e *GetPublisherupgradeprofilesStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "not found":
		*e = GetPublisherupgradeprofilesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetPublisherupgradeprofilesStatus: %v", v)
	}
}

// GetPublisherupgradeprofilesResponseBody - successful operation
type GetPublisherupgradeprofilesResponseBody struct {
	Data   []GetPublisherupgradeprofilesData  `json:"data,omitempty"`
	Status *GetPublisherupgradeprofilesStatus `json:"status,omitempty"`
	Total  *int                               `json:"total,omitempty"`
}

func (o *GetPublisherupgradeprofilesResponseBody) GetData() []GetPublisherupgradeprofilesData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetPublisherupgradeprofilesResponseBody) GetStatus() *GetPublisherupgradeprofilesStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetPublisherupgradeprofilesResponseBody) GetTotal() *int {
	if o == nil {
		return nil
	}
	return o.Total
}

type GetPublisherupgradeprofilesResponse struct {
	// successful operation
	TwoHundredApplicationJSONObject *GetPublisherupgradeprofilesResponseBody
	// Invalid request
	FourHundredApplicationJSONObject *GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPublisherupgradeprofilesResponse) GetTwoHundredApplicationJSONObject() *GetPublisherupgradeprofilesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetPublisherupgradeprofilesResponse) GetFourHundredApplicationJSONObject() *GetPublisherupgradeprofilesNPAPublisherUpgradeProfilesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetPublisherupgradeprofilesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPublisherupgradeprofilesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPublisherupgradeprofilesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}