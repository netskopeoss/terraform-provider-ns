resource "ns_npa_publishers_alerts_configuration" "my_npapublishersalertsconfiguration" {
  admin_users = [
    ["admin1@abc.com", "admin2@abc.com"],
  ]
  event_types = [
    "UPGRADE_FAILED",
  ]
  selected_users = "abc@xyz.com,def@xyz.com"
}