resource "volcenginecc_transitrouter_transit_router_route_entry" "TransitRouterRouteEntryDemo" {
  description                              = "Demo Example"
  destination_cidr_block                   = "192.168.1.101/*"
  transit_router_route_entry_name          = "test-Attachmentkua"
  transit_router_route_table_id            = "tr-rtb-mijcn***"
  transit_router_route_entry_next_hop_id   = "tr-attach-13fs****"
  transit_router_route_entry_next_hop_type = "Attachment"
  transit_router_route_entry_type          = "Propagated"
}