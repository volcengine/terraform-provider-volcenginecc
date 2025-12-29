resource "volcenginecc_dns_zone" "DnsZoneDemo" {
  zone_name    = "DnsZoneDemo"
  remark       = "test"
  project_name = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}