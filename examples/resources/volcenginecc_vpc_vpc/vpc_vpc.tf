resource "volcenginecc_vpc_vpc" "VPCDemo" {
  cidr_block            = "192.168.0.0/24"
  support_ipv_4_gateway = true
  enable_ipv_6          = false
  vpc_name              = "vpc-demo"
  description           = "VpcDemo Example"
  dns_servers = [
    "12.3.x.x"
  ]
  associate_cens = [
    {
      cen_id       = "cen-3re8cx4vwdibk5zsk2xxxx"
      cen_owner_id = "200000xxx"
      cen_status   = "Attaching"
    }
  ]
  nat_gateway_ids = [
    "ngw-2d6tp1y8zq41s58ozfdxxx"
  ]
  route_table_ids = [
    "vtb-29mkf8ft83l6o1e1hgixxx"
  ]
  security_group_ids = [
    "sg-29mkx39jb80741e1hgj2bxxx"
  ]
  project_name = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}