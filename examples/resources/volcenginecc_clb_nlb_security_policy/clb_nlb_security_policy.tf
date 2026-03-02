resource "volcenginecc_clb_nlb_security_policy" "CLBNLBSecurityPolicyDemo" {
  ciphers              = ["TLS_AES_128_CCM_SHA256", "TLS_AES_128_GCM_SHA256"]
  project_name         = "default"
  security_policy_name = "ccapi-test-1"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  tls_versions = ["TLSv1.3"]
}