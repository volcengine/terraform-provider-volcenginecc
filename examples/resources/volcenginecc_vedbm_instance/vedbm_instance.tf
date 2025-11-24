resource "volcenginecc_vedbm_instance" "VEDBMInstanceDemo" {
  instance_name          = "VEDBMInstanceDemo"
  db_engine_version      = "MySQL_8_0"
  db_minor_version       = "3.0"
  node_spec              = "vedb.mysql.x4.medium"
  zone_ids               = "cn-beijing-a"
  node_number            = 2
  vpc_id                 = "vpc-1a1vgeoxxxccg8nvepjykjnuw"
  subnet_id              = "subnet-3nrd6xxx3log0931ech3re2r"
  port                   = 3306
  super_account_name     = "username"
  super_account_password = "*******"
  time_zone              = "UTC +08:00"
  lower_case_table_names = "1"
  project_name           = "default"
  tags = [
    {
      key = "dev"
    value = "test" }
  ]
  deletion_protection = "enabled"
  number              = 1
  charge_detail = {
    charge_type = "PostPaid"
  }
}