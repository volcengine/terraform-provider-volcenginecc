resource "volcenginecc_vpc_traffic_mirror_filter" "VPCTrafficMirrorFilterDemo" {
  description                = "ccapi-terraform-test"
  traffic_mirror_filter_name = "test-vpc-traffic-mirror-filter-demo"
  project_name               = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}