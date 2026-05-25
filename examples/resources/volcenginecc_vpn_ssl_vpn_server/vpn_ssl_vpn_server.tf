resource "volcenginecc_vpn_ssl_vpn_server" "VPNSslVpnServerDemo" {
  compress            = false
  client_ip_pool      = "192.168.xxx.0/26"
  description         = "修改的SSL服务端描述"
  project_name        = "default"
  port                = 1195
  ssl_vpn_server_name = "更新的SSL服务端名称"
  local_subnets = [
    "192.168.1.0/24",
    "192.168.2.0/24"
  ]
  vpn_gateway_id             = "vgw-****"
  auth                       = "SHA1"
  cipher                     = "AES-128-CBC"
  client_cert_session_policy = "PreemptExisting"
  protocol                   = "TCP"
  tags = [
    {
      value = "env"
      key   = "test"
    }
  ]
}