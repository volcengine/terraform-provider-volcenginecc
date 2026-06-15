resource "volcenginecc_transitrouter_transit_router_bandwidth_package" "TransitRouterBandwidthPackageDemo" {
  bandwidth                      = 3
  billing_type                   = 1
  description                    = "test"
  line_operator                  = ""
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  period                         = 1
  period_unit                    = "Month"
  project_name                   = "default"
  tags = [{
    key   = "env"
    value = "test"
  }]
  transit_router_bandwidth_package_name = "ccapi-test"
}