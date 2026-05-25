resource "volcenginecc_directconnect_gateway_route" "DirectConnectGatewayRouteDemo" {
  direct_connect_gateway_id = "dcg-****"
  next_hop_id               = "dcv-****"
  destination_cidr_block    = "192.168.1.0/28"
}