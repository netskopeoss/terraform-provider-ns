// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Assessment struct {
	EeeSupport types.Bool   `tfsdk:"eee_support"`
	HddFree    types.String `tfsdk:"hdd_free"`
	HddTotal   types.String `tfsdk:"hdd_total"`
	IPAddress  types.String `tfsdk:"ip_address"`
	Latency    types.Int64  `tfsdk:"latency"`
	Version    types.String `tfsdk:"version"`
}