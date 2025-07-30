resource "volcenginecc_tos_bucket" "BucketDemo" {
  name                  = "bucket-demo"
  enable_version_status = "Enabled"
  tags                  = [
    {
      key   = "env"
      value = "test"
    }
  ]
  lifecycle_config = [
    {
      expiration = {
        date = "2027-01-18T00:00:00Z"
      }
      filter = {
        greater_than_include_equal = "Disabled"
        less_than_include_equal    = "Disabled"
        object_size_greater_than   = 123
        object_size_less_than      = 789
      }
      id     = "bucketdemo"
      prefix = "prefix"
      status = "Enabled"
    }
  ]
}