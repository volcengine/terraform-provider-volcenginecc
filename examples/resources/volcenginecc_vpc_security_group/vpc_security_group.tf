resource "volcenginecc_vpc_security_group" "SecurityGroupDemo" {
  vpc_id              = "vpc-xxxx"
  security_group_name = "SecurityGroupDemo"
  description         = "SecurityGroupDemo description"
  project_name        = "default"
  ingress_permissions = [
    {
      description     = "进的规则"
      policy          = "drop"
      port_end        = 1
      port_start      = 1
      priority        = 11
      protocol        = "tcp"
      cidr_ip         = "192.168.xx.0/26"
      prefix_list_id  = ""
      source_group_id = ""
    }
  ]
  egress_permissions = [
    {
      description     = "出的规则"
      policy          = "drop"
      port_end        = -1
      port_start      = -1
      priority        = 100
      protocol        = "all"
      cidr_ip         = "192.168.xx.0/26"
      prefix_list_id  = "list-xxx"
      source_group_id = "group-xxx"
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}