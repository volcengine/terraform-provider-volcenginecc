resource "volcenginecc_apig_upstream" "APIGUpstreamAIProviderDemo" {
  circuit_breaking_settings = {
    base_ejection_time   = 30
    consecutive_errors   = 5
    enable               = true
    interval             = 10
    max_ejection_percent = 20
    min_health_percent   = 60
  }
  comments = "this is a test"
  connection_pool_settings = {
    enable                      = true
    http_1_max_pending_requests = 1000
    idle_timeout                = 1000
    max_connections             = 1000
  }
  gateway_id = "gd6hagxxxxxkh7is70"
  load_balancer_settings = {
    lb_policy       = "SimpleLB"
    simple_lb       = "LEAST_CONN"
    warmup_duration = 5
  }
  name     = "test-7"
  protocol = "HTTP"
  tls_settings = {
    sni      = ""
    tls_mode = "DISABLE"
  }
  source_type = "Domain"
  upstream_spec = {
    domain = {
      domain_list = [
        {
          domain = "www.test7.com"
        port = 5566 }
      ]
    }
  }
}