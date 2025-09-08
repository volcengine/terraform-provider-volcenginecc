resource "volcenginecc_vpc_subnet" "SubnetDemo" {
  vpc_id      = volcenginecc_vpc_vpc.TestTerraform_vpc_vpc.id
  zone_id     = var.pre_zone_id
  subnet_name = "subnetDemo"
  description = "subnetDemo description"
  cidr_block  = "192.xxx.x.0/24"
  tags        = [
    {
      key   = "env"
      value = "test"
    }
  ]
}