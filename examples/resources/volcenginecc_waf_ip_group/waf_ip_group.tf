resource "volcenginecc_waf_ip_group" "example" {
  name    = "ccapi-ipgroup-full"
  ip_list = ["2.2.2.2", "10.0.0.0/8"]
}
