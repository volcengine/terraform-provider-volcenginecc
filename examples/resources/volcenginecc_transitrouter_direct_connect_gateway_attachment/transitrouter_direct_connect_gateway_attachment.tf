resource "volcenginecc_transitrouter_direct_connect_gateway_attachment" "DirectConnectGatewayAttachmentDemo" {
  direct_connect_gateway_id      = "dcg-1714vlpplhxc1f*****"
  transit_router_attachment_name = "test"
  transit_router_id              = "tr-2d69n1k0brncw58ozfdh6y6xq***"
  description                    = "test"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}