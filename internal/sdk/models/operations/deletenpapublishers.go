// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteNPAPublishersRequest struct {
	// publisher id
	PublisherID int `pathParam:"style=simple,explode=false,name=publisher_id"`
}

func (o *DeleteNPAPublishersRequest) GetPublisherID() int {
	if o == nil {
		return 0
	}
	return o.PublisherID
}

type DeleteNPAPublishersStatus string

const (
	DeleteNPAPublishersStatusSuccess DeleteNPAPublishersStatus = "success"
	DeleteNPAPublishersStatusError   DeleteNPAPublishersStatus = "error"
)

func (e DeleteNPAPublishersStatus) ToPointer() *DeleteNPAPublishersStatus {
	return &e
}

func (e *DeleteNPAPublishersStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = DeleteNPAPublishersStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteNPAPublishersStatus: %v", v)
	}
}

// DeleteNPAPublishersResponseBody - successful operation
type DeleteNPAPublishersResponseBody struct {
	Status *DeleteNPAPublishersStatus `json:"status,omitempty"`
}

func (o *DeleteNPAPublishersResponseBody) GetStatus() *DeleteNPAPublishersStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type DeleteNPAPublishersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	Object *DeleteNPAPublishersResponseBody
	// Invalid request
	PublishersResponse400 *shared.PublishersResponse400
}

func (o *DeleteNPAPublishersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNPAPublishersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNPAPublishersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteNPAPublishersResponse) GetObject() *DeleteNPAPublishersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteNPAPublishersResponse) GetPublishersResponse400() *shared.PublishersResponse400 {
	if o == nil {
		return nil
	}
	return o.PublishersResponse400
}