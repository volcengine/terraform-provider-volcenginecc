resource "volcenginecc_natgateway_nat_ip" "NatGatewayNatIpDemo" {
  nat_gateway_id     = "ngw-2hgk22kpfp5a874wjohxxxxx"
  nat_ip_description = "NatGatewayNatIpDemo description"
  nat_ip_name        = "NatGatewayNatIpDemo"
  nat_ip             = "192.168.xxx.xx"
}