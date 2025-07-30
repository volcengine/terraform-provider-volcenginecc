resource "volcenginecc_iam_role" "RoleDemo" {
  role_name             = "RoleDemo"
  description           = "role attach policies"
  max_session_duration  = 3600
  trust_policy_document = "{\n    \"Statement\": [\n        {\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"sts:AssumeRole\"\n            ],\n            \"Principal\": {\n                \"IAM\": [\n                    \"trn:iam::20000000xx:root\"\n                ]\n            }\n        }\n    ]\n}"
  tags                  = [
    {
      key   = "env"
      value = "test"
    }
  ]
}