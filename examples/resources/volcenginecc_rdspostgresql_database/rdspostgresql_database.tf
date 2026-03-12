resource "volcenginecc_rdspostgresql_database" "RDSPostgreSQLDatabaseDemo" {
  instance_id        = "postgres-8d1fcxxxxxx"
  db_name            = "ccapi-terraform-1001"
  character_set_name = "utf8"
  owner              = "test-1"
  collate            = "C"
  c_type             = "C.UTF-8"
}