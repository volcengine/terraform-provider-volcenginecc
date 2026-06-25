resource "volcenginecc_cr_vpc_endpoint" "CRVpcEndpointDemo" {
  registry = "registry_xxxxxx"
  vpcs = [
    {
      account_id = 21000000000000
      subnet_id  = "subnet_id_xxxx"
      vpc_id     = "vpc_id_xxxx"
    }
  ]
}