resource "volcenginecc_vedbm_backup" "vedbmbackupDemo" {
  instance_id             = "vedbm-i34lvm3j***"
  backup_type             = "Full"
  backup_method           = "Physical"
  backup_time             = "00:00Z-02:00Z"
  full_backup_period      = "Sunday"
  backup_retention_period = 11

}