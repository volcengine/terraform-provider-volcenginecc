resource "volcenginecc_organization_service_control_policy" "OrganizationServiceControlPolicyDemo" {
  policy_name = "ccapi-test"
  description = "ccapi-test-desc"
  statement = jsonencode(
    {
      "Statement" : [
        {
          "Effect" : "Deny",
          "Action" : [
            "iam:*"
          ],
          "Resource" : [
            "*"
          ],
          "Condition" : {
            "StringNotEqualsIfExists" : {
              "volc:UserName" : [
                "a"
              ]
            }
          }
        }
      ]
    }
  )
  targets = [
    {
      target_id   = "212***"
      target_type = "Account"
    },
    {
      target_id   = "21****"
      target_type = "Account"
    }
  ]
}