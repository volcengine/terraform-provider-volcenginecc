resource "volcenginecc_fwcenter_vpc_fire_wall_acl_rule" "Example" {
  action                 = "accept"
  description            = "this is a test"
  dest_port              = "60000"
  dest_port_type         = "port"
  destination            = "0.0.0.0/0"
  destination_type       = "net"
  domain_resolution_mode = ""
  end_time               = 1782662340
  ip_type                = "v4"
  prio                   = 1
  proto                  = "TCP"
  repeat_days            = [2, 3]
  repeat_end_time        = "04:00"
  repeat_start_time      = "02:00"
  repeat_type            = "Weekly"
  source                 = "0.0.0.0/0"
  source_type            = "net"
  start_time             = 1782403200
  status                 = true
  vpc_firewall_id        = "vfw-yeoxxxxxx4vvsy6xj"
}