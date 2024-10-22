// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
)

func (r *NPAPublishersDataSourceModel) RefreshFromSharedPublisherResponseData(resp *shared.PublisherResponseData) {
	if resp != nil {
		if resp.Assessment == nil {
			r.Assessment = types.StringNull()
		} else {
			assessmentResult, _ := json.Marshal(resp.Assessment)
			r.Assessment = types.StringValue(string(assessmentResult))
		}
		r.CommonName = types.StringPointerValue(resp.CommonName)
		r.Lbrokerconnect = types.BoolPointerValue(resp.Lbrokerconnect)
		r.Name = types.StringPointerValue(resp.Name)
		if resp.PublisherID != nil {
			r.PublisherID = types.Int64Value(int64(*resp.PublisherID))
		} else {
			r.PublisherID = types.Int64Null()
		}
		if resp.PublisherUpgradeProfileID != nil {
			r.PublisherUpgradeProfileID = types.Int64Value(int64(*resp.PublisherUpgradeProfileID))
		} else {
			r.PublisherUpgradeProfileID = types.Int64Null()
		}
		r.Registered = types.BoolPointerValue(resp.Registered)
		if resp.Status != nil {
			r.Status = types.StringValue(string(*resp.Status))
		} else {
			r.Status = types.StringNull()
		}
		if resp.StitcherID != nil {
			r.StitcherID = types.Int64Value(int64(*resp.StitcherID))
		} else {
			r.StitcherID = types.Int64Null()
		}
	}
}
