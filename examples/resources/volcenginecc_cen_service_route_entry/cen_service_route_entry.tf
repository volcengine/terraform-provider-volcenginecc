resource "volcenginecc_cen_service_route_entry" "CENServiceRouteEntryDemo" {
  cen_id                 = "cen-2v73nw1h8a03k6x7exxxxxx"
  destination_cidr_block = "100.x.0.0/11"
  service_vpc_id         = "vpc-iirjlwem73eo74o8cxxxxx"
  description            = "ccapi-test"
  publish_mode           = "Custom"
  publish_to_instances = [
    {
      instance_region_id = "cn-hongkong"
      instance_type      = "VPC"
      instance_id        = "vpc-2uzvkdpoof7cw17q7y2xxxxx"
    },
    {
      instance_region_id = "cn-shanghai"
      instance_type      = "VPC"
      instance_id        = "vpc-33glll2z1gem86k70bxxxxx"
    }
  ]
}