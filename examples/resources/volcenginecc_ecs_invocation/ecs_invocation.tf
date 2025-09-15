resource "volcenginecc_ecs_invocation" "InvocationDemo" {
  invocation_name        = "InvocationDemo"
  invocation_description = "InvocationDemo desc"
  working_dir            = "/home"
  username               = "InvocationDemo"
  windows_password       = "********"
  timeout                = 60
  instance_ids           = ["i-ye2v6l0pvkqc6inxxxxx"]
  repeat_mode            = "Rate"
  frequency              = "1h"
  launch_time            = "2025-08-30T11:10Z"
  recurrence_end_time    = "2025-08-31T11:04Z"
  project_name           = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  command_id = "cmd-ye28kugp249tzrexxxxx"
  parameters = "{\"dirname\":\"10\"}"
}