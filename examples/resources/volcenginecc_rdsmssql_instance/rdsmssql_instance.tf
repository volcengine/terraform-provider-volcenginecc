resource "volcenginecc_rdsmssql_instance" "RDSMsSQLInstanceDemo" {
  node_spec              = "rds.mssql.3il.x8.medium.s1"
  zone_id                = "cn-beijing-a"
  subnet_id              = "subnet-1a0zgr5e7hslc8nvepxxxxx"
  db_engine_version      = "SQLServer_2019_Std"
  instance_type          = "Basic"
  storage_space          = 20
  vpc_id                 = "vpc-3nqt4kq87xn2893xxxxx"
  instance_name          = "RDSMsSQLInstanceDemo"
  super_account_password = "********"
  server_collation       = "Chinese_PRC_CI_AS"
  time_zone              = "China Standard Time"
  charge_info = {
    charge_type = "PostPaid"
  }
  project_name     = "default"
  maintenance_time = "18:00Z-21:59Z"
  allow_list_ids   = ["acl-03f197e136c843b29e47de74e9xxxxx"]
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}