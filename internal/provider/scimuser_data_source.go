// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/speakeasy/terraform-provider-terraform/internal/provider/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SCIMUserDataSource{}
var _ datasource.DataSourceWithConfigure = &SCIMUserDataSource{}

func NewSCIMUserDataSource() datasource.DataSource {
	return &SCIMUserDataSource{}
}

// SCIMUserDataSource is the data source implementation.
type SCIMUserDataSource struct {
	client *sdk.TerraformProviderNs
}

// SCIMUserDataSourceModel describes the data model.
type SCIMUserDataSourceModel struct {
	Active     types.Bool                    `tfsdk:"active"`
	Emails     []tfTypes.Emails              `tfsdk:"emails"`
	ExternalID types.String                  `tfsdk:"external_id"`
	ID         types.String                  `tfsdk:"id"`
	Meta       *tfTypes.CreateSCIMGroupsMeta `tfsdk:"meta"`
	Name       *tfTypes.Name                 `tfsdk:"name"`
	Schemas    []types.String                `tfsdk:"schemas"`
	UserName   types.String                  `tfsdk:"user_name"`
}

// Metadata returns the data source type name.
func (r *SCIMUserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scim_user"
}

// Schema defines the schema for the data source.
func (r *SCIMUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SCIMUser DataSource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed: true,
			},
			"emails": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"primary": schema.BoolAttribute{
							Computed: true,
						},
						"value": schema.StringAttribute{
							Computed:    true,
							Description: `Email ID of the SCIM user`,
						},
					},
				},
			},
			"external_id": schema.StringAttribute{
				Computed:    true,
				Description: `Optional - Scim External ID`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `SCIM Used ID should be specified`,
			},
			"meta": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"resource_type": schema.StringAttribute{
						Computed:    true,
						Description: `resource type User/Group.`,
					},
				},
				Description: `meta information`,
			},
			"name": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"family_name": schema.StringAttribute{
						Computed:    true,
						Description: `last_name of the SCIM User.`,
					},
					"given_name": schema.StringAttribute{
						Computed:    true,
						Description: `first_name of the SCIM User.`,
					},
				},
				Description: `Family_name and given_name for the User`,
			},
			"schemas": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `schema of the resource`,
			},
			"user_name": schema.StringAttribute{
				Computed:    true,
				Description: `UPN name of the SCIM User`,
			},
		},
	}
}

func (r *SCIMUserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SCIMUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SCIMUserDataSourceModel
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

	id := data.ID.ValueString()
	request := operations.GetSCIMUsersByIDRequest{
		ID: id,
	}
	res, err := r.client.ScimUsers.GetSCIMUsersByID(ctx, request)
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
	if res.TwoHundredApplicationJSONObject == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsGetSCIMUsersByIDResponseBody(res.TwoHundredApplicationJSONObject)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}