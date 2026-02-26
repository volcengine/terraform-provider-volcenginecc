resource "volcenginecc_vpn_customer_gateway" "vpncustomergatewayDemo" {
  asn                   = 64513
  customer_gateway_name = "ccapi-test"
  description           = "testdesc"
  ip_address            = "115.***.151.5"
  ip_version            = "ipv4"
  project_name          = "default"
}