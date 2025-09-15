resource "volcenginecc_alb_certificate" "ALBCertificateDemo" {
  certificate_name = "ALBCertificateDemo"
  certificate_type = "CA"
  public_key       = "-----BEGIN CERTIFICATE-----xxxx-----END CERTIFICATE-----"
  description      = "ALBCertificateDemo description"
  project_name     = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}