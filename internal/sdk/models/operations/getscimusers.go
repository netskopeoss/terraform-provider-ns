// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetSCIMUsersRequest struct {
	// key eq value pair like userName(Nothing but UPN), externalId(scim_externalid) and pass the value to be searched for the key passed.  Example: userName eq "upn1" OR externalId eq "User-Ext_id"
	Filter     *string `queryParam:"style=form,explode=true,name=filter"`
	StartIndex *int64  `queryParam:"style=form,explode=true,name=startIndex"`
	Count      *int64  `queryParam:"style=form,explode=true,name=count"`
}

func (o *GetSCIMUsersRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetSCIMUsersRequest) GetStartIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.StartIndex
}

func (o *GetSCIMUsersRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

type GetSCIMUsersSCIMUsersStatus string

const (
	GetSCIMUsersSCIMUsersStatusFailed            GetSCIMUsersSCIMUsersStatus = "failed"
	GetSCIMUsersSCIMUsersStatusFiveHundredAndOne GetSCIMUsersSCIMUsersStatus = "501"
)

func (e GetSCIMUsersSCIMUsersStatus) ToPointer() *GetSCIMUsersSCIMUsersStatus {
	return &e
}

func (e *GetSCIMUsersSCIMUsersStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "failed":
		fallthrough
	case "501":
		*e = GetSCIMUsersSCIMUsersStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSCIMUsersSCIMUsersStatus: %v", v)
	}
}

// GetSCIMUsersSCIMUsersResponse501ResponseBody - Internal error Failed
type GetSCIMUsersSCIMUsersResponse501ResponseBody struct {
	Status  *GetSCIMUsersSCIMUsersStatus `json:"status,omitempty"`
	Schemas *string                      `json:"schemas,omitempty"`
	Detail  *string                      `json:"detail,omitempty"`
}

func (o *GetSCIMUsersSCIMUsersResponse501ResponseBody) GetStatus() *GetSCIMUsersSCIMUsersStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetSCIMUsersSCIMUsersResponse501ResponseBody) GetSchemas() *string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *GetSCIMUsersSCIMUsersResponse501ResponseBody) GetDetail() *string {
	if o == nil {
		return nil
	}
	return o.Detail
}

// GetSCIMUsersSCIMUsersResponseResponseBody - Not authorized to execte the specific API.
type GetSCIMUsersSCIMUsersResponseResponseBody struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetSCIMUsersSCIMUsersResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type GetSCIMUsersStatus string

const (
	GetSCIMUsersStatusFourHundred GetSCIMUsersStatus = "400"
	GetSCIMUsersStatusFailed      GetSCIMUsersStatus = "failed"
)

func (e GetSCIMUsersStatus) ToPointer() *GetSCIMUsersStatus {
	return &e
}

func (e *GetSCIMUsersStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "400":
		fallthrough
	case "failed":
		*e = GetSCIMUsersStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSCIMUsersStatus: %v", v)
	}
}

// GetSCIMUsersSCIMUsersResponseBody - Failed
type GetSCIMUsersSCIMUsersResponseBody struct {
	Status      *GetSCIMUsersStatus `json:"status,omitempty"`
	Description *string             `json:"description,omitempty"`
	// schema of the resource
	Schemas []string `json:"schemas,omitempty"`
}

func (o *GetSCIMUsersSCIMUsersResponseBody) GetStatus() *GetSCIMUsersStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetSCIMUsersSCIMUsersResponseBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetSCIMUsersSCIMUsersResponseBody) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

// GetSCIMUsersName - Family_name and given_name for the User
type GetSCIMUsersName struct {
	// last_name of the SCIM User.
	FamilyName *string `json:"familyName,omitempty"`
	// first_name of the SCIM User.
	GivenName *string `json:"givenName,omitempty"`
}

func (o *GetSCIMUsersName) GetFamilyName() *string {
	if o == nil {
		return nil
	}
	return o.FamilyName
}

func (o *GetSCIMUsersName) GetGivenName() *string {
	if o == nil {
		return nil
	}
	return o.GivenName
}

type GetSCIMUsersEmails struct {
	// Email ID of the SCIM user
	Value   *string `json:"value,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
}

func (o *GetSCIMUsersEmails) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *GetSCIMUsersEmails) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}

type Resources struct {
	// SCIM User ID
	ID *string `json:"id,omitempty"`
	// UPN name of the SCIM User
	UserName *string `json:"userName,omitempty"`
	// Family_name and given_name for the User
	Name   *GetSCIMUsersName    `json:"name,omitempty"`
	Active *bool                `json:"active,omitempty"`
	Emails []GetSCIMUsersEmails `json:"emails,omitempty"`
	// Optional - Scim External ID
	ExternalID *string `json:"externalId,omitempty"`
}

func (o *Resources) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Resources) GetUserName() *string {
	if o == nil {
		return nil
	}
	return o.UserName
}

func (o *Resources) GetName() *GetSCIMUsersName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Resources) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *Resources) GetEmails() []GetSCIMUsersEmails {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *Resources) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

// GetSCIMUsersResponseBody - Successfully added/overwritten the User.
type GetSCIMUsersResponseBody struct {
	// schema of the resource
	Schemas      []string    `json:"schemas,omitempty"`
	TotalResults *int64      `json:"totalResults,omitempty"`
	StartIndex   *int64      `json:"startIndex,omitempty"`
	ItemsPerPage *int64      `json:"itemsPerPage,omitempty"`
	Resources    []Resources `json:"Resources,omitempty"`
}

func (o *GetSCIMUsersResponseBody) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *GetSCIMUsersResponseBody) GetTotalResults() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalResults
}

func (o *GetSCIMUsersResponseBody) GetStartIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.StartIndex
}

func (o *GetSCIMUsersResponseBody) GetItemsPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.ItemsPerPage
}

func (o *GetSCIMUsersResponseBody) GetResources() []Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

type GetSCIMUsersResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully added/overwritten the User.
	TwoHundredApplicationJSONObject *GetSCIMUsersResponseBody
	// Failed
	FourHundredApplicationJSONObject *GetSCIMUsersSCIMUsersResponseBody
	// Not authorized to execte the specific API.
	FourHundredAndThreeApplicationJSONObject *GetSCIMUsersSCIMUsersResponseResponseBody
	// Internal error Failed
	FiveHundredAndOneApplicationJSONObject *GetSCIMUsersSCIMUsersResponse501ResponseBody
}

func (o *GetSCIMUsersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSCIMUsersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSCIMUsersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSCIMUsersResponse) GetTwoHundredApplicationJSONObject() *GetSCIMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetSCIMUsersResponse) GetFourHundredApplicationJSONObject() *GetSCIMUsersSCIMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetSCIMUsersResponse) GetFourHundredAndThreeApplicationJSONObject() *GetSCIMUsersSCIMUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetSCIMUsersResponse) GetFiveHundredAndOneApplicationJSONObject() *GetSCIMUsersSCIMUsersResponse501ResponseBody {
	if o == nil {
		return nil
	}
	return o.FiveHundredAndOneApplicationJSONObject
}