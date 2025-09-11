resource "volcenginecc_clb_nlb" "NLBDemo" {
  ipv_4_network_type = "intranet"
  security_group_ids = ["sg-13fg4nslwnwu83n6nu5a7xxxx"]
  vpc_id = "vpc-13ffun9olqz9c3n6nu43bxxxx"
  load_balancer_name = "NLBDemo"
  description = "NLBDemom"
  project_name = "default"
  ip_address_version = "ipv4"
  cross_zone_enabled = false
  zone_mappings = [
    {
      subnet_id = "subnet-mjoyynjt59ts5smt1bncxxxx"
      zone_id = "cn-beijing-a"
      ipv_4_address = "192.168.xx.2"    }
  ]
  modification_protection_status = "ConsoleProtection"
  tags = [
    {
      key = "env"
      value = "test"    }
  ]
}