resource "volcenginecc_rdspostgresql_instance" "RDSPostgreSQLInstanceDemo" {
  db_engine_version = "PostgreSQL_14"
  node_info = [
    {
      zone_id   = "cn-beijing-a"
      node_spec = "rds.postgres.1c2g"
      node_type = "Primary"
    },
    {
      zone_id   = "cn-beijing-a"
      node_spec = "rds.postgres.1c2g"
      node_type = "Secondary"
    }
  ]
  storage_type = "LocalSSD"
  vpc_id       = "vpc-***"
  subnet_id    = "subnet-***"
  charge_detail = {
    charge_type = "PostPaid"
  }
}