resource "volcenginecc_iam_saml_provider" "IAMSamlProviderDemo" {
  saml_provider_name             = "ccapi-test"
  description                    = "ccapi-test"
  encoded_saml_metadata_document = "PD94bWwgdmVyc2lvbxxxxxx=="
  sso_type                       = 1
  status                         = 1
}