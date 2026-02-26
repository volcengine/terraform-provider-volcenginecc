resource "volcenginecc_rdsmysql_backup" "rdsmysqlbackupDemo" {
  instance_id   = "mysql-9e2c59****"
  backup_method = "Logical"
  backup_meta = [
    {
      database = "test",
      tables   = []

    }
  ]
}