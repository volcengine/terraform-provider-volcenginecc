resource "volcenginecc_cloudmonitor_alert_template" "example" {
  name         = "ccapi-alert-template-full"
  description  = "ccapi cloudmonitor alert template full"
  notify_mode  = "Resource"
  silence_time = 10
  no_data = {
    enable           = true
    evaluation_count = 3
    level            = "notice"
  }
  recovery_notify = {
    enable = true
  }
  template_rules = [{
    condition_operator  = "&&"
    evaluation_count    = 3
    evaluation_interval = 3
    level_conditions = [{
      level = "warning"
      conditions = [{
        comparison_operator = ">"
        evaluation_window   = 1
        metric_name         = "istio_requests_total"
        metric_unit         = "Count"
        statistics          = "sum"
        threshold           = "80"
      }]
    }]
    multiple_conditions = true
    name                = "full-rule"
    namespace           = "VCM_AgentKitMcp"
    sub_namespace       = "MCP_Service"
  }]
  project_name = "default"
  object_groups = [{
    alert_methods     = ["Email", "Webhook"]
    contact_group_ids = ["contact_group_id-xxxx"]
    effect_start_at   = "00:00"
    effect_end_at     = "23:59"
    notify_templates = [{
      channel            = "email"
      notify_template_id = "2049xxxxxxx81664"
    }]
    object_group_id = "2090xxxxxxxxxx9875456"
    webhook_ids     = ["20058xxxxxx87200"]
  }]

}
