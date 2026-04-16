resource "volcenginecc_privatelink_vpc_endpoint" "PrivateLinkVpcEndpointDemo" {
  description         = "this is a test"
  endpoint_name       = "ccapi-test-1001"
  ip_address_versions = ["ipv4", "ipv6"]
  private_dns_enabled = true
  project_name        = "default"
  security_group_ids  = ["sg-btg9xxxxxc5h0b2u913txw"]
  service_id          = "epsvc-13fpxxxxx03n6nu5omuqde"
  service_name        = "com.volces.xxxx.cn-beijing.api-ipv6.ticket"
  tags = [{
    key   = "env"
    value = "test"
  }]
  vpc_id = "vpc-btg9hmxxxxx5h0b2tnp1on8"
  zones = [{
    private_ip_address    = "192.168.xxx.xx"
    private_ipv_6_address = "2406:d440:103:xxxxx:xxxxx:7eef:35c4:8d79"
    subnet_id             = "subnet-ijif1pxxxxx4o8culvzg85"
    zone_id               = "cn-beijing-a"
  }]

}