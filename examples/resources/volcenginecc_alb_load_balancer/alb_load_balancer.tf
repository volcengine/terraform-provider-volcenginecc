resource "volcenginecc_alb_load_balancer" "ALBLoadBalancerDemo" {
  type                       = "public"
  address_ip_version         = "DualStack"
  load_balancer_name         = "ccapi-alb-4"
  description                = "Create by ccaip"
  vpc_id                     = "vpc-rrco37ovjq4gv0x5xxxxx"
  load_balancer_billing_type = 1
  delete_protection          = "off"
  bandwidth_package_id       = "bwp-1a1i9jjnawidc8nvexxxxx"
  ipv_6_bandwidth_package_id = "bwp-1a1i9jjnawidc8nvxxxxx"
  eip_billing_config = {
    isp                             = "BGP"
    billing_type                    = 3
    bandwidth                       = 1
    security_protection_types       = "AntiDDoS_Enhanced"
    security_protection_instance_id = 743
  }
  zone_mappings = [
    {
      subnet_id = "subnet-rrwqhg3qzxfkv0xxxxxx"
    zone_id = "cn-beijing-a" },
    {
      subnet_id = "subnet-btnzu3hrc0005h0xxxxx"
    zone_id = "cn-beijing-b" }
  ]
  ipv_6_eip_billing_config = {
    isp          = "BGP"
    billing_type = 3
    bandwidth    = 1
  }
  project_name                   = "default"
  modification_protection_status = "NonProtection"
  load_balancer_edition          = "Standard"
  waf_protection_enabled         = "off"
  tags = [
    {
      key = "test"
    value = "env" }
  ]
}