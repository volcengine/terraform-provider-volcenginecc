resource "volcenginecc_clb_acl" "AclDemo" {
  acl_name     = "AclDemo"
  description  = "AclDemo Example"
  project_name = "default"
  acl_entries = [
    {
      description = "AclDemo description"
    entry = "1.1.2.2/32" },
    {
      description = "AclDemo description"
    entry = "2.2.2.2/32" }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}