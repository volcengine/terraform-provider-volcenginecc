resource "volcenginecc_cen_route_entry" "CENRouteEntryDemo" {
  instance_type          = "VPC"
  instance_region_id     = "cn-beijing"
  instance_id            = "vpc-iirjlwem73eo74o8ctxxxxx"
  destination_cidr_block = "192.x.0.0/24"
  cen_id                 = "cen-2v73nw1h8a03k6x7e8xxxxx"
}