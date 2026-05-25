resource "volcenginecc_vpn_vpn_gateway_route" "VpnGatewayRouteDemo" {
  destination_cidr_block = "192.168.0.0/25"
  next_hop_id            = "vgc-****"
  vpn_gateway_id         = "vgw-****"
}