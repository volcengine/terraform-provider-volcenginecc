resource "volcenginecc_cloudidentity_permission_set" "CloudIdentityPermissionSetDemo" {
  description      = "cc-test"
  name             = "cctest-test"
  relay_state      = "https://console.xxxxx.com/xxxxx"
  session_duration = 3600
  permission_policies = [
    {
      permission_policy_name     = "AdministratorAccess"
      permission_policy_type     = "System"
      permission_policy_document = ""
    },
    {
      permission_policy_name     = "IAMFullAccess"
      permission_policy_type     = "System"
      permission_policy_document = ""
    },
    {
      permission_policy_name     = "InlinePolicy"
      permission_policy_type     = "Inline"
      permission_policy_document = "{\n    \"Statement\": [\n        {\n            \"Effect\": \"Allow\",\n            \"Action\": [\n                \"vpc:*\"\n            ],\n            \"Resource\": [\n                \"*\"\n            ]\n        }\n    ]\n}"
    }
  ]
}