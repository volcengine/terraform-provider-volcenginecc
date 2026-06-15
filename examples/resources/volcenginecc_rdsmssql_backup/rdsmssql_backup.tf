resource "volcenginecc_rdsmssql_backup" "RDSMsSQLBackupDemo" {
  instance_id = "mssql-099bda9abfaf"
  backup_type = "Full"
  backup_meta = [
    {
      db_name = "user"
    }
  ]
}