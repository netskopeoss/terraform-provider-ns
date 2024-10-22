---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_publishers_list Data Source - terraform-provider-ns"
subcategory: ""
description: |-
  The NPA Publisher is a software package that enables private application
  connectivity between your data center and the Netskope cloud. It is a crucial
  component of Netskope’s Private Access (NPA) solution, which provides zero-trust
  network access (ZTNA) to private applications and data in hybrid IT environments.
  This data source supports the list of all Publisher objects.
  Features may require additional licensing, please work with account team to enable.
---

# ns_npa_publishers_list (Data Source)

The NPA Publisher is a software package that enables private application
connectivity between your data center and the Netskope cloud. It is a crucial 
component of Netskope’s Private Access (NPA) solution, which provides zero-trust 
network access (ZTNA) to private applications and data in hybrid IT environments.

This data source supports the list of all Publisher objects.

Features may require additional licensing, please work with account team to enable.

## Example Usage

```terraform
data "ns_npa_publishers_list" "my_npapublisherslist" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `data` (Attributes) (see [below for nested schema](#nestedatt--data))
- `status` (String)
- `total` (Number)

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `publishers` (Attributes List) (see [below for nested schema](#nestedatt--data--publishers))

<a id="nestedatt--data--publishers"></a>
### Nested Schema for `data.publishers`

Read-Only:

- `apps_count` (Number)
- `assessment` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--assessment))
- `common_name` (String)
- `connected_apps` (List of String)
- `lbrokerconnect` (Boolean)
- `publisher_id` (Number)
- `publisher_name` (String)
- `publisher_upgrade_profiles_external_id` (Number)
- `registered` (Boolean)
- `status` (String)
- `stitcher_id` (Number)
- `upgrade_failed_reason` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--upgrade_failed_reason))
- `upgrade_request` (Boolean)
- `upgrade_status` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--upgrade_status))

<a id="nestedatt--data--publishers--assessment"></a>
### Nested Schema for `data.publishers.assessment`

Read-Only:

- `assessment` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--assessment--assessment))
- `two` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--assessment--two))

<a id="nestedatt--data--publishers--assessment--assessment"></a>
### Nested Schema for `data.publishers.assessment.two`

Read-Only:

- `eee_support` (Boolean)
- `hdd_free` (String)
- `hdd_total` (String)
- `ip_address` (String)
- `latency` (Number)
- `version` (String)


<a id="nestedatt--data--publishers--assessment--two"></a>
### Nested Schema for `data.publishers.assessment.two`



<a id="nestedatt--data--publishers--upgrade_failed_reason"></a>
### Nested Schema for `data.publishers.upgrade_failed_reason`

Read-Only:

- `two` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--upgrade_failed_reason--two))
- `upgrade_failed_reason` (Attributes) (see [below for nested schema](#nestedatt--data--publishers--upgrade_failed_reason--upgrade_failed_reason))

<a id="nestedatt--data--publishers--upgrade_failed_reason--two"></a>
### Nested Schema for `data.publishers.upgrade_failed_reason.upgrade_failed_reason`


<a id="nestedatt--data--publishers--upgrade_failed_reason--upgrade_failed_reason"></a>
### Nested Schema for `data.publishers.upgrade_failed_reason.upgrade_failed_reason`

Read-Only:

- `detail` (String)
- `error_code` (Number)
- `timestamp` (Number)
- `version` (String)



<a id="nestedatt--data--publishers--upgrade_status"></a>
### Nested Schema for `data.publishers.upgrade_status`

Read-Only:

- `upstat` (String)


