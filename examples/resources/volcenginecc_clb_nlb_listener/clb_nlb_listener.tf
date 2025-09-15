resource "volcenginecc_clb_nlb_listener" "CLBNLBListenerDemo" {
  load_balancer_id   = "nlb-11zz9w3jqptz449iegfwvxxxx"
  protocol           = "TCP"
  port               = 0
  server_group_id    = "rsp-11zz9wdewa3uo49ieggq8xxxx"
  connection_timeout = 60
  description        = "CLBNLBListenerDemo description"
  enabled            = true
  listener_name      = "CLBNLBListenerDemo"
  tags = [
    {
      key = "env"
    value = "Test" }
  ]
}