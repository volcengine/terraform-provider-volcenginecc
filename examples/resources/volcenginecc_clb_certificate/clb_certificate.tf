resource "volcenginecc_clb_certificate" "CertificatDemo" {
  certificate_name = "CertificatDemo"
  private_key      = "-----BEGIN RSA PRIVATE KEY-----***----END RSA PRIVATE KEY-----"
  public_key       = "-----BEGIN CERTIFICATE-----***-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----***-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----***-----END CERTIFICATE-----"
  description      = "CertificateDemo Example"
  project_name     = "default"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}