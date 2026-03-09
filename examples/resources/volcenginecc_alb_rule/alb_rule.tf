resource "volcenginecc_alb_rule" "AlbRuleDemo" {
  description = "test-rewrite"
  domain      = "www.xxxx.test.com"
  listener_id = "lsn-bdwhrmgvcyyo8dv4xxxxxx"
  rewrite_config = {
    rewrite_path = "/test"
  }
  rewrite_enabled       = "on"
  rule_action           = ""
  traffic_limit_enabled = "off"
  traffic_limit_qps     = 0
  url                   = "/"
  forward_group_config = {
    server_group_tuples = [{
      server_group_id = "rsp-1pf4pgyq8zitc845wfxxxxxx",
      weight          = 100
      },
      {
        server_group_id = "rsp-1pf4pgyq8zitc845wfxxxxxx",
        weight          = 100
    }]
    sticky_session_enabled = "on"
    sticky_session_timeout = 2000
  }
}