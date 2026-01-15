resource "volcenginecc_iam_group" "GroupDemo" {
  user_group_name = "GroupDemo"
  description     = "GroupDemo-Description"
  display_name    = "GroupDemo-DisplayName"
  users = [
    {
    user_name = "demo" }
  ]
  attached_policies = [
    {
      policy_type = "System"
      policy_name = "ECSFullAccess"
      policy_scope = [
        {
          attach_time          = "20230810T071***Z"
          policy_scope_type    = "Project"
          project_display_name = "demo"
        project_name = "Project" }
    ] }
  ]
}