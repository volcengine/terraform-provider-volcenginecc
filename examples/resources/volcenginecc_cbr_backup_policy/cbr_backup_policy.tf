resource "volcenginecc_cbr_backuppolicy" "CBRBackupPolicyDemo" {
  backup_type       = "INCREMENTAL"
  crontab           = "0 2,1,0 * * 1,2"
  enable_policy     = false
  name              = "CBRBackupPolicyDemo"
  retention_day     = 67
  retention_num_max = -1
  retention_num_min = 2
}