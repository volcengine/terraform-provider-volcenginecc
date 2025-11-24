resource "volcenginecc_natgateway_dnatentry" "NatGatewayDnatentryDemo" {
  dnat_entry_name = "NatGatewayDnatentryDemo"
  protocol        = "udp"
  internal_ip     = "192.168.xxx.53"
  internal_port   = "20-25"
  external_ip     = "115.190.xxx.9"
  external_port   = "25-30"
  nat_gateway_id  = "ngw-2pc28yhdpbx8g227qo1sxxxxx"
  port_type       = "specified"
}