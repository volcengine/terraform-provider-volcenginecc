resource "volcenginecc_ecs_launch_template_version" "EcsLaunchTemplateVersionDemo" {
  deployment_set_group_number = 0
  deployment_set_id           = "dps-ydzc5xxxxxx5b9vu"
  description                 = "ecs_launch_template_version test"
  eip = {
    bandwidth                       = 1
    bandwidth_package_id            = "bwp-1a1fvbxxxxxx8nvepl9jzfmf"
    billing_type                    = 0
    isp                             = "BGP"
    release_with_instance           = true
    security_protection_instance_id = 0
    security_protection_types       = ["AntiDDoS_Enhanced"]
  }
  host_name             = "myname"
  image_id              = "image-yzpvmxxxxxxgqcsdamq"
  image_name            = "Ubuntu 24.04 with LLM Knowledge Base 64 bit"
  instance_charge_type  = "PostPaid"
  instance_name         = "test-version"
  instance_type_id      = "ecs.g4i.large"
  keep_image_credential = false
  key_pair_name         = "MigrationKey-job-yecd7dromy38dfaxgxt8"
  launch_template_id    = "lt-yefdzjlbxxxxxxylmkj"
  network_interfaces = [
    {
      security_group_ids = ["sg-1jolcljxxxxxx1n7ampe70vpx"]
    subnet_id = "subnet-btd4nxxxxxxb2tl1jdsb" }
  ]
  project_name = "default"
  scheduled_instance = {
    scheduled_instance_description = "test"
    scheduled_instance_name        = "test-template"
  }
  security_enhancement_strategy = "Active"
  spot_strategy                 = "NoSpot"
  suffix_index                  = 1
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  unique_suffix       = false
  user_data           = "ZWNobyBoZWxsbyBlY3Mh"
  version_description = "test"
  volumes = [
    {
      delete_with_instance            = true
      extra_performance_iops          = 0
      extra_performance_throughput_mb = 0
      extra_performance_type_id       = ""
      size                            = 50
      snapshot_id                     = ""
    volume_type = "ESSD_PL0" }
  ]
  vpc_id  = "vpc-1jolcldhxxxxxxmq5q7yms"
  zone_id = "cn-beijing-a"
}