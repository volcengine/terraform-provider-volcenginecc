resource "volcenginecc_tls_project" "TLSProjectDemo" {
  project_name     = "ccapi_test"
  description      = "ccapi_test"
  iam_project_name = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}