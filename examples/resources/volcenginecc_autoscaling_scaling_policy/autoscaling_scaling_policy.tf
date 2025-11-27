resource "volcenginecc_autoscaling_scaling_policy" "AutoScalingScalingPolicyDemo" {
  scaling_group_id    = "scg-ye9vu9ztco9ht5lxxxxx"
  scaling_policy_name = "AutoScalingScalingPolicyDemo-定时任务"
  scaling_policy_type = "Scheduled"
  adjustment_type     = "PercentChangeInCapacity"
  adjustment_value    = 1
  cooldown            = 86400
  scheduled_policy = {
    launch_time         = "2025-12-21T11:58Z"
    recurrence_end_time = ""
    recurrence_type     = ""
    recurrence_value    = ""
  }
  is_enabled_policy = false
}