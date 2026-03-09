resource "volcenginecc_kms_key" "KMSKeyDemo" {
  keyring_name           = "KMSKeyDemo"
  key_name               = "KMSKeyDemoKeyName"
  key_spec               = "SYMMETRIC_256"
  description            = "description KMSKeyDemo"
  key_usage              = "ENCRYPT_DECRYPT"
  protection_level       = "HSM"
  rotate_state           = "Enable"
  origin                 = "CloudKMS"
  multi_region           = false
  rotate_interval        = 900
  key_enable_operation   = 1
  key_rotation_operation = 1
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}