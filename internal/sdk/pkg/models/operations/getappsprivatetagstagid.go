// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetAppsPrivateTagsTagIDRequest struct {
	// tag id
	TagID int `pathParam:"style=simple,explode=false,name=tag_id"`
}

func (o *GetAppsPrivateTagsTagIDRequest) GetTagID() int {
	if o == nil {
		return 0
	}
	return o.TagID
}

// GetAppsPrivateTagsTagIDResponseResponseBody - Invalid request
type GetAppsPrivateTagsTagIDResponseResponseBody struct {
	Result *string `json:"result,omitempty"`
	Status *int64  `json:"status,omitempty"`
}

func (o *GetAppsPrivateTagsTagIDResponseResponseBody) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *GetAppsPrivateTagsTagIDResponseResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAppsPrivateTagsTagIDData struct {
	TagID   *int    `json:"tag_id,omitempty"`
	TagName *string `json:"tag_name,omitempty"`
}

func (o *GetAppsPrivateTagsTagIDData) GetTagID() *int {
	if o == nil {
		return nil
	}
	return o.TagID
}

func (o *GetAppsPrivateTagsTagIDData) GetTagName() *string {
	if o == nil {
		return nil
	}
	return o.TagName
}

type GetAppsPrivateTagsTagIDStatus string

const (
	GetAppsPrivateTagsTagIDStatusSuccess  GetAppsPrivateTagsTagIDStatus = "success"
	GetAppsPrivateTagsTagIDStatusNotFound GetAppsPrivateTagsTagIDStatus = "not found"
)

func (e GetAppsPrivateTagsTagIDStatus) ToPointer() *GetAppsPrivateTagsTagIDStatus {
	return &e
}

func (e *GetAppsPrivateTagsTagIDStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "not found":
		*e = GetAppsPrivateTagsTagIDStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAppsPrivateTagsTagIDStatus: %v", v)
	}
}

// GetAppsPrivateTagsTagIDResponseBody - successful operation
type GetAppsPrivateTagsTagIDResponseBody struct {
	Data   []GetAppsPrivateTagsTagIDData  `json:"data,omitempty"`
	Status *GetAppsPrivateTagsTagIDStatus `json:"status,omitempty"`
}

func (o *GetAppsPrivateTagsTagIDResponseBody) GetData() []GetAppsPrivateTagsTagIDData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetAppsPrivateTagsTagIDResponseBody) GetStatus() *GetAppsPrivateTagsTagIDStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetAppsPrivateTagsTagIDResponse struct {
	// successful operation
	TwoHundredApplicationJSONObject *GetAppsPrivateTagsTagIDResponseBody
	// Invalid request
	FourHundredApplicationJSONObject *GetAppsPrivateTagsTagIDResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAppsPrivateTagsTagIDResponse) GetTwoHundredApplicationJSONObject() *GetAppsPrivateTagsTagIDResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetAppsPrivateTagsTagIDResponse) GetFourHundredApplicationJSONObject() *GetAppsPrivateTagsTagIDResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetAppsPrivateTagsTagIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAppsPrivateTagsTagIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAppsPrivateTagsTagIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}