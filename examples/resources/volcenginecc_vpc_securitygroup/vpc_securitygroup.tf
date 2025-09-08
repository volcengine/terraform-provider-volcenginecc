resource "volcenginecc_vpc_securitygroup" "SecurityGroupDemo" {
  vpc_id              = volcenginecc_vpc_vpc.TestTerraform_vpc_vpc.id
  security_group_name = "SecurityGroupDemo"
  description         = "SecurityGroupDemo description"
  project_name        = var.project_name
  ingress_permissions = [
    {
      description     = "进的规则1"
      policy          = "drop"
      port_end        = 1
      port_start      = 1
      priority        = 11
      protocol        = "tcp"
      cidr_ip         = "192.xxx.x.0/26"
      prefix_list_id  = ""
      source_group_id = ""
    }
  ]
  egress_permissions = [
    {
      description     = "出的规则1"
      policy          = "drop"
      port_end        = -1
      port_start      = -1
      priority        = 100
      protocol        = "all"
      cidr_ip         = "192.xxx.x.0/26"
      prefix_list_id  = ""
      source_group_id = ""
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}