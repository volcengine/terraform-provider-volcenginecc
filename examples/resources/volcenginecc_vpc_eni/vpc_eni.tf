resource "volcenginecc_vpc_eni" "EniDemo" {
  network_interface_name = "EniDemo"
  subnet_id              = "vpc_subnet-xxxx"
  security_group_ids     = ["vpc_security_group-xxxx"]
  instance_id            = "i-ye498lwge85i3z3kxxxx"
  project_name           = "default"
  primary_ip_address = {
    private_ip_address = "192.168.x.130"
    associated_elastic_ip = {
      allocation_id = "eip-2f80zqjduo6ps4f4pzzcxxxxx"
    }
  }
  secondary_private_ip_address_count = 2
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}