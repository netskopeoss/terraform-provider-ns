// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PublishersBulkResponseStatus string

const (
	PublishersBulkResponseStatusSuccess  PublishersBulkResponseStatus = "success"
	PublishersBulkResponseStatusNotFound PublishersBulkResponseStatus = "not found"
)

func (e PublishersBulkResponseStatus) ToPointer() *PublishersBulkResponseStatus {
	return &e
}

func (e *PublishersBulkResponseStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "not found":
		*e = PublishersBulkResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PublishersBulkResponseStatus: %v", v)
	}
}

type PublishersBulkResponse struct {
	Status *PublishersBulkResponseStatus `json:"status,omitempty"`
	Data   []PublisherBulkItem           `json:"data,omitempty"`
}

func (o *PublishersBulkResponse) GetStatus() *PublishersBulkResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PublishersBulkResponse) GetData() []PublisherBulkItem {
	if o == nil {
		return nil
	}
	return o.Data
}