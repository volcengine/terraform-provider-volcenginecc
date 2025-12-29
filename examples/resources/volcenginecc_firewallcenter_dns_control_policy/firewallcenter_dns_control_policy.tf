resource "volcenginecc_firewallcenter_dns_control_policy" "FirewallCenterDnsControlPolicyDemo" {
  description          = "FirewallCenterDnsControlPolicyDemo test"
  destination          = "www.xxxx.com"
  destination_type     = "domain"
  internet_firewall_id = "ifw-yebxxxxx0iac7gxxxxx"
  sources = [
    {
      region = "cn-beijing"
    vpc_id = "vpc-3rehw4xxxxk2ixxxxx" }
  ]
}