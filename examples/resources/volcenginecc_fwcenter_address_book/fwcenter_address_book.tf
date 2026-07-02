resource "volcenginecc_fwcenter_address_book" "Example" {
  group_name        = "IPv6手动地址簿"
  group_type        = "ipv6"
  description       = "测试IPv6手动地址簿"
  cloud_firewall_id = "CloudFirewallxxxxxxxam7tr9t02w"
  auto_update_type  = "Manual"
  address_detail_list = [{
    description = "测试IPv6网段"
    address     = "2001:xxxx::/32"
    }, {
    description = "测试IPv6地址"
    address     = "2001:xxxx::1"
  }]
}