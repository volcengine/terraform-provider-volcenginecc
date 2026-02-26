resource "volcenginecc_vpc_ha_vip" "VPCHAVIPDemo" {
  associated_eip_id        = "eip-1a1tzpbhmsphc8nvexxxxx"
  associated_instance_ids  = ["eni-btw4pmtcabr45h0b2xxxxx", "eni-iiotcbjp7r4074o8xxxxx", "eni-3nrjyim6prm68931ebbxxxxx"]
  associated_instance_type = "NetworkInterface"
  description              = "this is a test"
  ha_vip_name              = "HaVip-test"
  ip_address               = "192.***.0.9"
  subnet_id                = "subnet-rrwqhg3qzxfkv0x57xxxxx"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}