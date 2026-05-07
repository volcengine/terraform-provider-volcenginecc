resource "volcenginecc_cr_endpoint_acl_policy" "CREndpointAclPolicyDemo" {
  registry    = "xxxxx"
  type        = "Public"
  entry       = "0.0.0.0/0"
  description = "test acl policy"
}