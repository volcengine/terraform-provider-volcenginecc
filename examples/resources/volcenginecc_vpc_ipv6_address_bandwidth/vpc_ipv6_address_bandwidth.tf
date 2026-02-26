resource "volcenginecc_vpc_ipv6_address_bandwidth" "VPCIpv6AddressBandwidthDemo" {
  bandwidth            = 200
  billing_type         = 2
  ipv_6_address        = "2406:d440:10a:****:92b:****:90b6:4f09"
  project_name         = "default"
  bandwidth_package_id = "bwp-1vm41dmikjr451j8exxxxx"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}