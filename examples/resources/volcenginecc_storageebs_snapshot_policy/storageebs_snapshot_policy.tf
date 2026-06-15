resource "volcenginecc_storageebs_snapshot_policy" "StorageEBSSnapshotPolicyDemo" {
  auto_snapshot_policy_name = "ccapi-test"
  time_points               = ["0", "1"]
  repeat_weekdays           = ["2", "5"]
  retention_days            = 7
  project_name              = "default"
  add_volume_ids            = ["vol-3x594jtg0g4cxxxxx", "vol-3x594j4dvp4i4txxxxx"]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}