resource "volcenginecc_alb_load_balancer" "ALBLoadBalancerDemo" {
  type                       = "private"
  address_ip_version         = "IPv4"
  load_balancer_name         = "ALBLoadBalancerDemo"
  description                = "ALBLoadBalancerDemo description"
  vpc_id                     = "vpc-rrco37ovjq4gv0x58zfxxxx"
  load_balancer_billing_type = 1
  bandwidth_package_id       = "bwp-1c099l94j13b45e8j6zz5xxxx"
  delete_protection          = "off"
  zone_mappings = [
    {
      subnet_id = "subnet-rrwqhg3qzxfkv0x57g3xxxx"
    zone_id = "cn-test-a" }
  ]
  project_name                   = "default"
  modification_protection_status = "NonProtection"
  modification_protection_reason = ""
  load_balancer_edition          = "Standard"
  waf_protection_enabled         = "off"
  tags = [
    {
      key = "test"
    value = "env" }
  ]
}