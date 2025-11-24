resource "volcenginecc_autoscaling_scaling_group" "AutoScalingScalingGroupDemo" {
  health_check_type = "ECS"
  scaling_mode      = "release"
  instances_distribution = {
    compensate_with_on_demand                = true
    on_demand_base_capacity                  = 0
    on_demand_percentage_above_base_capacity = 50
    spot_instance_remedy                     = true
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  launch_template_overrides = [
    {
      instance_type = "ecs.g4il.large"
    price_limit = 0 }
  ]
  launch_template_id = "lt-ye9dsl3ucffig992xxxxfncd"
  project_name       = "default"
  server_group_attributes = [
    {
      port            = 8080
      server_group_id = "rsp-13g2ktevofhfk3n6nxxxu4mqjq0o"
      type            = "CLB"
    weight = 50 },
    {
      port            = 8080
      server_group_id = "rsp-bde4j3yuu8008dv4xxxx0nkn203g"
      type            = "ALB"
    weight = 50 }
  ]
  launch_template_version   = "Default"
  subnet_ids                = ["subnet-3nrd6qaq3log0931ecxxxh3re2r"]
  default_cooldown          = 300
  desire_instance_number    = -1
  instance_terminate_policy = "OldestScalingConfigurationWithOldestInstance"
  max_instance_number       = 6
  min_instance_number       = 1
  multi_az_policy           = "PRIORITY"
  scaling_group_name        = "AutoScalingScalingGroupDemo"
}