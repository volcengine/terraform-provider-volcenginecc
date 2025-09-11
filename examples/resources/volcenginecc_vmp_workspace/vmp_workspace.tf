resource "volcenginecc_vmp_workspace" "WorkspaceDemo" {
  username                  = "WorkspaceDemo"
  password                  = "***********"
  name                      = "WorkspaceDemo"
  description               = "WorkspaceDemo"
  delete_protection_enabled = false
  instance_type_id          = "vmp.standard.xxx"
  project_name              = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}