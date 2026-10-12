resource "volcenginecc_vke_nodepool_scaling_policy" "example" {
  cluster_id       = "cluster_id-xxxx"
  node_pool_id     = "node_pool_id-xxxxx"
  type             = "Recurrence"
  adjustment_value = 2
  enabled          = true
  scheduled_policy = {
    launch_time         = "2026-09-19T00:00Z"
    recurrence_type     = "Daily"
    recurrence_value    = "1"
    recurrence_end_time = "2026-10-19T00:00Z"
  }
  scheduled_instance_policy = {
    elastic_scheduled_instance_type = "Esi"
    start_delivery_time             = "2026-09-19T04:00Z"
    end_delivery_time               = "2026-09-19T04:10Z"
    best_effort                     = false
  }

}
