resource "volcenginecc_rdspostgresql_schema" "RDSPostgreSQLSchemaDemo" {
  instance_id = "postgres-8d1fcxxxxx"
  db_name     = "ccapi-terraform-1001"
  schema_name = "ccapi-terraform-1"
  owner       = "test"
}