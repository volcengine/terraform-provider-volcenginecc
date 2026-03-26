resource "volcenginecc_cen_cen" "CENCENDemo" {
  cen_name     = "CENCENDemo"
  description  = "CENCENDemo descripiton"
  project_name = "iac"
  tags = [
    {
      key = "dev"
    value = "test" }
  ]
  instances = [
    {
      instance_id        = "vpc-3psq16gaiw4qo6csxyv***",
      instance_owner_id  = "21074****",
      instance_region_id = "cn-beijing",
      instance_type      = "VPC"
    }
  ]
}