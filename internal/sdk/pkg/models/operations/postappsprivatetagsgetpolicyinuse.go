// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostAppsPrivateTagsGetpolicyinuseRequestBody struct {
	Ids []string `json:"ids,omitempty"`
}

func (o *PostAppsPrivateTagsGetpolicyinuseRequestBody) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

// PostAppsPrivateTagsGetpolicyinuseResponseResponseBody - Invalid request
type PostAppsPrivateTagsGetpolicyinuseResponseResponseBody struct {
	Result *string `json:"result,omitempty"`
	Status *int64  `json:"status,omitempty"`
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponseResponseBody) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponseResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type PostAppsPrivateTagsGetpolicyinuseData struct {
	Token *string `json:"token,omitempty"`
}

func (o *PostAppsPrivateTagsGetpolicyinuseData) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

type PostAppsPrivateTagsGetpolicyinuseStatus string

const (
	PostAppsPrivateTagsGetpolicyinuseStatusSuccess  PostAppsPrivateTagsGetpolicyinuseStatus = "success"
	PostAppsPrivateTagsGetpolicyinuseStatusNotFound PostAppsPrivateTagsGetpolicyinuseStatus = "not found"
)

func (e PostAppsPrivateTagsGetpolicyinuseStatus) ToPointer() *PostAppsPrivateTagsGetpolicyinuseStatus {
	return &e
}

func (e *PostAppsPrivateTagsGetpolicyinuseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "not found":
		*e = PostAppsPrivateTagsGetpolicyinuseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostAppsPrivateTagsGetpolicyinuseStatus: %v", v)
	}
}

// PostAppsPrivateTagsGetpolicyinuseResponseBody - successful operation
type PostAppsPrivateTagsGetpolicyinuseResponseBody struct {
	Data   []PostAppsPrivateTagsGetpolicyinuseData  `json:"data,omitempty"`
	Status *PostAppsPrivateTagsGetpolicyinuseStatus `json:"status,omitempty"`
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponseBody) GetData() []PostAppsPrivateTagsGetpolicyinuseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponseBody) GetStatus() *PostAppsPrivateTagsGetpolicyinuseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type PostAppsPrivateTagsGetpolicyinuseResponse struct {
	// successful operation
	TwoHundredApplicationJSONObject *PostAppsPrivateTagsGetpolicyinuseResponseBody
	// Invalid request
	FourHundredApplicationJSONObject *PostAppsPrivateTagsGetpolicyinuseResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponse) GetTwoHundredApplicationJSONObject() *PostAppsPrivateTagsGetpolicyinuseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponse) GetFourHundredApplicationJSONObject() *PostAppsPrivateTagsGetpolicyinuseResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAppsPrivateTagsGetpolicyinuseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}