resource "volcenginecc_rdsmysql_instance" "RDSMySQLInstanceDemo" {
  deletion_protection = "Disabled"
  db_engine_version   = "MySQL_5_7"
  nodes = [
    {
      zone_id   = "cn-beijing-a"
      node_spec = "rds.mysql.c.s.1c2g"
    node_type = "Primary" },
    {
      zone_id   = "cn-beijing-a"
      node_spec = "rds.mysql.c.s.1c2g"
    node_type = "Secondary" },
    {
      zone_id   = "cn-beijing-a"
      node_spec = "rds.mysql.c.s.1c2g"
    node_type = "ReadOnly" }
  ]
  storage_type           = "CloudESSD_PL0"
  storage_space          = 100
  instance_type          = "DoubleNode"
  vpc_id                 = "vpc-rrco37ovjq4gv0x58xxxxx"
  subnet_id              = "subnet-rrwqhg3qzxfkv0x57xxxxx"
  instance_name          = "RDSMySQLInstanceDemo-按量计费"
  super_account_name     = "test_account"
  super_account_password = "***********"
  lower_case_table_names = "1"
  db_time_zone           = "UTC +08:00"
  charge_detail = {
    charge_type = "PostPaid"
    auto_renew  = false
    number      = 1
  }
  allow_list_ids = [
    "acl-8cde5e16f44143788234ca4489xxxxx",
    "acl-31f6053bd6be4cff88c1b205d20xxxxx"
  ]
  port = 3306
  maintenance_window = {
    day_kind = "Week"
    day_of_week = [
      "Monday"
    ]
    maintenance_time = "18:00Z-21:59Z"
  }
  auto_storage_scaling_config = {
    enable_storage_auto_scale = true
    storage_threshold         = 10
    storage_upper_bound       = 3000
  }
  project_name = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}