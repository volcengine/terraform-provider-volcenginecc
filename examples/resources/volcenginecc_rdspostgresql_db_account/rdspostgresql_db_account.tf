resource "volcenginecc_rdspostgresql_db_account" "RDSPostgreSQLDBAccountNormalDemo" {
  instance_id          = "postgres-xxxxxxx"
  account_name         = "ccapi-test-Normal"
  account_password     = "**********"
  account_type         = "Normal"
  account_privileges   = "Inherit,Login,CreateRole,CreateDB"
  not_allow_privileges = ["DDL"]
}