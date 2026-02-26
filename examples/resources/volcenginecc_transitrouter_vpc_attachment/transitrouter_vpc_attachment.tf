resource "volcenginecc_transitrouter_vpc_attachment" "TransitRouterVpcAttachmentDemo" {
  appliance_mode_enabled = false
  attach_points = [
    {
      subnet_id = "subnet-ijifxxxxx8culvzg85"
    zone_id = "cn-beijing-a" }
  ]
  auto_publish_route_enabled = false
  description                = "test"
  ipv_6_enabled              = false
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  transit_router_attachment_name = "ccapi-test-1"
  transit_router_id              = "tr-mjl8xxxxxg5smt1boobol4"
  vpc_id                         = "vpc-btg9hmxxxxx0b2tnp1on8"
}