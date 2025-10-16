resource "volcenginecc_clb_listener" "ClbListenerDemo" {
  acl_ids                  = ["acl-mjd01crvppts5smt1bsxxxxx"]
  acl_status               = "on"
  acl_type                 = "black"
  connection_drain_enabled = "on"
  connection_drain_timeout = 900
  description              = "ClbListenerDemo description"
  enabled                  = "on"
  established_timeout      = 300
  server_group_id          = "rsp-rs11ie8u6neov0x58bxxxxx"
  health_check = {
    enabled             = "on"
    healthy_threshold   = 3
    interval            = 5
    timeout             = 2
    unhealthy_threshold = 3
    port                = 0
    udp_request         = "test"
    udp_expect          = "test"
  }
  listener_name       = "ClbListenerDemo"
  load_balancer_id    = "clb-rs1187938g00v0x58nxxxxx"
  persistence_timeout = 1000
  persistence_type    = "source_ip"
  protocol            = "UDP"
  port                = 5001
  scheduler           = "wrr"
  bandwidth           = 1
  proxy_protocol_type = "standard"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}