resource "volcenginecc_storageebs_snapshot" "StorageEBSSnapshotDemo" {
  volume_id     = "vol-3wt6pip1dy3qu*****"
  snapshot_name = "StorageEBSSnapshotDemo"
  description   = "StorageEBSSnapshotDemo description"
  project_name  = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}