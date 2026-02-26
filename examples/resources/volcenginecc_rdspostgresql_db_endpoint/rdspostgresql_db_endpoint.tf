resource "volcenginecc_rdspostgresql_db_endpoint" "RdsPostgresqlDbEndpointDemo" {
  endpoint_name   = "ccapi-test-1"
  endpoint_type   = "Custom"
  instance_id     = "postgres-9dxxxxxd"
  nodes           = "Primary"
  read_write_mode = "ReadWrite"
}