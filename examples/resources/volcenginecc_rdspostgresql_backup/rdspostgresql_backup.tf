resource "volcenginecc_rdspostgresql_backup" "RDSPostgreSQLBackupDemo" {
  backup_method = "Logical"
  backup_scope  = "Database"
  instance_id   = "postgres-09xxxxxxxxx"
  backup_meta = [{
    db_name = "test_db"
  }]
  backup_description = "逻辑增量库级备份并下载测试"
  backup_policy = {
    data_incr_backup_periods   = "Wednesday"
    full_backup_period         = "Tuesday,Thursday,Saturday"
    hourly_incr_backup_enable  = true
    full_backup_time           = "01:00Z-02:00Z"
    increment_backup_frequency = 2
    backup_retention_period    = 14
    wal_log_space_limit_enable = true
  }
}