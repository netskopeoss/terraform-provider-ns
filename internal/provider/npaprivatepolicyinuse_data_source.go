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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NPAPrivatePolicyInUseDataSource{}
var _ datasource.DataSourceWithConfigure = &NPAPrivatePolicyInUseDataSource{}

func NewNPAPrivatePolicyInUseDataSource() datasource.DataSource {
	return &NPAPrivatePolicyInUseDataSource{}
}

// NPAPrivatePolicyInUseDataSource is the data source implementation.
type NPAPrivatePolicyInUseDataSource struct {
	client *sdk.TerraformProviderNs
}

// NPAPrivatePolicyInUseDataSourceModel describes the data model.
type NPAPrivatePolicyInUseDataSourceModel struct {
	Data []tfTypes.ResponseBody `tfsdk:"data"`
	Ids  []types.String         `tfsdk:"ids"`
}

// Metadata returns the data source type name.
func (r *NPAPrivatePolicyInUseDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_npa_private_policy_in_use"
}

// Schema defines the schema for the data source.
func (r *NPAPrivatePolicyInUseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NPAPrivatePolicyInUse DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"token": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["success", "not found"]`,
						},
					},
				},
				Description: `successful operation`,
			},
			"ids": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (r *NPAPrivatePolicyInUseDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *NPAPrivatePolicyInUseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NPAPrivatePolicyInUseDataSourceModel
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

	request := *data.ToOperationsGetNPAPolicyInUseRequestBody()
	res, err := r.client.GetNPAPolicyInUse(ctx, request)
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
	if !(res.ResponseBodies != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromResponseBody(res.ResponseBodies)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}