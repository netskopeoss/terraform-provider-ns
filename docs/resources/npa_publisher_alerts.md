---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ns_npa_publisher_alerts Resource - terraform-provider-ns"
subcategory: ""
description: |-
  NPAPublisherAlerts Resource
---

# ns_npa_publisher_alerts (Resource)

NPAPublisherAlerts Resource

## Example Usage

```terraform
resource "ns_npa_publisher_alerts" "my_npapublisheralerts" {
  admin_users = [
    ["admin1@abc.com", "admin2@abc.com"],
  ]
  event_types = [
    "UPGRADE_STARTED",
  ]
  selected_users = "abc@xyz.com,def@xyz.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `admin_users` (List of String)
- `event_types` (List of String)
- `selected_users` (String)

### Read-Only

- `status` (String) must be one of ["success", "not found", "failure"]

