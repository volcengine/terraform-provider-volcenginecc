resource "volcenginecc_emr_node_group" "EMRNodeGroupDemo" {
  cluster_id         = "emr-cluster-123456789xxxxx"
  zone_id            = "cn-beijing-a"
  node_group_name    = "CoreGroup-1002"
  node_group_type    = "CORE"
  with_public_ip     = false
  target_disk_size   = 120
  node_count         = 2
  charge_type        = "POST"
  ecs_instance_types = ["ecs.r1ie.xlarge"]
  subnet_ids         = ["subnet-rrwqhg3qzxfkv0xxxxxxx"]
  system_disk = {
    volume_type = "ESSD_FlexPL"
    size        = 120
  }
  data_disks = [
    {
      volume_type = "ESSD_FlexPL"
      size        = 80
      count       = 1
    }
  ]
}