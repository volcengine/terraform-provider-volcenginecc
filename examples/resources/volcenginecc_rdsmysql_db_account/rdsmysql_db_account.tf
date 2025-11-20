resource "volcenginecc_rdsmysql_db_account" "RDSMySQLDBAccountDemo" {
  account_desc     = "RDSMySQLDBAccountDemo desc"
  instance_id      = "mysql-a5b5caexxxxx"
  account_name     = "RDSMySQLDBAccountDemo"
  account_password = "********"
  account_type     = "Normal"
  account_privileges = [
    {
      account_privilege        = "Custom"
      account_privilege_detail = ["CREATE", "DROP", "REFERENCES", "INDEX"]
    db_name = "test" },
    {
      account_privilege        = "Global"
      account_privilege_detail = ["PROCESS", "REPLICATION SLAVE", "REPLICATION CLIENT", "DROP", "ALTER"]
    db_name = "" }
  ]
  host                    = "%"
  table_column_privileges = []
}