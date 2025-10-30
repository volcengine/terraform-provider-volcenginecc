resource "volcenginecc_directconnect_direct_connect_gateway" "DirectConnectDirectConnectGatewayDemo" {
  description                 = "DirectConnectDirectConnectGatewayDemo-Description"
  direct_connect_gateway_name = "DirectConnectDirectConnectGatewayDemo"
  enable_ipv_6                = false
  project_name                = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}