resource "volcenginecc_id_service" "IDServiceDemo" {
  permission_space_id = "fd46c764-e94e-4c94-ba69-6bfxxxxxxxx"
  backend = {
    timeout_seconds = 60
    enable_tls      = false
    backend_port    = 80
    backend_domain  = "api.internal.com"
    protocol        = "HTTP"
  }
  description  = "测试枚举分支覆盖"
  service_name = "test-other-service"
  routes = [{
    path     = "/api/v1/orders"
    priority = 20
    api_spec = {
      extractors = [{
        path          = "/orderId"
        is_auth       = false
        resource_type = "Order"
        source        = 2
      }]
      action_type  = "write"
      action_value = "query"
      response_filters = [{
        mode          = 2
        is_auth       = false
        resource_type = "Order"
      }]
      identifier_value = "/order/id"
      identifier_type  = 2
    }
    route_name         = "order-query"
    resource_type      = "Collection"
    method             = "GET"
    auth_resource_type = "Order"
    path_match_type    = "Prefix"
  }]
  tags = [{
    value = "env"
    key   = "test"
  }]
}