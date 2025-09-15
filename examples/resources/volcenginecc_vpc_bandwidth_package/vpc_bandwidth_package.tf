resource "volcenginecc_vpc_bandwidth_package" "BandwidthPackageDemo" {
  period      = 1
  period_unit = 1
  protocol    = "Dual-stack"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  bandwidth_package_name = "BandwidthPackageDemo"
  description            = "BandwidthPackageDemo"
  isp                    = "BGP"
  billing_type           = 2
  bandwidth              = 2
  project_name           = "default"
  eip_addresses = [
    {
      allocation_id = "eip-2f90y2a8sczcw4f4pzzx1xxxx"
    eip_address = "101.126.*.143" }
  ]
}