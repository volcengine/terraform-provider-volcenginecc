resource "volcenginecc_fwcenter_control_policy" "Example" {
  direction         = "in"
  action            = "accept"
  source            = "0.0.0.0/0"
  source_type       = "net"
  destination       = "0.0.0.0/0"
  destination_type  = "net"
  proto             = "TCP"
  dest_port         = "60000"
  dest_port_type    = "port"
  description       = "this is a test"
  ip_type           = "v4"
  prio              = 1
  status            = true
  repeat_type       = "Weekly"
  repeat_days       = [2, 3]
  repeat_start_time = "02:00"
  repeat_end_time   = "04:00"
  start_time        = 1782230400
  end_time          = 1782489540
}