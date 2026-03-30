resource "volcenginecc_vmp_workspace" "WorkspaceDemo" {
  auth_type                 = "BearerToken"
  bearer_token              = "M3cSN7gssM09-6wO8vdqo_xxxxxxxx"
  delete_protection_enabled = false
  description               = "test workspace"
  instance_type_id          = "vmp.standard.30d"
  name                      = "terraform_test_BearerToken"
  project_name              = "default"
  public_access_enabled     = true
  public_query_bandwidth    = 2
  public_write_bandwidth    = 50
  search_latency_offset     = "32s"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}