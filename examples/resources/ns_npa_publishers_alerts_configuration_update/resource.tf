resource "ns_npa_publishers_alerts_configuration_update" "my_npapublishersalertsconfigurationupdate" {
  admin_users = [
    ["admin1@abc.com", "admin2@abc.com"],
  ]
  event_types = [
    "UPGRADE_SUCCEEDED",
  ]
  selected_users = "abc@xyz.com,def@xyz.com"
}