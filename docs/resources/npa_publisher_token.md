---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_publisher_token Resource - terraform-provider-ns"
subcategory: ""
description: |-
  The NPA Publisher is a software package that enables private application
  connectivity between your data center and the Netskope cloud. It is a crucial
  component of Netskope’s Private Access (NPA) solution, which provides zero-trust
  network access (ZTNA) to private applications and data in hybrid IT environments.
  This resource supports the creation and retrival of a registration token.
---

# ns_npa_publisher_token (Resource)

The NPA Publisher is a software package that enables private application
connectivity between your data center and the Netskope cloud. It is a crucial 
component of Netskope’s Private Access (NPA) solution, which provides zero-trust 
network access (ZTNA) to private applications and data in hybrid IT environments.

This resource supports the creation and retrival of a registration token.

## Example Usage

```terraform
resource "ns_npa_publisher_token" "my_npapublishertoken" {
  publisher_id = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `publisher_id` (Number) publisher id. Requires replacement if changed.

### Read-Only

- `token` (String)

