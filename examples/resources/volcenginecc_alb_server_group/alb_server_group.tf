resource "volcenginecc_alb_server_group" "AlbServergroupDemo" {
  vpc_id            = "vpc-13f8***"
  server_group_name = "test-servergroup"
  server_group_type = "instance"
  protocol          = "HTTP"
  scheduler         = "wrr"
  health_check = {
    enabled             = "on"
    port                = 80
    interval            = 2
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
    method              = "GET"
    http_version        = "HTTP1.1"
    uri                 = "/health"
    http_code           = "http_2xx"
  }
  sticky_session_config = {
    sticky_session_enabled = "on"
    sticky_session_type    = "insert"
    cookie_timeout         = 1000
  }
  servers = [
    {
      instance_id = "i-ye3***"
      type        = "ecs"
      ip          = "192.168.1.**"
      port        = 80
      weight      = 10
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}