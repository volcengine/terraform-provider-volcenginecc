resource "volcenginecc_organization_account" "OrganizationAccountDemo" {
  account_name  = "testdemo"
  description   = "test-desc"
  org_unit_id   = "7538034613*****190"
  show_name     = "test-show-name"
  allow_console = 1
  secure_contact_info = {
    new_email = ""
    new_phone = "*********"
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}