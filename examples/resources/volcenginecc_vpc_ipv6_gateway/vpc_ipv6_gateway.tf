resource "volcenginecc_vpc_ipv6_gateway" "Ipv6GatewayDemo" {
  vpc_id       = "vpc-3nrh1tqschwcg931exxxxx"
  name         = "terraform-test-2"
  description  = "Create by terraform"
  project_name = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}