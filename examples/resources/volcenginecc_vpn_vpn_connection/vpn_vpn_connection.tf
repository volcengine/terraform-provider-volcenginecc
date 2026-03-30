resource "volcenginecc_vpn_vpn_connection" "VpnVpnConnectionDemo" {
  vpn_connection_name = "terraform-test"
  description         = "IPsec-test"
  vpn_gateway_id      = "vgw-3nqn2s36hu1a89xxxxxxx"
  project_name        = "default"
  log_enabled         = true
  local_subnet        = ["0.0.0.0/2"]
  remote_subnet       = ["0.0.0.0/2"]
  customer_gateway_id = "cgw-vzzoltnuu41s865ykxxxxxx"
  negotiate_instantly = false
  ike_config = {
    psk       = "88888888"
    version   = "ikev1"
    mode      = "aggressive"
    auth_alg  = "sha1"
    enc_alg   = "aes"
    dh_group  = "group2"
    lifetime  = 86400
    local_id  = "0.0.0.0"
    remote_id = "0.0.0.0"
  }
  ipsec_config = {
    auth_alg = "sha1"
    dh_group = "group2"
    enc_alg  = "aes"
    lifetime = 86400
  }
  bgp_info = {
    enable_bgp   = true
    tunnel_cidr  = "169.xxx.30.40/30"
    local_bgp_ip = "169.xxx.30.41"
  }
  attach_type   = "VpnGateway"
  nat_traversal = true
  dpd_action    = "restart"
}
