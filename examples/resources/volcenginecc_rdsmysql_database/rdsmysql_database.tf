resource "volcenginecc_rdsmysql_database" "DatabaseDemo" {
  character_set_name = "utf8"
  database_privileges = [
    {
      account_name      = "a***"
      account_privilege = "ReadOnly"
    host = "%" }
  ]
  description = "Demo Example"
  instance_id = "mysql-779***"
  name        = "rdstest"
}