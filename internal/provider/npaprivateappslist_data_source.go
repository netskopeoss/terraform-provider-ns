// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/netskope/terraform-provider-ns/internal/provider/types"
	"github.com/netskope/terraform-provider-ns/internal/sdk"
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NPAPrivateAppsListDataSource{}
var _ datasource.DataSourceWithConfigure = &NPAPrivateAppsListDataSource{}

func NewNPAPrivateAppsListDataSource() datasource.DataSource {
	return &NPAPrivateAppsListDataSource{}
}

// NPAPrivateAppsListDataSource is the data source implementation.
type NPAPrivateAppsListDataSource struct {
	client *sdk.TerraformProviderNs
}

// NPAPrivateAppsListDataSourceModel describes the data model.
type NPAPrivateAppsListDataSourceModel struct {
	Data   []tfTypes.PrivateAppsGetResponse `tfsdk:"data"`
	Fields types.String                     `tfsdk:"fields"`
}

// Metadata returns the data source type name.
func (r *NPAPrivateAppsListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_npa_private_apps_list"
}

// Schema defines the schema for the data source.
func (r *NPAPrivateAppsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NPAPrivateAppsList DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"allow_unauthenticated_cors": schema.BoolAttribute{
									Computed: true,
								},
								"allow_uri_bypass": schema.BoolAttribute{
									Computed: true,
								},
								"app_id": schema.Int64Attribute{
									Computed: true,
								},
								"app_name": schema.StringAttribute{
									Computed: true,
								},
								"app_option": schema.SingleNestedAttribute{
									Computed:   true,
									Attributes: map[string]schema.Attribute{},
								},
								"bypass_uris": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"clientless_access": schema.BoolAttribute{
									Computed: true,
								},
								"host": schema.StringAttribute{
									Computed: true,
								},
								"is_user_portal_app": schema.BoolAttribute{
									Computed: true,
								},
								"protocols": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.Int64Attribute{
												Computed: true,
											},
											"port": schema.StringAttribute{
												Computed: true,
											},
											"service_id": schema.Int64Attribute{
												Computed: true,
											},
											"transport": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"real_host": schema.StringAttribute{
									Computed: true,
								},
								"service_publisher_assignments": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"primary": schema.BoolAttribute{
												Computed: true,
											},
											"publisher_id": schema.Int64Attribute{
												Computed: true,
											},
											"reachability": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"error_code": schema.Int64Attribute{
														Computed: true,
													},
													"error_string": schema.StringAttribute{
														Computed: true,
													},
													"reachable": schema.BoolAttribute{
														Computed: true,
													},
												},
											},
											"service_id": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
								"tags": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"tag_id": schema.Int64Attribute{
												Computed: true,
											},
											"tag_name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"trust_self_signed_certs": schema.BoolAttribute{
									Computed: true,
								},
								"uribypass_header_value": schema.StringAttribute{
									Computed: true,
								},
								"use_publisher_dns": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["success", "not found"]`,
						},
						"total": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
				Description: `successful operation`,
			},
			"fields": schema.StringAttribute{
				Optional:    true,
				Description: `Return values only from specified fields`,
			},
		},
	}
}

func (r *NPAPrivateAppsListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.TerraformProviderNs)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.TerraformProviderNs, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NPAPrivateAppsListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NPAPrivateAppsListDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	fields := new(string)
	if !data.Fields.IsUnknown() && !data.Fields.IsNull() {
		*fields = data.Fields.ValueString()
	} else {
		fields = nil
	}
	request := operations.ListNPAPrivateAppsRequest{
		Fields: fields,
	}
	res, err := r.client.ListNPAPrivateApps(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.PrivateAppsGetResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrivateAppsGetResponse(res.PrivateAppsGetResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}