resource "volcenginecc_alb_acl" "AlbAclDemo" {
  acl_name     = "AlbAclDemo"
  project_name = "default"
  acl_entries = [
    {
      description = "AlbAclDemo description"
    entry = "1.1.1.1/32" }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}