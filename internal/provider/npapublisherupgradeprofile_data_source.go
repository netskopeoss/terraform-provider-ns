// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/netskope/terraform-provider-ns/internal/sdk"
	"github.com/netskope/terraform-provider-ns/internal/sdk/pkg/models/operations"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NPAPublisherUpgradeProfileDataSource{}
var _ datasource.DataSourceWithConfigure = &NPAPublisherUpgradeProfileDataSource{}

func NewNPAPublisherUpgradeProfileDataSource() datasource.DataSource {
	return &NPAPublisherUpgradeProfileDataSource{}
}

// NPAPublisherUpgradeProfileDataSource is the data source implementation.
type NPAPublisherUpgradeProfileDataSource struct {
	client *sdk.SDK
}

// NPAPublisherUpgradeProfileDataSourceModel describes the data model.
type NPAPublisherUpgradeProfileDataSourceModel struct {
	DockerTag   types.String `tfsdk:"docker_tag"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Frequency   types.String `tfsdk:"frequency"`
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ReleaseType types.String `tfsdk:"release_type"`
	Timezone    types.String `tfsdk:"timezone"`
}

// Metadata returns the data source type name.
func (r *NPAPublisherUpgradeProfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_npa_publisher_upgrade_profile"
}

// Schema defines the schema for the data source.
func (r *NPAPublisherUpgradeProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NPAPublisherUpgradeProfile DataSource",

		Attributes: map[string]schema.Attribute{
			"docker_tag": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"frequency": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.Int64Attribute{
				Required:    true,
				Description: `publisher upgrade profile id`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"release_type": schema.StringAttribute{
				Computed: true,
			},
			"timezone": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *NPAPublisherUpgradeProfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NPAPublisherUpgradeProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *NPAPublisherUpgradeProfileDataSourceModel
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

	upgradeProfileID := int(data.ID.ValueInt64())
	request := operations.GetInfrastructurePublisherupgradeprofilesUpgradeProfileIDRequest{
		UpgradeProfileID: upgradeProfileID,
	}
	res, err := r.client.NPAPublisherUpgradeProfiles.Read(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.PublisherUpgradeProfileResponse == nil || res.PublisherUpgradeProfileResponse.Data == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.PublisherUpgradeProfileResponse.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}