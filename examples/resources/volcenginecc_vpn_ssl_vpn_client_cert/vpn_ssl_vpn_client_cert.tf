resource "volcenginecc_vpn_ssl_vpn_client_cert" "VPNSslVpnClientCertDemo" {
  ssl_vpn_server_id        = "vss-xxxx"
  ssl_vpn_client_cert_name = "ccapi-test-client-cert"
  description              = "ccapi-test"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}