resource "volcenginecc_cloudidentity_group" "CloudIdentityGroupDemo" {
  description  = "ccapi-multi-des"
  display_name = "test for ccapi"
  group_name   = "ccapi-multi"
  join_type    = "Manual"
  members = [
    {
      user_id = "***********"
    },
    {
      user_id = "*************"
    }
  ]
}