resource "volcenginecc_clb_rule" "RuleDemo" {
  listener_id = "lsn-13****"
  description = "RuleDemo Example"
  domain      = "www.***.com"
  action_type = "Forward"

  server_group_id = "rsp-mj***"
  url             = "/co3cee"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}