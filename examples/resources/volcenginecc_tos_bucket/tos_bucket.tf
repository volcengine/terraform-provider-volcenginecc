resource "volcenginecc_tos_bucket" "BucketDemo" {
  name                  = "bucket-demo"
  storage_class         = "STANDARD"
  enable_version_status = "Enabled"
  bucket_type           = "fns"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  policy = "{\"Version\":\"1.0\",\"Statement\":[{\"Sid\":\"f8fd\",\"Effect\":\"Allow\",\"Principal\":[\"*******\"],\"Action\":[\"tos:Get*\",\"tos:List*\",\"tos:HeadBucket\"],\"Resource\":[\"trn:tos:::************\",\"trn:tos:::************/*\"]}]}"
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
      lifecycle_rule_id = "bucketdemo"
      prefix            = "prefix"
    status = "Enabled" }
  ]
}