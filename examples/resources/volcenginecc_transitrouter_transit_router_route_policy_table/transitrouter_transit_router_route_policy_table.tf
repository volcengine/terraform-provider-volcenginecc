resource "volcenginecc_transitrouter_transit_router_route_policy_table" "example" {
  transit_router_id                      = "router_id-xxx"
  transit_router_route_policy_table_name = "ccapi-route-policy-table-full"
  description                            = "ccapi transit router route policy table"
  direction                              = "Out"

}
