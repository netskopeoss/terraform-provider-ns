// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"npa-publisher/internal/sdk/pkg/models/shared"
)

func (r *NPAPublisherReleasesDataSourceModel) RefreshFromGetResponse(resp *shared.PublishersReleaseGetResponse) {
	r.Data = nil
	for _, dataItem := range resp.Data {
		var data1 ReleaseItem
		if dataItem.DockerTag != nil {
			data1.DockerTag = types.StringValue(*dataItem.DockerTag)
		} else {
			data1.DockerTag = types.StringNull()
		}
		if dataItem.Name != nil {
			data1.Name = types.StringValue(*dataItem.Name)
		} else {
			data1.Name = types.StringNull()
		}
		if dataItem.Version != nil {
			data1.Version = types.StringValue(*dataItem.Version)
		} else {
			data1.Version = types.StringNull()
		}
		r.Data = append(r.Data, data1)
	}
	if resp.Status != nil {
		r.Status = types.StringValue(string(*resp.Status))
	} else {
		r.Status = types.StringNull()
	}
}