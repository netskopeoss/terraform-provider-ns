---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_policy_list Data Source - terraform-provider-ns"
subcategory: ""
description: |-
  NPAPolicyList DataSource
---

# ns_npa_policy_list (Data Source)

NPAPolicyList DataSource

## Example Usage

```terraform
data "ns_npa_policy_list" "my_npapolicylist" {
  filter    = "...my_filter..."
  limit     = 7
  offset    = 8
  sortby    = "...my_sortby..."
  sortorder = "...my_sortorder..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filter` (String) Query string based on query operaters
- `limit` (Number) Max number of policies to retrieve. Default will be all policies.
- `offset` (Number) The offset of the first policy in the list to retrieve.
- `sortby` (String) Sort retrieved policies by specified field. Default is policy id
- `sortorder` (String) Sort in either asc or desc order. The default is asc order

### Read-Only

- `data` (Attributes List) (see [below for nested schema](#nestedatt--data))

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `rule_data` (Attributes) (see [below for nested schema](#nestedatt--data--rule_data))
- `rule_id` (String)
- `rule_name` (String)

<a id="nestedatt--data--rule_data"></a>
### Nested Schema for `data.rule_data`

Read-Only:

- `access_method` (List of String)
- `b_negate_net_location` (Boolean)
- `b_negate_src_countries` (Boolean)
- `classification` (String)
- `dlp_actions` (Attributes List) (see [below for nested schema](#nestedatt--data--rule_data--dlp_actions))
- `external_dlp` (Boolean)
- `json_version` (Number)
- `match_criteria_action` (Attributes) (see [below for nested schema](#nestedatt--data--rule_data--match_criteria_action))
- `net_location_obj` (List of String)
- `organization_units` (List of String)
- `policy_type` (String) must be one of ["private-app"]
- `private_app_ids` (List of String)
- `private_app_tag_ids` (List of String)
- `private_app_tags` (List of String)
- `private_apps` (List of String)
- `private_apps_with_activities` (Attributes List) (see [below for nested schema](#nestedatt--data--rule_data--private_apps_with_activities))
- `show_dlp_profile_action_table` (Boolean)
- `src_countries` (List of String)
- `user_groups` (List of String)
- `user_type` (String) must be one of ["user"]
- `users` (List of String)
- `version` (Number)

<a id="nestedatt--data--rule_data--dlp_actions"></a>
### Nested Schema for `data.rule_data.dlp_actions`

Read-Only:

- `actions` (List of String)
- `dlp_profile` (String)


<a id="nestedatt--data--rule_data--match_criteria_action"></a>
### Nested Schema for `data.rule_data.match_criteria_action`

Read-Only:

- `action_name` (String) must be one of ["allow", "block"]


<a id="nestedatt--data--rule_data--private_apps_with_activities"></a>
### Nested Schema for `data.rule_data.private_apps_with_activities`

Read-Only:

- `activities` (Attributes List) (see [below for nested schema](#nestedatt--data--rule_data--private_apps_with_activities--activities))
- `app_name` (String)

<a id="nestedatt--data--rule_data--private_apps_with_activities--activities"></a>
### Nested Schema for `data.rule_data.private_apps_with_activities.app_name`

Read-Only:

- `activity` (String) must be one of ["any"]
- `list_of_constraints` (List of String)

