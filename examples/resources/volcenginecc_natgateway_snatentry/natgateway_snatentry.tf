resource "volcenginecc_natgateway_snatentry" "NatGatewaySnatentryDemo" {
  nat_gateway_id  = "ngw-2pc28yhdpbx8g227qo1xxxxx"
  eip_id          = "eip-iivdtssgbdog74o8cuxxxxx,eip-btbv1pk36g3k5h0b2vxxxxx"
  snat_entry_name = "私有网络"
  source_cidr     = "0.0.0.0/0"
}