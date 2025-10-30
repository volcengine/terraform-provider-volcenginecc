resource "volcenginecc_transitrouter_transit_router_route_table" "TransitRouterTransitRouterRouteTableDemo" {
  transit_router_id               = "tr-mj7mc0paq******"
  description                     = "TransitRouterTransitRouterRouteTableDemo-Description"
  transit_router_route_table_name = "TransitRouterTransitRouterRouteTableDemo"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}