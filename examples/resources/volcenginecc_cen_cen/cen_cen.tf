resource "volcenginecc_cen_cen" "CENCENDemo" {
  cen_name     = "CENCENDemo"
  description  = "CENCENDemo descripiton"
  project_name = "iac"
  tags = [
    {
      key = "dev"
    value = "test" }
  ]
}