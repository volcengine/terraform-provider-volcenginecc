resource "volcenginecc_kms_key" "KMSKeyDemo" {
  keyring_name     = "KMSKeyDemo"
  key_name         = "KMSKeyDemoKeyName"
  key_spec         = "SYMMETRIC_256"
  description      = "description KMSKeyDemo"
  key_usage        = "ENCRYPT_DECRYPT"
  protection_level = "HSM"
  rotate_state     = "Enable"
  origin           = "CloudKMS"
  multi_region     = false
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}