// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/netskope/terraform-provider-ns/internal/provider/types"
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/shared"
)

func (r *PolicyGroupListDataSourceModel) RefreshFromSharedNpaPolicygroupResponse(resp *shared.NpaPolicygroupResponse) {
	if resp != nil {
		r.Data = []tfTypes.NpaPolicygroupResponseItem{}
		if len(r.Data) > len(resp.Data) {
			r.Data = r.Data[:len(resp.Data)]
		}
		for dataCount, dataItem := range resp.Data {
			var data1 tfTypes.NpaPolicygroupResponseItem
			data1.CanBeEditedDeleted = types.StringPointerValue(dataItem.CanBeEditedDeleted)
			data1.GroupName = types.StringPointerValue(dataItem.GroupName)
			data1.ID = types.StringPointerValue(dataItem.ID)
			if dataCount+1 > len(r.Data) {
				r.Data = append(r.Data, data1)
			} else {
				r.Data[dataCount].CanBeEditedDeleted = data1.CanBeEditedDeleted
				r.Data[dataCount].GroupName = data1.GroupName
				r.Data[dataCount].ID = data1.ID
			}
		}
	}
}
