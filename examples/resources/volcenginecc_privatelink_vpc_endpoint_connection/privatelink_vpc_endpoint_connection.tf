resource "volcenginecc_privatelink_vpc_endpoint_connection" "PrivateLinkVpcEndpointConnectionDemo" {
  endpoint_id = "ep-2rxb5wrxxxxxxxukcknq"
  service_id  = "epsvc-1mg2xrmsxxxxxxxxconsso"
  resources_allocate = [{
    resource_id = "clb-13g8lgxxxxxxx6nu5ajtpp2"
    zone_id     = "cn-beijing-c"
    }, {
    resource_id = "clb-13fx74xxxxxxxnu4g85ggo"
    zone_id     = "cn-beijing-b"
  }]
}