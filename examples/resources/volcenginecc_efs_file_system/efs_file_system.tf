resource "volcenginecc_efs_file_system" "EFSFileSystemDemo" {
  file_system_name    = "cc-test"
  description         = "description"
  charge_type         = "PayAsYouGo"
  zone_id             = "cn-beijing-a"
  instance_type       = "Premium"
  performance_density = "Premium_125"
  project_name        = "default"
  performance = {
    bandwidth_mode        = "Provisioned"
    provisioned_bandwidth = 300,
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}