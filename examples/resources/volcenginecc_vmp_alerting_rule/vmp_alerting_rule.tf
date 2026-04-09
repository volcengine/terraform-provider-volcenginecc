resource "volcenginecc_vmp_alerting_rule" "VMPAlertingRuleDemo" {
  name        = "alert-rule-demo"
  description = "this is a alert rule demo"
  type        = "vmp/PromQL"
  query = {
    prom_ql      = "http_requests_total{method=\"PUT\", status=\"404\"}"
    workspace_id = "e9e4b146-0ba9-41ac-af2b-xxxxxx"
  }
  levels = [
    {
      level      = "P0"
      for        = "1s"
      comparator = "<"
      threshold  = 5
    },
    {
      level      = "P1"
      for        = "6m"
      comparator = "<="
      threshold  = 10
    },
    {
      level      = "P2"
      for        = "10m"
      comparator = "!="
      threshold  = 20
    }
  ]
  notify_policy_id       = "2890eb5e-f383-4a7b-810f-xxxxxxx"
  notify_group_policy_id = "a8490e9b-4e5d-4a86-xxxx-xxxxxxx"
  annotations = [
    {
      name  = "Key"
      value = "Value"
    }
  ]
  labels = [
    {
      name  = "LabelsKey"
      value = "LabelsValue"
    }
  ]
  status = "Running"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}