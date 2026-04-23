resource "volcenginecc_tos_bucket_encryption" "TOSBucketEncryptionDemo" {
  kms_master_key_id = "trn:kms:cn-beijing:21xxxxxxxxxxx:keyrings/ccapi-text/keys/terratest-kms-key-u-xxxxx"
  name              = "ccapi-test"
  sse_algorithm     = "kms"
}