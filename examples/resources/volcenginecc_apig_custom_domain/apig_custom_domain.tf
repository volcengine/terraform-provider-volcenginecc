resource "volcenginecc_apig_custom_domain" "ApigCustomDomainDemo" {
  service_id     = "sd50d3h5xxxm0t5xxxxx"
  domain         = "www.****.com"
  protocol       = ["HTTP", "HTTPS"]
  certificate_id = "cert-775906d873xxx5bc9d1d372b5dxxxxx"
  comments       = "ApigCustomDomainDemo custom domain"
  ssl_redirect   = true
}