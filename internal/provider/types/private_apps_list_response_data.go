// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type PrivateAppsListResponseData struct {
	AllowUnauthenticatedCors    types.Bool                       `tfsdk:"allow_unauthenticated_cors"`
	AllowURIBypass              types.Bool                       `tfsdk:"allow_uri_bypass"`
	AppID                       types.Int64                      `tfsdk:"app_id"`
	AppOption                   *AppOption                       `tfsdk:"app_option"`
	BypassUris                  []types.String                   `tfsdk:"bypass_uris"`
	ClientlessAccess            types.Bool                       `tfsdk:"clientless_access"`
	Host                        types.String                     `tfsdk:"host"`
	IsUserPortalApp             types.Bool                       `tfsdk:"is_user_portal_app"`
	Protocols                   []ProtocolResponseItem           `tfsdk:"protocols"`
	RealHost                    types.String                     `tfsdk:"real_host"`
	ServicePublisherAssignments []ServicePublisherAssignmentItem `tfsdk:"service_publisher_assignments"`
	Tags                        []TagItem                        `tfsdk:"tags"`
	TrustSelfSignedCerts        types.Bool                       `tfsdk:"trust_self_signed_certs"`
	UribypassHeaderValue        types.String                     `tfsdk:"uribypass_header_value"`
	UsePublisherDNS             types.Bool                       `tfsdk:"use_publisher_dns"`
}