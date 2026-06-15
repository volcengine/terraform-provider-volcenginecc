resource "volcenginecc_rdspostgresql_instance_ssl" "RDSPostgreSQLInstanceSSLDemo" {
  reload_ssl_certificate = true
  instance_id            = "postgres-60xxxx5ed9"
  force_encryption       = true
}