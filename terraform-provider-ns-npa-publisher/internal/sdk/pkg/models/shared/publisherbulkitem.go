// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PublisherBulkItemAssessment struct {
}

type PublisherBulkItemStatus string

const (
	PublisherBulkItemStatusConnected     PublisherBulkItemStatus = "connected"
	PublisherBulkItemStatusNotRegistered PublisherBulkItemStatus = "not registered"
)

func (e PublisherBulkItemStatus) ToPointer() *PublisherBulkItemStatus {
	return &e
}

func (e *PublisherBulkItemStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "connected":
		fallthrough
	case "not registered":
		*e = PublisherBulkItemStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PublisherBulkItemStatus: %v", v)
	}
}

type PublisherBulkItemUpgradeFailedReason struct {
}

type PublisherBulkItemUpgradeStatus struct {
}

type PublisherBulkItem struct {
	Assessment                *PublisherBulkItemAssessment          `json:"assessment,omitempty"`
	CommonName                *string                               `json:"common_name,omitempty"`
	ID                        *int                                  `json:"id,omitempty"`
	Lbrokerconnect            *bool                                 `json:"lbrokerconnect,omitempty"`
	Name                      *string                               `json:"name,omitempty"`
	PublisherUpgradeProfileID *int                                  `json:"publisher_upgrade_profile_id,omitempty"`
	Registered                *bool                                 `json:"registered,omitempty"`
	Status                    *PublisherBulkItemStatus              `json:"status,omitempty"`
	StitcherID                *int                                  `json:"stitcher_id,omitempty"`
	Tags                      []TagItem                             `json:"tags,omitempty"`
	UpgradeFailedReason       *PublisherBulkItemUpgradeFailedReason `json:"upgrade_failed_reason,omitempty"`
	UpgradeRequest            *bool                                 `json:"upgrade_request,omitempty"`
	UpgradeStatus             *PublisherBulkItemUpgradeStatus       `json:"upgrade_status,omitempty"`
}

func (o *PublisherBulkItem) GetAssessment() *PublisherBulkItemAssessment {
	if o == nil {
		return nil
	}
	return o.Assessment
}

func (o *PublisherBulkItem) GetCommonName() *string {
	if o == nil {
		return nil
	}
	return o.CommonName
}

func (o *PublisherBulkItem) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PublisherBulkItem) GetLbrokerconnect() *bool {
	if o == nil {
		return nil
	}
	return o.Lbrokerconnect
}

func (o *PublisherBulkItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PublisherBulkItem) GetPublisherUpgradeProfileID() *int {
	if o == nil {
		return nil
	}
	return o.PublisherUpgradeProfileID
}

func (o *PublisherBulkItem) GetRegistered() *bool {
	if o == nil {
		return nil
	}
	return o.Registered
}

func (o *PublisherBulkItem) GetStatus() *PublisherBulkItemStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PublisherBulkItem) GetStitcherID() *int {
	if o == nil {
		return nil
	}
	return o.StitcherID
}

func (o *PublisherBulkItem) GetTags() []TagItem {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *PublisherBulkItem) GetUpgradeFailedReason() *PublisherBulkItemUpgradeFailedReason {
	if o == nil {
		return nil
	}
	return o.UpgradeFailedReason
}

func (o *PublisherBulkItem) GetUpgradeRequest() *bool {
	if o == nil {
		return nil
	}
	return o.UpgradeRequest
}

func (o *PublisherBulkItem) GetUpgradeStatus() *PublisherBulkItemUpgradeStatus {
	if o == nil {
		return nil
	}
	return o.UpgradeStatus
}