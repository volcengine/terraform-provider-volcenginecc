resource "volcenginecc_apig_gateway_service" "ApigGatewayServiceDemo" {
  service_name = "ApigGatewayServiceDemo"
  gateway_id   = "gd3vehjs7npja181xxxxx"
  protocol     = ["HTTP", "HTTPS"]
  auth_spec = {
    enable = false
  }
  comments = "ApigGatewayServiceDemo-test"
}