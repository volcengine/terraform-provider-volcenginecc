resource "volcenginecc_vepfs_mount_service" "VEPFSMountServiceDemo" {
  mount_service_name = "test"
  project            = "default"
  node_type          = "ecs.g4i.large"
  subnet_id          = "subnet-1a185dj3sh1q8*****"
  vpc_id             = "vpc-1a17v350xa8lc8n*****"
  zone_id            = "cn-beijing-a"
  attach_file_systems = [
    {
      file_system_id = "vepfs-cnbj607*****"
      customer_path  = "/vepfs-cnbj*****"
    }
  ]
}