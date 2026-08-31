resource "volcenginecc_vpc_ipv4_gateway" "example" {
  ipv_4_gateway_name = "ccapi-ipv4-gw-full"
  description        = "ccapi ipv4 gateway full case"
  vpc_id             = "vpc-xxxxx"
}
