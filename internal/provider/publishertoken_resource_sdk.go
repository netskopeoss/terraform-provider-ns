// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"ns/internal/sdk/pkg/models/operations"
)

func (r *PublisherTokenResourceModel) ToCreateSDKType() *operations.PostInfrastructurePublishersPublisherIDRegistrationTokenRequest {
	publisherID := int(r.PublisherID.ValueInt64())
	out := operations.PostInfrastructurePublishersPublisherIDRegistrationTokenRequest{
		PublisherID: publisherID,
	}
	return &out
}

func (r *PublisherTokenResourceModel) RefreshFromCreateResponse(resp *operations.PostInfrastructurePublishersPublisherIDRegistrationTokenData) {
	r.Token = types.StringValue(resp.Token)
}