resource "volcenginecc_iam_user" "UserDemo" {
  user_name   = "UserDemo"
  description = "user for tf demo"
  groups      = [
    "UserGroupDemo"
  ]
  policies = [
    {
      policy_name = "TOSReadOnlyAccess"
      policy_type = "System"
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}