// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type PublisherBulkItem struct {
	Assessment                *PublisherBulkItemAssessment `tfsdk:"assessment"`
	CommonName                types.String                 `tfsdk:"common_name"`
	ID                        types.Int64                  `tfsdk:"id"`
	Lbrokerconnect            types.Bool                   `tfsdk:"lbrokerconnect"`
	Name                      types.String                 `tfsdk:"name"`
	PublisherUpgradeProfileID types.Int64                  `tfsdk:"publisher_upgrade_profile_id"`
	Registered                types.Bool                   `tfsdk:"registered"`
	Status                    types.String                 `tfsdk:"status"`
	StitcherID                types.Int64                  `tfsdk:"stitcher_id"`
	Tags                      []TagItem                    `tfsdk:"tags"`
	UpgradeFailedReason       *PublisherBulkItemAssessment `tfsdk:"upgrade_failed_reason"`
	UpgradeRequest            types.Bool                   `tfsdk:"upgrade_request"`
	UpgradeStatus             *PublisherBulkItemAssessment `tfsdk:"upgrade_status"`
}