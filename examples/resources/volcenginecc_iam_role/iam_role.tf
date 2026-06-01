resource "volcenginecc_iam_role" "RoleDemo" {
  role_name            = "RoleDemo"
  description          = "role attach policies"
  max_session_duration = 3600
  trust_policy_document = jsonencode(
    {
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "sts:AssumeRole"
          ],
          "Principal" : {
            "Service" : [
              "auto_scaling"
            ]
          }
        }
      ]
    }
  )
  policies = [
    { policy_name = "VPCFullAccess", policy_type = "System" },
    { policy_name = "SGFullAccess", policy_type = "System" },
    { policy_name = "RTCReadOnlyAccess", policy_type = "System" },
    { policy_name = "IAMReadOnlyAccess", policy_type = "System" },
    { policy_name = "EIPFullAccess", policy_type = "System" },
    { policy_name = "AdministratorAccess", policy_type = "System" },
    { policy_name = "SSLAccessCLBRolePolicy", policy_type = "System" },
    { "policy_name" = "ServiceRoleForNatGateway", "policy_type" = "System" },
    { "policy_name" = "ccapi-policy", "policy_type" = "Custom" }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}