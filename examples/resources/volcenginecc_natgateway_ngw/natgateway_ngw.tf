resource "volcenginecc_natgateway_ngw" "NatGatewayNGWDemo" {
  spec             = "Small"
  vpc_id           = "vpc-2f8kicbjkot8g4f4pzyyxxxxx"
  subnet_id        = "subnet-3hicda1321wqo3nkipk4xxxxx"
  nat_gateway_name = "NatGatewayNGWDemo"
  billing_type     = 2
  project_name     = "default"
  network_type     = "internet"
  description      = "NatGatewayNGWDemo"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}