resource "volcenginecc_cloudmonitor_contact_group" "CloudMonitorContactGroupDemo" {
  name        = "TestContactGroup"
  description = "TestContactGroup Description"
  contacts = [
    {
      contact_id = "20562729774143xxxx"
    },
    {
      contact_id = "20562729184541xxxx"
    },
    {
      contact_id = "20562728672541xxxx"
    }
  ]
}