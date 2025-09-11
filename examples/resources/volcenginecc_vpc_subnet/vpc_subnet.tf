resource "volcenginecc_vpc_subnet" "SubnetDemo" {
  vpc_id      = "vpc-xxxx"
  zone_id     = "cn-beijing"
  subnet_name = "subnetDemo"
  description = "subnetDemo description"
  cidr_block  = "192.168.xx.0/24"
  tags        = [
    {
      key   = "env"
      value = "test"
    }
  ]
}