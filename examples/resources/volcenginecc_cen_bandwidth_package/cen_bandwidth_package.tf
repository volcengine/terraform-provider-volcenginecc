resource "volcenginecc_cen_bandwidth_package" "CENBandwidthPackageDemo" {
  bandwidth                  = 3
  billing_type               = 4
  cen_bandwidth_package_name = "ccapi-test"
  cen_ids = [
    "cen-2v73nw1h8a03k6x7exxxxx"
  ]
  description                    = "this is a cen test"
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  project_name                   = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}