resource "volcenginecc_privatelink_endpoint_service" "PrivateLinkEndpointServiceDemo" {
  service_type          = "Interface"
  ip_address_versions   = ["ipv4"]
  service_resource_type = "CLB"
  resources = [
    {
      instance_id = ""
      resource_id = "clb-rr0o8ni4dxxxx58wxxxxx"
      zone_ids    = []
    }
  ]
  auto_accept_enabled = true
  private_dns_enabled = true
  private_dns_type    = "public"
  private_dns_name    = "*.www.example.com"
  description         = "PrivateLinkEndpointServiceDemo description"
  project_name        = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}