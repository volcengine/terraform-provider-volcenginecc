resource "volcenginecc_alb_health_check_template" "ALBHealthCheckTemplateDemo" {
  description                = "asdfgh"
  health_check_domain        = "test.com"
  health_check_http_code     = "http_2xx,http_3xx"
  health_check_http_version  = "HTTP1.0"
  health_check_interval      = 2
  health_check_method        = "HEAD"
  health_check_port          = 0
  health_check_protocol      = "HTTP"
  health_check_template_name = "test-test-1001"
  health_check_timeout       = 2
  health_check_uri           = "/"
  healthy_threshold          = 3
  unhealthy_threshold        = 3
  project_name               = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]

}