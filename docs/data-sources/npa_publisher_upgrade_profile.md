---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_publisher_upgrade_profile Data Source - terraform-provider-ns"
subcategory: ""
description: |-
  The NPA Publisher is a software package that enables private application
  connectivity between your data center and the Netskope cloud. It is a crucial
  component of Netskope’s Private Access (NPA) solution, which provides zero-trust
  network access (ZTNA) to private applications and data in hybrid IT environments.
  This data object queries and returns an upgrade profile by supplying 'external_id' of your profile
---

# ns_npa_publisher_upgrade_profile (Data Source)

The NPA Publisher is a software package that enables private application
connectivity between your data center and the Netskope cloud. It is a crucial
component of Netskope’s Private Access (NPA) solution, which provides zero-trust
network access (ZTNA) to private applications and data in hybrid IT environments.

This data object queries and returns an upgrade profile by supplying 'external_id' of your profile

## Example Usage

```terraform
data "ns_npa_publisher_upgrade_profile" "my_npapublisherupgradeprofile" {
  publisher_upgrade_profile_id = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `publisher_upgrade_profile_id` (Number) publisher upgrade profile external_id

### Read-Only

- `created_at` (String)
- `docker_tag` (String)
- `enabled` (Boolean)
- `frequency` (String)
- `name` (String)
- `next_update_time` (Number)
- `num_associated_publisher` (Number)
- `release_type` (String)
- `timezone` (String)
- `updated_at` (String)
- `upgrading_stage` (Number)
- `will_start` (Boolean)

