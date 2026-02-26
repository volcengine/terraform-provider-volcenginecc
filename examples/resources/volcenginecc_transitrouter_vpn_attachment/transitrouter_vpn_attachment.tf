resource "volcenginecc_transitrouter_vpn_attachment" "TransitRouterVpnAttachmentDemo" {
  description = "test"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  transit_router_attachment_name = "ccapi-test-1"
  transit_router_id              = "tr-mjl8zkxxxxxsmt1boobol4"
  transit_router_route_table_id  = "tr-rtb-mjlxxxxxo5smt1bnewrhh"
  vpn_connection_id              = "vgc-ij0yxxxxx474o8cux0n08t"
  zone_id                        = "cn-beijing-a"
}