// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DeleteInfrastructurePublishersPublisherIDRequest struct {
	// publisher id
	PublisherID int `pathParam:"style=simple,explode=false,name=publisher_id"`
}

func (o *DeleteInfrastructurePublishersPublisherIDRequest) GetPublisherID() int {
	if o == nil {
		return 0
	}
	return o.PublisherID
}

// DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody - Invalid request
type DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody struct {
	Result *string `json:"result,omitempty"`
	Status *int64  `json:"status,omitempty"`
}

func (o *DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody) GetResult() *string {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteInfrastructurePublishersPublisherIDStatus string

const (
	DeleteInfrastructurePublishersPublisherIDStatusSuccess DeleteInfrastructurePublishersPublisherIDStatus = "success"
	DeleteInfrastructurePublishersPublisherIDStatusError   DeleteInfrastructurePublishersPublisherIDStatus = "error"
)

func (e DeleteInfrastructurePublishersPublisherIDStatus) ToPointer() *DeleteInfrastructurePublishersPublisherIDStatus {
	return &e
}

func (e *DeleteInfrastructurePublishersPublisherIDStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteInfrastructurePublishersPublisherIDStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteInfrastructurePublishersPublisherIDStatus: %v", v)
	}
}

// DeleteInfrastructurePublishersPublisherIDResponseBody - successful operation
type DeleteInfrastructurePublishersPublisherIDResponseBody struct {
	Status *DeleteInfrastructurePublishersPublisherIDStatus `json:"status,omitempty"`
}

func (o *DeleteInfrastructurePublishersPublisherIDResponseBody) GetStatus() *DeleteInfrastructurePublishersPublisherIDStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteInfrastructurePublishersPublisherIDResponse struct {
	// successful operation
	TwoHundredApplicationJSONObject *DeleteInfrastructurePublishersPublisherIDResponseBody
	// Invalid request
	FourHundredApplicationJSONObject *DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteInfrastructurePublishersPublisherIDResponse) GetTwoHundredApplicationJSONObject() *DeleteInfrastructurePublishersPublisherIDResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *DeleteInfrastructurePublishersPublisherIDResponse) GetFourHundredApplicationJSONObject() *DeleteInfrastructurePublishersPublisherIDNPAPublishersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *DeleteInfrastructurePublishersPublisherIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteInfrastructurePublishersPublisherIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteInfrastructurePublishersPublisherIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}