resource "volcenginecc_hbase_instance" "HBaseInstanceDemo" {
  region_id            = "cn-beijing"
  multi_az             = false
  zone_id              = "cn-beijing-a"
  instance_name        = "HBaseInstanceDemo"
  engine_version       = "HBase_2.0"
  instance_type        = "Standard"
  master_spec          = "hbase.x1.medium"
  rs_count             = 2
  rs_spec              = "hbase.x1.large"
  storage_capacity     = 20000
  vpc_id               = "vpc-rrco37ovjq4gv0x5xxxxx"
  project_name         = "default"
  charge_type          = "PrePaid"
  purchase_months      = 1
  auto_renew           = false
  subnet_id            = "subnet-rrwqhg3qzxfkv0x57xxxxx"
  deletion_protection  = "disabled"
  enable_cloud_storage = false
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  storage_type = "HdfsSsd"
  enable_auth  = true
}