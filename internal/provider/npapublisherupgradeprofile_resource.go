// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/netskope/terraform-provider-ns/internal/sdk"
	"github.com/netskope/terraform-provider-ns/internal/sdk/models/operations"
	"strconv"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NPAPublisherUpgradeProfileResource{}
var _ resource.ResourceWithImportState = &NPAPublisherUpgradeProfileResource{}

func NewNPAPublisherUpgradeProfileResource() resource.Resource {
	return &NPAPublisherUpgradeProfileResource{}
}

// NPAPublisherUpgradeProfileResource defines the resource implementation.
type NPAPublisherUpgradeProfileResource struct {
	client *sdk.TerraformProviderNs
}

// NPAPublisherUpgradeProfileResourceModel describes the resource data model.
type NPAPublisherUpgradeProfileResourceModel struct {
	CreatedAt                 types.String `tfsdk:"created_at"`
	DockerTag                 types.String `tfsdk:"docker_tag"`
	Enabled                   types.Bool   `tfsdk:"enabled"`
	Frequency                 types.String `tfsdk:"frequency"`
	Name                      types.String `tfsdk:"name"`
	NextUpdateTime            types.Int64  `tfsdk:"next_update_time"`
	NumAssociatedPublisher    types.Int64  `tfsdk:"num_associated_publisher"`
	PublisherUpgradeProfileID types.Int64  `tfsdk:"publisher_upgrade_profile_id"`
	ReleaseType               types.String `tfsdk:"release_type"`
	Timezone                  types.String `tfsdk:"timezone"`
	UpdatedAt                 types.String `tfsdk:"updated_at"`
	UpgradingStage            types.Int64  `tfsdk:"upgrading_stage"`
	WillStart                 types.Bool   `tfsdk:"will_start"`
}

func (r *NPAPublisherUpgradeProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_npa_publisher_upgrade_profile"
}

func (r *NPAPublisherUpgradeProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "The NPA Publisher is a software package that enables private application\nconnectivity between your data center and the Netskope cloud. It is a crucial\ncomponent of Netskope’s Private Access (NPA) solution, which provides zero-trust\nnetwork access (ZTNA) to private applications and data in hybrid IT environments.\n\nThis resource creates a publisher upgrade profile.    \n",
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"docker_tag": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Docker Tag of the release version you wish to install. \` + "\n" +
					`Docker Tag for releases can be obtained from: \` + "\n" +
					`` + "`" + `api/v2/infrastructure/publishers/releases` + "`" + `` + "\n" +
					``,
			},
			"enabled": schema.BoolAttribute{
				Required: true,
				MarkdownDescription: `Is this updgrade profile enabled.` + "\n" +
					`* ` + "`" + `true` + "`" + ` - Enabled` + "\n" +
					`* ` + "`" + `false` + "`" + ` - Disabled` + "\n" +
					``,
			},
			"frequency": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `Frequency of updates. This frequency is in a CRON format. \` + "\n" +
					`┌───────────── minute (0–59) \` + "\n" +
					`│ ┌───────────── hour (0–23) \` + "\n" +
					`│ │ ┌───────────── day of the month (1–31) \` + "\n" +
					`│ │ │ ┌───────────── month (1–12) (Leave as *) \` + "\n" +
					`│ │ │ │ ┌───────────── day of the week (MON, TUE, WED, THU, FRI, SAT, SUN) \` + "\n" +
					`0 0 1 * TUE => (Midnight, Weekly, Tuesday)` + "\n" +
					``,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"next_update_time": schema.Int64Attribute{
				Computed: true,
			},
			"num_associated_publisher": schema.Int64Attribute{
				Computed: true,
			},
			"publisher_upgrade_profile_id": schema.Int64Attribute{
				Computed:    true,
				Description: `publisher upgrade profile external_id`,
			},
			"release_type": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `This is the Release Type that is to be installed. \` + "\n" +
					`Release Type for releases can be obtained from: \` + "\n" +
					`` + "`" + `api/v2/infrastructure/publishers/releases` + "`" + `` + "\n" +
					`` + "\n" +
					`must be one of ["Beta", "Latest", "Latest-1", "Latest-2"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Beta",
						"Latest",
						"Latest-1",
						"Latest-2",
					),
				},
			},
			"timezone": schema.StringAttribute{
				Required: true,
				MarkdownDescription: `The timezone for which the upgrade triggers. \` + "\n" +
					`Please see enum for accepted values.` + "\n" +
					`` + "\n" +
					`must be one of ["Africa/Cairo", "Africa/Casablanca", "Africa/Johannesburg", "Africa/Nairobi", "America/Argentina/Buenos_Aires", "America/Caracas", "America/Godthab", "America/Lima", "America/Mazatlan", "America/Santiago", "America/Tijuana", "Asia/Almaty", "Asia/Baghdad", "Asia/Baku", "Asia/Calcutta", "Asia/Dhaka", "Asia/Harbin", "Asia/Jakarta", "Asia/Jerusalem", "Asia/Kabul", "Asia/Karachi", "Asia/Kathmandu", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Muscat", "Asia/Rangoon", "Asia/Taipei", "Asia/Tehran", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Cape_Verde", "Australia/Adelaide", "Australia/Brisbane", "Australia/Darwin", "Australia/Hobart", "Australia/Perth", "Australia/Sydney", "Brazil/East", "Canada/Atlantic", "Canada/Central", "Canada/Newfoundland", "Canada/Saskatchewan", "Europe/Amsterdam", "Europe/Athens", "Europe/Copenhagen", "Europe/Helsinki", "Europe/London", "Europe/Minsk", "Europe/Moscow", "Europe/Paris", "Europe/Prague", "Europe/Sarajevo", "Japan", "Mexico/General", "Pacific/Auckland", "Pacific/Fiji", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Samoa", "Pacific/Tongatapu", "US/Alaska", "US/Arizona", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Mountain", "US/Pacific"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Africa/Cairo",
						"Africa/Casablanca",
						"Africa/Johannesburg",
						"Africa/Nairobi",
						"America/Argentina/Buenos_Aires",
						"America/Caracas",
						"America/Godthab",
						"America/Lima",
						"America/Mazatlan",
						"America/Santiago",
						"America/Tijuana",
						"Asia/Almaty",
						"Asia/Baghdad",
						"Asia/Baku",
						"Asia/Calcutta",
						"Asia/Dhaka",
						"Asia/Harbin",
						"Asia/Jakarta",
						"Asia/Jerusalem",
						"Asia/Kabul",
						"Asia/Karachi",
						"Asia/Kathmandu",
						"Asia/Krasnoyarsk",
						"Asia/Kuala_Lumpur",
						"Asia/Muscat",
						"Asia/Rangoon",
						"Asia/Taipei",
						"Asia/Tehran",
						"Asia/Vladivostok",
						"Asia/Yakutsk",
						"Asia/Yerevan",
						"Atlantic/Azores",
						"Atlantic/Cape_Verde",
						"Australia/Adelaide",
						"Australia/Brisbane",
						"Australia/Darwin",
						"Australia/Hobart",
						"Australia/Perth",
						"Australia/Sydney",
						"Brazil/East",
						"Canada/Atlantic",
						"Canada/Central",
						"Canada/Newfoundland",
						"Canada/Saskatchewan",
						"Europe/Amsterdam",
						"Europe/Athens",
						"Europe/Copenhagen",
						"Europe/Helsinki",
						"Europe/London",
						"Europe/Minsk",
						"Europe/Moscow",
						"Europe/Paris",
						"Europe/Prague",
						"Europe/Sarajevo",
						"Japan",
						"Mexico/General",
						"Pacific/Auckland",
						"Pacific/Fiji",
						"Pacific/Guadalcanal",
						"Pacific/Guam",
						"Pacific/Samoa",
						"Pacific/Tongatapu",
						"US/Alaska",
						"US/Arizona",
						"US/East-Indiana",
						"US/Eastern",
						"US/Hawaii",
						"US/Mountain",
						"US/Pacific",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"upgrading_stage": schema.Int64Attribute{
				Computed: true,
			},
			"will_start": schema.BoolAttribute{
				Computed: true,
			},
		},
	}
}

func (r *NPAPublisherUpgradeProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.TerraformProviderNs)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.TerraformProviderNs, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NPAPublisherUpgradeProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *NPAPublisherUpgradeProfileResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedPublisherUpgradeProfilePostRequest()
	res, err := r.client.NPAPublisherUpgradeProfile.CreateNPAPublisherUpgradeProfile(ctx, request)
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
	if !(res.PublisherUpgradeProfileResponse != nil && res.PublisherUpgradeProfileResponse.Data != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPublisherUpgradeProfileResponseData(res.PublisherUpgradeProfileResponse.Data)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NPAPublisherUpgradeProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *NPAPublisherUpgradeProfileResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var publisherUpgradeProfileID int
	publisherUpgradeProfileID = int(data.PublisherUpgradeProfileID.ValueInt64())

	request := operations.GetNPAPublisherUpgradeProfileRequest{
		PublisherUpgradeProfileID: publisherUpgradeProfileID,
	}
	res, err := r.client.NPAPublisherUpgradeProfile.GetNPAPublisherUpgradeProfile(ctx, request)
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
	if !(res.PublisherUpgradeProfileGetResponse != nil && res.PublisherUpgradeProfileGetResponse.Data != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPublisherUpgradeProfileGetResponseData(res.PublisherUpgradeProfileGetResponse.Data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NPAPublisherUpgradeProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *NPAPublisherUpgradeProfileResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var publisherUpgradeProfileID int
	publisherUpgradeProfileID = int(data.PublisherUpgradeProfileID.ValueInt64())

	publisherUpgradeProfilePutRequest := *data.ToSharedPublisherUpgradeProfilePutRequest()
	request := operations.UpdateNPAPublisherUpgradeProfileRequest{
		PublisherUpgradeProfileID:         publisherUpgradeProfileID,
		PublisherUpgradeProfilePutRequest: publisherUpgradeProfilePutRequest,
	}
	res, err := r.client.NPAPublisherUpgradeProfile.UpdateNPAPublisherUpgradeProfile(ctx, request)
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
	if !(res.PublisherUpgradeProfileResponse != nil && res.PublisherUpgradeProfileResponse.Data != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPublisherUpgradeProfileResponseData(res.PublisherUpgradeProfileResponse.Data)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NPAPublisherUpgradeProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *NPAPublisherUpgradeProfileResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	var publisherUpgradeProfileID int
	publisherUpgradeProfileID = int(data.PublisherUpgradeProfileID.ValueInt64())

	request := operations.DeleteNPAPublisherUpgradeProfileRequest{
		PublisherUpgradeProfileID: publisherUpgradeProfileID,
	}
	res, err := r.client.NPAPublisherUpgradeProfile.DeleteNPAPublisherUpgradeProfile(ctx, request)
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

}

func (r *NPAPublisherUpgradeProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	publisherUpgradeProfileID, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid ID", fmt.Sprintf("ID must be an integer but was %s", req.ID))
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("publisher_upgrade_profile_id"), int64(publisherUpgradeProfileID))...)
}