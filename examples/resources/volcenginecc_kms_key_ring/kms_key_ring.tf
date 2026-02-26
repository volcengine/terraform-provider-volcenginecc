resource "volcenginecc_kms_key_ring" "kmskeyringDemo" {
  description  = "this is a test update"
  keyring_name = "ccapi-test"
  keyring_type = "CustomKeyring"
  project_name = "default"
}