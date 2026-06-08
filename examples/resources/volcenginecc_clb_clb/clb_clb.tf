resource "volcenginecc_clb_clb" "CLBDemo" {
  load_balancer_name             = "CLBDemo"
  load_balancer_spec             = "small_1"
  address_ip_version             = "ipv4"
  bypass_security_group_enabled  = "off"
  description                    = "CLBDemo description"
  load_balancer_billing_type     = 2
  master_zone_id                 = "cn-beijing-a"
  slave_zone_id                  = "cn-beijing-b"
  modification_protection_reason = "xx"
  modification_protection_status = "xx"
  period                         = 1
  period_unit                    = "Month"
  project_name                   = "default"
  region_id                      = "cn-beijing"
  subnet_id                      = "subnet-rrwqhg3qzxfkv0x57g3xxxx"
  type                           = "public"
  vpc_id                         = "vpc-rrco37ovjq4gv0x58zfxxxx"
  eip = {
    isp              = "BGP"
    bandwidth        = 1
    eip_billing_type = 3
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  access_log = {
    bucket_name    = "ccapi-test"
    tls_project_id = "5554e74f-0413-4375-ad99-e544xxxxx"
    tls_topic_id   = "7f3bc374-5e1d-4984-83fc-0e5a5xxxxx"
  }
}