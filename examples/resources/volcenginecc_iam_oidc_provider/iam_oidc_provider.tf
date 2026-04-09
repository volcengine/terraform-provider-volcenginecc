resource "volcenginecc_iam_oidc_provider" "IAMOidcProviderDemo" {
  oidc_provider_name  = "ccapi-test"
  thumbprints         = ["b676ffa3179e8812093a1b5eafee876ae7a6aaf231078dad1bfbxxxxxx"]
  client_ids          = ["sts.test1.com", "sts.test2.com"]
  description         = "this is a test"
  issuance_limit_time = 10
  issuer_url          = "https://oidc-vke-cn-xxx.tos-cn-boe.volces.com/test"
}