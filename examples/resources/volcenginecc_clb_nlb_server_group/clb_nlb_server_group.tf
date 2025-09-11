resource "volcenginecc_clb_nlb_server_group" "NLBServerGroupDemo" {
  bypass_security_group_enabled = false
  health_check = {
    enabled = false
  }
  vpc_id = "vpc-13f8k4dwdsydc3n6nu5rxxxxx"
  project_name = "default"
  server_group_name = "NLBServerGroupDemo.Aa1xxxxx"
  type = "instance"
  protocol = "UDP"
  description = "test.IPV4.UDP.WLC_="
  scheduler = "wlc"
  ip_address_version = "ipv4"
  any_port_enabled = false
  connection_drain_enabled = false
  preserve_client_ip_enabled = false
  session_persistence_enabled = false
  proxy_protocol_type = "off"
  servers = [
    {
      instance_id = "i-ye2fvd0qo0bw80ctxxxxx"
      type = "ecs"
      ip = "192.168.xx.76"
      port = 10
      weight = 50    },
    {
      instance_id = "eni-mirt64nt1xq85smt1a3xxxxx"
      type = "eni"
      ip = "192.168.xx.77"
      port = 20
      weight = 50    }
  ]
  tags = [
    {
      key = "env"
      value = "test"    }
  ]
}