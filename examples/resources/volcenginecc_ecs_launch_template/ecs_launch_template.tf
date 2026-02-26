resource "volcenginecc_ecs_launch_template" "EcsLaunchTemplateDemo" {
  launch_template_version = {
    deployment_set_group_number = 0
    deployment_set_id           = "dps-ydzc5xxxxxxkak3c5b9vu"
    description                 = "ecs_launch_template test"
    eip = {
      bandwidth                       = 1
      bandwidth_package_id            = "bwp-1a1fvbxxxxxxvepl9jzfmf"
      billing_type                    = 0
      isp                             = "BGP"
      release_with_instance           = true
      security_protection_instance_id = 0
      security_protection_types       = ["AntiDDoS_Enhanced"]
    }
    host_name             = "myname"
    hpc_cluster_id        = ""
    image_id              = "image-yzpvmk9xxxxxxgqcsdamq"
    image_name            = "Ubuntu 24.04 with LLM Knowledge Base 64 bit"
    instance_charge_type  = "PostPaid"
    instance_name         = "ccapi-dx-0"
    instance_type_id      = "ecs.g4i.large"
    keep_image_credential = false
    key_pair_name         = "MigrationKey-job-yecd7dromy38dfaxgxt8"
    network_interfaces = [
      {
        security_group_ids = ["sg-1jolcljxxxxxxpe70vpx"]
      subnet_id = "subnet-btd4nhxxxxxxb2tl1jdsb" }
    ]
    project_name = "default"
    scheduled_instance = {
      scheduled_instance_description = "test"
      scheduled_instance_name        = "test-template"
    }
    security_enhancement_strategy = "Active"
    spot_price_limit              = 0
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
    vpc_id  = "vpc-1jolcldhxxxxxxq7yms"
    zone_id = "cn-beijing-a"
  }
  launch_template_name         = "test"
  launch_template_project_name = "default"
  launch_template_tags = [
    {
      key = "env"
    value = "test" }
  ]
}