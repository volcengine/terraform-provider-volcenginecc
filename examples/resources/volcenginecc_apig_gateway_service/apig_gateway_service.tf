resource "volcenginecc_apig_gateway_service" "ApigGatewayServiceDemo" {
  service_name = "ccapi-terraform-1"
  gateway_id   = "gd6l9lbilgrmdxxxxxx"
  protocol     = ["HTTP", "HTTPS"]
  auth_spec = {
    enable = true
  }
  comments     = "test"
  service_type = "AIProvider"
  service_network_spec = {
    enable_public_network  = true
    enable_private_network = false
    private_network_ip     = []
  }
}