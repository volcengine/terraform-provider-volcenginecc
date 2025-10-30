resource "volcenginecc_cr_registry" "CRRegistryDemo" {
  project = "default"
  name    = "test"
  type    = "Enterprise"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}