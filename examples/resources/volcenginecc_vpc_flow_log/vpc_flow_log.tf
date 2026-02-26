resource "volcenginecc_vpc_flow_log" "VPCFlowLogDemo" {
  aggregation_interval = 10
  description          = "this is a test flow log"
  flow_log_name        = "FlowLog-ccapi-test"
  log_project_name     = "ccapi-test-flow-log"
  log_topic_name       = "test-flow-log"
  project_name         = "default"
  resource_id          = "vpc-rrco37ovjq4gv0xxxxxx"
  resource_type        = "vpc"
  status               = "Active"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  traffic_type = "All"
}