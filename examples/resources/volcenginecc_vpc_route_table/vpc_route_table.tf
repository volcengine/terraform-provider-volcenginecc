resource "volcenginecc_vpc_route_table" "routeTableDemo" {
  vpc_id               = "vpc-*****"
  route_table_name     = "routeTableDemo"
  description          = "routeTableDemo description"
  project_name         = "default"
  subnet_ids           = ["subnet-rrx*****"]
  custom_route_entries = [
    {
      destination_cidr_block = "192.168.*.0/30"
      next_hop_type          = "Instance"
      next_hop_id            = "i-****"
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}