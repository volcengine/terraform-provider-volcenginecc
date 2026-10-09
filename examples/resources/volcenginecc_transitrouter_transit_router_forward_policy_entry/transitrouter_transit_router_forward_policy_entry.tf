resource "volcenginecc_transitrouter_transit_router_forward_policy_entry" "example" {
  transit_router_forward_policy_table_id = "policy_table_xxxxx"
  transit_router_route_table_id          = "route_table-xxxxx"
  source_cidr_block                      = "172.16.0.0/12"
  destination_cidr_block                 = "192.168.0.0/16"
  priority                               = 100
  description                            = "ccapi transit router forward policy entry full"

}
