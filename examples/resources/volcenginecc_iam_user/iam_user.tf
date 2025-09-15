resource "volcenginecc_iam_user" "UserDemo" {
  user_name   = "UserDemo"
  description = "user description"
  groups = [
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