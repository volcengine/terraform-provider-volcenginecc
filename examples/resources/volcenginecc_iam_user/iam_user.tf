resource "volcenginecc_iam_user" "userDemo" {
  user_name    = "userDemo"
  display_name = "userDemo"
  description  = "userDemo Description"
  mobile_phone = "1234342****"
  policies     = [
    {
      policy_name = "ReadOnlyAccess"
      policy_type = "System"
    }
  ]
  tags = [
    {
      key   = "enc"
      value = "test"
    }
  ]
  login_profile = {
    safe_auth_type            = "phone"
    safe_auth_exempt_required = 1
    safe_auth_exempt_unit     = 0
    safe_auth_exempt_duration = 10
    login_allowed             = true
    password                  = "********"
    password_reset_required   = false
    safe_auth_flag            = true
  }
  security_config = {
    safe_auth_type = "phone"
  }
}