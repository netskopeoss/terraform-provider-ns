---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_publisher_upgrade_profile Resource - terraform-provider-ns"
subcategory: ""
description: |-
  The NPA Publisher is a software package that enables private application
  connectivity between your data center and the Netskope cloud. It is a crucial
  component of Netskope’s Private Access (NPA) solution, which provides zero-trust
  network access (ZTNA) to private applications and data in hybrid IT environments.
  This resource creates a publisher upgrade profile.
---

# ns_npa_publisher_upgrade_profile (Resource)

The NPA Publisher is a software package that enables private application
connectivity between your data center and the Netskope cloud. It is a crucial
component of Netskope’s Private Access (NPA) solution, which provides zero-trust
network access (ZTNA) to private applications and data in hybrid IT environments.

This resource creates a publisher upgrade profile.

## Example Usage

```terraform
resource "ns_npa_publisher_upgrade_profile" "my_npapublisherupgradeprofile" {
  docker_tag   = 8690
  enabled      = true
  frequency    = "0 0 1 * TUE"
  name         = "My Upgrade Profile"
  release_type = "Latest"
  timezone     = "US/Eastern"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `docker_tag` (String) Docker Tag of the release version you wish to install. \
Docker Tag for releases can be obtained from: \
`api/v2/infrastructure/publishers/releases`
- `enabled` (Boolean) Is this updgrade profile enabled.
* `true` - Enabled
* `false` - Disabled
- `frequency` (String) Frequency of updates. This frequency is in a CRON format. \
┌───────────── minute (0–59) \
│ ┌───────────── hour (0–23) \
│ │ ┌───────────── day of the month (1–31) \
│ │ │ ┌───────────── month (1–12) (Leave as *) \
│ │ │ │ ┌───────────── day of the week (MON, TUE, WED, THU, FRI, SAT, SUN) \
0 0 1 * TUE => (Midnight, Weekly, Tuesday)
- `name` (String)
- `release_type` (String) This is the Release Type that is to be installed. \
Release Type for releases can be obtained from: \
`api/v2/infrastructure/publishers/releases`

must be one of ["Beta", "Latest", "Latest-1", "Latest-2"]
- `timezone` (String) The timezone for which the upgrade triggers. \
Please see enum for accepted values.

must be one of ["Africa/Cairo", "Africa/Casablanca", "Africa/Johannesburg", "Africa/Nairobi", "America/Argentina/Buenos_Aires", "America/Caracas", "America/Godthab", "America/Lima", "America/Mazatlan", "America/Santiago", "America/Tijuana", "Asia/Almaty", "Asia/Baghdad", "Asia/Baku", "Asia/Calcutta", "Asia/Dhaka", "Asia/Harbin", "Asia/Jakarta", "Asia/Jerusalem", "Asia/Kabul", "Asia/Karachi", "Asia/Kathmandu", "Asia/Krasnoyarsk", "Asia/Kuala_Lumpur", "Asia/Muscat", "Asia/Rangoon", "Asia/Taipei", "Asia/Tehran", "Asia/Vladivostok", "Asia/Yakutsk", "Asia/Yerevan", "Atlantic/Azores", "Atlantic/Cape_Verde", "Australia/Adelaide", "Australia/Brisbane", "Australia/Darwin", "Australia/Hobart", "Australia/Perth", "Australia/Sydney", "Brazil/East", "Canada/Atlantic", "Canada/Central", "Canada/Newfoundland", "Canada/Saskatchewan", "Europe/Amsterdam", "Europe/Athens", "Europe/Copenhagen", "Europe/Helsinki", "Europe/London", "Europe/Minsk", "Europe/Moscow", "Europe/Paris", "Europe/Prague", "Europe/Sarajevo", "Japan", "Mexico/General", "Pacific/Auckland", "Pacific/Fiji", "Pacific/Guadalcanal", "Pacific/Guam", "Pacific/Samoa", "Pacific/Tongatapu", "US/Alaska", "US/Arizona", "US/East-Indiana", "US/Eastern", "US/Hawaii", "US/Mountain", "US/Pacific"]

### Read-Only

- `created_at` (String)
- `next_update_time` (Number)
- `num_associated_publisher` (Number)
- `publisher_upgrade_profile_id` (Number) publisher upgrade profile external_id
- `updated_at` (String)
- `upgrading_stage` (Number)
- `will_start` (Boolean)

## Import

Import is supported using the following syntax:

```shell
terraform import ns_npa_publisher_upgrade_profile.my_ns_npa_publisher_upgrade_profile "0"
```