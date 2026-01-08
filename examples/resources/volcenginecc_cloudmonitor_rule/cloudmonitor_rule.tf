resource "volcenginecc_cloudmonitor_rule" "CloudMonitorRuleDemo" {
  rule_name     = "CloudMonitorRuleDemo"
  description   = "this is a description"
  rule_type     = "static"
  namespace     = "VCM_AgentKitMcp"
  sub_namespace = "MCP_Service"
  original_dimensions = {
    key    = "ResourceID"
    values = ["*"]
  }
  level               = "warning"
  evaluation_count    = 1
  enable_state        = "enable"
  regions             = ["cn-beijing"]
  multiple_conditions = false
  no_data = {
    enable           = false
    evaluation_count = 4
  }
  level_conditions = [
    {
      level = "warning"
      conditions = [
        {
          metric_name         = "istio_requests_total"
          statistics          = "sum"
          comparison_operator = ">"
          threshold           = "100000"
        metric_unit = "Count" }
    ] }
  ]
  conditions = [
    {
      metric_name         = "istio_requests_total"
      statistics          = "sum"
      comparison_operator = ">"
      threshold           = "100000"
    metric_unit = "Count" }
  ]
  condition_operator = "&&"
  silence_time       = 5
  recovery_notify = {
    enable = true
  }
  project_name = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  alert_methods     = ["Email", "Webhook"]
  webhook_ids       = ["2005838xxxxx67xxxxx"]
  contact_group_ids = ["20041xxxxxx06xxxxx"]
  effect_start_at   = "00:00"
  effect_end_at     = "23:59"
}