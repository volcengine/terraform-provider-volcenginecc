resource "volcenginecc_cloudmonitor_event_rule" "CloudMonitorEventRuleDemo" {
  rule_name      = "ccapi-test-1001"
  description    = "autoscaling scale in error"
  event_bus_name = "default"
  event_source   = "autoscaling"
  event_type = [
    "autoscaling:ScalingGroup:ScaleInError"
  ]
  enable_state    = "enable"
  level           = "critical"
  notification_id = "2049745192xxxxxx"
  filter_pattern_input = jsonencode({
    Data = {
      autoscaling = ["1", "3"]
      ecs         = ["2"]
    }
  })
}