// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type UpgradePublisherResponse struct {
	AppsCount                  types.Int64                                  `tfsdk:"apps_count"`
	Assessment                 *UpgradePublisherResponseAssessment          `tfsdk:"assessment"`
	Capabilities               *PublisherResponseCapabilities               `tfsdk:"capabilities"`
	CommonName                 types.String                                 `tfsdk:"common_name"`
	ConnectedApps              []types.String                               `tfsdk:"connected_apps"`
	Lbrokerconnect             types.Bool                                   `tfsdk:"lbrokerconnect"`
	Name                       types.String                                 `tfsdk:"name"`
	PublisherID                types.Int64                                  `tfsdk:"publisher_id"`
	PublisherUpgradeProfilesID types.Int64                                  `tfsdk:"publisher_upgrade_profiles_id"`
	Registered                 types.Bool                                   `tfsdk:"registered"`
	Status                     types.String                                 `tfsdk:"status"`
	SticherPop                 types.String                                 `tfsdk:"sticher_pop"`
	StitcherID                 types.Int64                                  `tfsdk:"stitcher_id"`
	UpgradeFailedReason        *UpgradePublisherResponseUpgradeFailedReason `tfsdk:"upgrade_failed_reason"`
	UpgradeRequest             types.Bool                                   `tfsdk:"upgrade_request"`
	UpgradeStatus              *PublisherResponseUpgradeStatus              `tfsdk:"upgrade_status"`
}