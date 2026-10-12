resource "volcenginecc_tidb_instance" "example" {
  display_name = "ccapi-tidb-instance-full"
  auto_scaling = {
    min_rcu = 2000
    max_rcu = 4000
  }
  encryption_config = {
    enhanced_encryption_enabled = true
  }
  project_name           = "default"
  high_availability_type = "REGIONAL"
  region_id              = "cn-beijing"
  service_plan           = "Essential"
  vpc_id                 = "vpc-xxx"
  zones = [{
    zone_id   = "zone-xxx"
    subnet_id = "subnet-xxxx"
  }]
  automated_backup_policy = {
    start_time     = "01:00"
    retention_days = 1
  }
  tags = [{
    key   = "env"
    value = "full"
    }, {
    key   = "team"
    value = "ccapi"
  }]

}
