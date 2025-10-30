resource "volcenginecc_vpc_network_acl" "NetworkAclDemo" {
  vpc_id           = "vpc-rrco37ovjq4gv0x58xxxxxx"
  network_acl_name = "NetworkAclDemo"
  description      = "NetworkAclDemo-Description"
  project_name     = "default"
  ingress_acl_entries = [
    {
      cidr_ip                = "10.0.1.0/24"
      description            = "默认规则"
      network_acl_entry_name = "test-entries"
      policy                 = "drop"
      port                   = "80/80"
    protocol = "tcp" }
  ]
  egress_acl_entries = [
    {
      cidr_ip                = "0.0.0.0/0"
      description            = "默认规则"
      network_acl_entry_name = ""
      policy                 = "accept"
      port                   = "-1/-1"
    protocol = "all" }
  ]
  resources = [
    {
    resource_id = "subnet-3nrjlvvxo4gsg931ebxxxxxx" }
  ]
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}