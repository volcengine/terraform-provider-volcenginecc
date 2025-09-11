resource "volcenginecc_transitrouter_transit_router" "TransitRouterDemo" {
  transit_router_name = "TransitRouterDemo"
  asn = 4200001111
  description = "TransitRouterDemo"
  project_name = "default"
  tags = [
    {
      key = "env"
      value = "test"    }
  ]
}