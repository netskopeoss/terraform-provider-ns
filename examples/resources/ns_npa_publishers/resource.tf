resource "ns_npa_publishers" "my_npapublishers" {
  lbrokerconnect                = false
  name                          = "Sherri Mohr"
  publisher_upgrade_profiles_id = 1
  tags = [
    {
      tag_id   = 3
      tag_name = "...my_tag_name..."
    },
  ]
}