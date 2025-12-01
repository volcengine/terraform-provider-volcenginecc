resource "volcenginecc_vepfs_instance" "VEPFSInstanceDemo" {
  file_system_name = "VEPFSInstanceDemo"
  zone_id          = "cn-beijing-a"
  charge_type      = "PayAsYouGo"
  file_system_type = "VePFS"
  store_type       = "Advance_100"
  protocol_type    = "VePFS"
  project_name     = "default"
  capacity         = 8
  vpc_id           = "vpc-3nqt4kq87xn28931eclxxxxx"
  subnet_id        = "subnet-1a0zgr5e7hslc8nvepkxxxxxx"
  version_number   = "1.4.0"
  enable_restripe  = true
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}