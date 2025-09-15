resource "volcenginecc_filenas_instance" "FileNASInstanceDemo" {
  file_system_name = "FileNASInstanceDemo"
  capacity = {
    total = 105
  }
  charge_type      = "PayAsYouGo"
  file_system_type = "Extreme"
  protocol_type    = "NFS"
  zone_id          = "cn-beijing-x"
}