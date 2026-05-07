resource "volcenginecc_cr_registry" "CRRegistryDemo" {
  project = "default"
  name    = "test"
  type    = "Enterprise"
  endpoint = {
    enabled = true
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}