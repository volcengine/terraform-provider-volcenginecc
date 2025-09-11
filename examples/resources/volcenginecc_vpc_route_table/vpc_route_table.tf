resource "volcenginecc_vpc_route_table" "RouteTableDemo" {
  vpc_id               = "vpc-xxxxx"
  route_table_name     = "routeTableDemo"
  description          = "routeTableDemo description"
  project_name         = "default"
  subnet_ids           = ["subnet-rrxxxxxx"]
  custom_route_entries = [
    {
      destination_cidr_block = "192.168.x.0/30"
      next_hop_type          = "Instance"
      next_hop_id            = "i-xxxx"
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}