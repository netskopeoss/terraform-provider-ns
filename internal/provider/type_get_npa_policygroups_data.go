// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type GetNpaPolicygroupsData struct {
	CanBeEditedDeleted types.Int64  `tfsdk:"can_be_edited_deleted"`
	GroupID            types.Int64  `tfsdk:"group_id"`
	GroupName          types.String `tfsdk:"group_name"`
	GroupPinnedID      types.Int64  `tfsdk:"group_pinned_id"`
	GroupProdID        types.Int64  `tfsdk:"group_prod_id"`
	GroupType          types.Int64  `tfsdk:"group_type"`
	ModifyTime         types.String `tfsdk:"modify_time"`
	ModifyType         types.String `tfsdk:"modify_type"`
}