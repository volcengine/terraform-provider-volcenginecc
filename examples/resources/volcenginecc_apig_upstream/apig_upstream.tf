resource "volcenginecc_apig_upstream" "APIGUpstreamAIProviderDemo" {
  name        = "APIGUpstreamAIProviderDemo"
  gateway_id  = "gd45elb819ma6giexxxxx"
  comments    = "APIGUpstreamAIProviderDemo"
  protocol    = "HTTP"
  source_type = "K8S"
  upstream_spec = {
    k8_s_service = {
      namespace = "namespace"
      name      = "server"
      port      = 2001
    }
  }
  load_balancer_settings = {
    lb_policy = "SimpleLB"
    simple_lb = "RANDOM"
  }
  circuit_breaking_settings = {
    enable               = true
    consecutive_errors   = 5
    interval             = 10000
    base_ejection_time   = 30000
    max_ejection_percent = 20
    min_health_percent   = 60
  }
  tls_settings = {
    tls_mode = "DISABLE"
  }
}