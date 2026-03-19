resource "volcenginecc_autoscaling_scaling_lifecycle_hook" "ScalingLifecycleHookDemo" {
  scaling_group_id       = "scg-yexxxxxxv6yn56hnvlm"
  lifecycle_hook_name    = "ccapi-dx-1"
  lifecycle_hook_type    = "SCALE_IN"
  lifecycle_hook_timeout = 1800
  lifecycle_hook_policy  = "CONTINUE"
  lifecycle_command = {
    command_id = "cmd-ybv3xxxxxxx51qxnx"
    parameters = "{\"KEY_PAIR_ID\":\"1024\",\"KEY_PAIR_TIMEOUT\":\"360\"}"
  }

}