resource "volcenginecc_storageebs_snapshot_group" "StorageEBSSnapshotGroupDemo" {
  name         = "ccapi-test"
  project_name = "default"
  description  = "this-is_test"
  instance_id  = "i-yekg7jsw00bw8xxxxx"
  volume_ids   = ["vol-3x1q75ir0t4qbgxxxxx", "vol-3x1q75ir0t4qbgxxxxx", "vol-3x1q75ir0t4qbgxxxxx", "vol-3x1q75ir0t4qbghxxxxx"]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}