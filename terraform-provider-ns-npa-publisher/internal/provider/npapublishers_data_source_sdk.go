// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"npa-publisher/internal/sdk/pkg/models/shared"
)

func (r *NPAPublishersDataSourceModel) RefreshFromGetResponse(resp *shared.PublishersGetResponse) {
	if resp.Data == nil {
		r.Data = nil
	} else {
		r.Data = &Data{}
		r.Data.Publishers = nil
		for _, publishersItem := range resp.Data.Publishers {
			var publishers1 Publisher
			publishers1.AppsCount = types.Int64Value(publishersItem.AppsCount)
			if publishersItem.Assessment.Assessment != nil {
				publishers1.Assessment.Assessment = &Assessment{}
				publishers1.Assessment.Assessment.EeeSupport = types.BoolValue(publishersItem.Assessment.Assessment.EeeSupport)
				publishers1.Assessment.Assessment.HddFree = types.StringValue(publishersItem.Assessment.Assessment.HddFree)
				publishers1.Assessment.Assessment.HddTotal = types.StringValue(publishersItem.Assessment.Assessment.HddTotal)
				publishers1.Assessment.Assessment.IPAddress = types.StringValue(publishersItem.Assessment.Assessment.IPAddress)
				publishers1.Assessment.Assessment.Latency = types.Int64Value(publishersItem.Assessment.Assessment.Latency)
				publishers1.Assessment.Assessment.Version = types.StringValue(publishersItem.Assessment.Assessment.Version)
			}
			if publishersItem.Assessment.PublisherAssessment2 != nil {
				publishers1.Assessment.PublisherAssessment2 = &PublisherAssessment2{}
			}
			publishers1.CommonName = types.StringValue(publishersItem.CommonName)
			publishers1.ConnectedApps = nil
			for _, v := range publishersItem.ConnectedApps {
				publishers1.ConnectedApps = append(publishers1.ConnectedApps, types.StringValue(v))
			}
			publishers1.Lbrokerconnect = types.BoolValue(publishersItem.Lbrokerconnect)
			publishers1.PublisherID = types.Int64Value(publishersItem.PublisherID)
			publishers1.PublisherName = types.StringValue(publishersItem.PublisherName)
			publishers1.PublisherUpgradeProfilesExternalID = types.Int64Value(publishersItem.PublisherUpgradeProfilesExternalID)
			publishers1.Registered = types.BoolValue(publishersItem.Registered)
			publishers1.Status = types.StringValue(publishersItem.Status)
			if publishersItem.StitcherID.Integer != nil {
				if publishersItem.StitcherID.Integer != nil {
					publishers1.StitcherID.Integer = types.Int64Value(*publishersItem.StitcherID.Integer)
				} else {
					publishers1.StitcherID.Integer = types.Int64Null()
				}
			}
			if publishersItem.StitcherID.PublisherStitcherID2 != nil {
				publishers1.StitcherID.PublisherStitcherID2 = &PublisherAssessment2{}
			}
			publishers1.Tags = nil
			for _, tagsItem := range publishersItem.Tags {
				var tags1 types.String
				tags1Result, _ := json.Marshal(tagsItem)
				tags1 = types.StringValue(string(tags1Result))
				publishers1.Tags = append(publishers1.Tags, tags1)
			}
			if publishersItem.UpgradeFailedReason.PublisherUpgradeFailedReason2 != nil {
				publishers1.UpgradeFailedReason.PublisherUpgradeFailedReason2 = &PublisherAssessment2{}
			}
			if publishersItem.UpgradeFailedReason.UpgradeFailedReason != nil {
				publishers1.UpgradeFailedReason.UpgradeFailedReason = &UpgradeFailedReason{}
				publishers1.UpgradeFailedReason.UpgradeFailedReason.Detail = types.StringValue(publishersItem.UpgradeFailedReason.UpgradeFailedReason.Detail)
				publishers1.UpgradeFailedReason.UpgradeFailedReason.ErrorCode = types.Int64Value(publishersItem.UpgradeFailedReason.UpgradeFailedReason.ErrorCode)
				publishers1.UpgradeFailedReason.UpgradeFailedReason.Timestamp = types.Int64Value(publishersItem.UpgradeFailedReason.UpgradeFailedReason.Timestamp)
				publishers1.UpgradeFailedReason.UpgradeFailedReason.Version = types.StringValue(publishersItem.UpgradeFailedReason.UpgradeFailedReason.Version)
			}
			publishers1.UpgradeRequest = types.BoolValue(publishersItem.UpgradeRequest)
			publishers1.UpgradeStatus.Upstat = types.StringValue(publishersItem.UpgradeStatus.Upstat)
			r.Data.Publishers = append(r.Data.Publishers, publishers1)
		}
	}
	if resp.Status != nil {
		r.Status = types.StringValue(*resp.Status)
	} else {
		r.Status = types.StringNull()
	}
	if resp.Total != nil {
		r.Total = types.Int64Value(int64(*resp.Total))
	} else {
		r.Total = types.Int64Null()
	}
}