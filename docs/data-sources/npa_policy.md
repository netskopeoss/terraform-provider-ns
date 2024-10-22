---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_policy Data Source - terraform-provider-ns"
subcategory: ""
description: |-
  NPAPolicy DataSource
---

# ns_npa_policy (Data Source)

NPAPolicy DataSource

## Example Usage

```terraform
data "ns_npa_policy" "my_npapolicy" {
  id = "f45b0adb-9c35-4083-913c-9506eec1a765"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `rule_data` (Attributes) (see [below for nested schema](#nestedatt--rule_data))
- `rule_id` (String)
- `rule_name` (String)

<a id="nestedatt--rule_data"></a>
### Nested Schema for `rule_data`

Read-Only:

- `access_method` (List of String)
- `b_negate_net_location` (Boolean)
- `b_negate_src_countries` (Boolean)
- `classification` (List of String)
- `description` (String)
- `dlp_actions` (Attributes List) (see [below for nested schema](#nestedatt--rule_data--dlp_actions))
- `dlp_profile` (List of String)
- `external_dlp` (Boolean)
- `json_version` (Number)
- `match_criteria_action` (Attributes) (see [below for nested schema](#nestedatt--rule_data--match_criteria_action))
- `net_location_obj` (List of String)
- `organization_units` (List of String)
- `os` (List of String)
- `policy_type` (String) must be one of ["private-app"]
- `private_app_ids` (List of String)
- `private_app_tag_ids` (List of String)
- `private_app_tags` (List of String)
- `private_apps` (List of String)
- `private_apps_with_activities` (Attributes List) (see [below for nested schema](#nestedatt--rule_data--private_apps_with_activities))
- `show_dlp_profile_action_table` (Boolean)
- `src_countries` (List of String)
- `user_groups` (List of String)
- `user_type` (String) must be one of ["user"]
- `users` (List of String)
- `version` (Number)

<a id="nestedatt--rule_data--dlp_actions"></a>
### Nested Schema for `rule_data.dlp_actions`

Read-Only:

- `actions` (Attributes List) (see [below for nested schema](#nestedatt--rule_data--dlp_actions--actions))
- `dlp_profile` (String)

<a id="nestedatt--rule_data--dlp_actions--actions"></a>
### Nested Schema for `rule_data.dlp_actions.actions`

Read-Only:

- `action_name` (String)



<a id="nestedatt--rule_data--match_criteria_action"></a>
### Nested Schema for `rule_data.match_criteria_action`

Read-Only:

- `action_name` (String) must be one of ["allow", "block"]
- `template` (String)


<a id="nestedatt--rule_data--private_apps_with_activities"></a>
### Nested Schema for `rule_data.private_apps_with_activities`

Read-Only:

- `activities` (Attributes List) (see [below for nested schema](#nestedatt--rule_data--private_apps_with_activities--activities))
- `app_name` (String)

<a id="nestedatt--rule_data--private_apps_with_activities--activities"></a>
### Nested Schema for `rule_data.private_apps_with_activities.activities`

Read-Only:

- `activity` (String) must be one of ["any"]
- `list_of_constraints` (List of String)


