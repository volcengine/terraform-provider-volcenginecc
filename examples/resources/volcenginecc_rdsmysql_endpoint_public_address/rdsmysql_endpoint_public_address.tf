resource "volcenginecc_rdsmysql_endpoint_public_address" "EndpointPublicAddressDemo" {
  instance_id   = "mysql-7623a9xxxxxx"
  endpoint_id   = "mysql-7623a9xxxxx-custom-xxxx"
  eip_id        = "eip-2f8vbiyym6txckjzxxxxx"
  domain_prefix = "ccapi-terraform"
  port          = 23306
}