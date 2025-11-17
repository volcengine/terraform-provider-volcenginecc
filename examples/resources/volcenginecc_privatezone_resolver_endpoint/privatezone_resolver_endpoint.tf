resource "volcenginecc_privatezone_resolver_endpoint" "PrivateZoneResolverEndpointDemo" {
  name          = "PrivateZoneResolverEndpointDemo"
  vpc_id        = "vpc-3nrh1tqschwcg931eaqxxxxx"
  endpoint_type = "IPv4"
  ip_configs = [
    {
      az_id     = "cn-beijing-a"
      subnet_id = "subnet-bt50na0bf6kg5h0b2u1xxxxx"
      ip        = "192.168.xx.44"
    },
    {
      az_id     = "cn-beijing-b"
      subnet_id = "subnet-1a14u8n59jdvk8nvepjyxxxxx"
      ip        = "192.168.xx.154"
    }
  ]
  project_name = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  vpc_region = "cn-beijing"
  direction  = "OUTBOUND"
}