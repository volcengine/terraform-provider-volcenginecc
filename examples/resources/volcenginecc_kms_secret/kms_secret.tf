resource "volcenginecc_kms_secret" "KMSSecretDemo" {
  secret_name    = "secret-ccapi-Generic"
  version_name   = "ccapi-v1"
  project_name   = "default"
  description    = "description test ccapi Generic"
  secret_type    = "Generic"
  encryption_key = "trn:kms:cn-beijing:********:keyrings/ccapi-text/keys/ccapi-terraform-1101"
  secret_value = jsonencode({
    "key1" = "value1"
    "key2" = "value2"
  })
}