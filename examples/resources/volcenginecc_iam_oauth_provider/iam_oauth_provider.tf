resource "volcenginecc_iam_oauth_provider" "IAMOauthProviderDemo" {
  oauth_provider_name = "tf_oauth"
  sso_type            = 2
  status              = 2
  description         = "terraform test"
  client_id           = "4003ba40864e4215b929142dxxxxx"
  client_secret       = "ab44f642d6fb4a559fd5e5caedxxxxx"
  user_info_url       = "https://security-api.xxxxx.com/qa/sso/oauth/xxxxx"
  token_url           = "https://security-api.xxxxx.com/qa/sso/oauth/xxxxx"
  authorize_url       = "https://security-api.xxxxxx.com/qa/sso/oauth/xxxxxx"
  authorize_template  = "$${authEndpoint}?client_id=$${clientId}&scope=$${scope}&response_type=code&state=12345"
  scope               = "ccapi terraform"
  identity_map_type   = 1
  idp_identity_key    = "username"
}