resource "volcenginecc_vpn_vpn_gateway" "VpnVpnGatewayDemo" {
  bandwidth           = 5
  billing_type        = 2
  description         = "this is a test description"
  dual_tunnel_enabled = false
  ip_stack_type       = "ipv4_only"
  ip_version          = "ipv4"
  vpn_gateway_name    = "VpnVpnGatewayDemo"
  ipsec_enabled       = true
  ssl_enabled         = true
  ssl_max_connections = 5
  project_name        = "default"
  vpc_id              = "vpc-3nr6adxxxxu8931eb64y4z2"
  subnet_id           = "subnet-btepcsxxxxw5h0b2u6hppyd"
}