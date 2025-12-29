resource "volcenginecc_cloudidentity_user" "CloudIdentityUserDemo" {
  user_name               = "CloudIdentityUserDemo"
  display_name            = "CloudIdentityUserDemo"
  description             = "CloudIdentityUserDemo description"
  email                   = "xxx@163.com"
  phone                   = "***********"
  password                = "********"
  password_reset_required = true
}