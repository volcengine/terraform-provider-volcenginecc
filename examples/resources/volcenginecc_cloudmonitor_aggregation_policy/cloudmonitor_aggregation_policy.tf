resource "volcenginecc_cloudmonitor_aggregation_policy" "primary_cloudmonitor_aggregationpolicy_full" {
  aggregations = [{
    group_by       = ["rule_id", "region"]
    group_interval = 60
    group_wait     = 30
    level          = "all"
    repeat_downgrade = {
      enable              = true
      max_repeat_interval = 3600
      mode                = "doubled"
    }
    repeat_interval = 120
  }]
  description = "完整all兜底及官方示例降频配置"
  is_default  = false
  name        = "ccapi-agp-9011199b-full"
}