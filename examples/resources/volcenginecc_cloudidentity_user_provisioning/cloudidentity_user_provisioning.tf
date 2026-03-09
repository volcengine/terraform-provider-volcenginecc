resource "volcenginecc_cloudidentity_userprovisioning" "CloudIdentityUserProvisioningDemo" {
  target_id                = "2103612***"
  principal_type           = "Group"
  principal_id             = "8990752****"
  duplication_strategy     = "KeepBoth"
  deletion_strategy        = "Delete"
  identity_source_strategy = "BindConflictUser"
  description              = "ccapi-test"
  duplication_suffix       = "-cctest"
}