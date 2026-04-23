resource "volcenginecc_tos_bucket_inventory" "TOSBucketInventoryDemo" {
  bucket_name = "ccapi-test"
  destination = {
    tos_bucket_destination = {
      account_id = "21xxxxxxxxx"
      bucket     = "ccapi-test"
      format     = "CSV"
      prefix     = "tos_bucket_inventory"
      role       = "TosArchiveTOSxxxxxx"
    }
  }
  filter = {
    prefix = "sdfsaf"
  }
  included_object_versions = "Current"
  inventory_id             = "test"
  is_enabled               = true
  is_un_compressed         = false
  optional_fields = {
    field = ["StorageClass", "IntelligentTieringAccessTier"]
  }
  schedule = {
    frequency = "Weekly"
  }
}