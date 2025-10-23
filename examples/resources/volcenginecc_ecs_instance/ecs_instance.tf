resource "volcenginecc_ecs_instance" "EcsInstanceDemo" {
  password                    = "**********"
  instance_charge_type        = "PostPaid"
  instance_name               = "EcsInstanceDemo"
  spot_strategy               = "NoSpot"
  hostname                    = "EcsInstanceDemo"
  description                 = "EcsInstanceDemo Description"
  project_name                = "default"
  instance_type               = "ecs.g4i.large"
  deletion_protection         = false
  zone_id                     = "cn-beijing-a"
  deployment_set_id           = "dps-ydzccfzqjoaa98kxxxxx"
  deployment_set_group_number = 1
  image = {
    image_id = "image-aagd56zrw2jtdroxxxxx"
  }
  vpc_id = "vpc-rrco37ovjq4gv0x58zxxxxx"
  eip_address = {
    charge_type           = "PayByTraffic"
    bandwidth_mbps        = 5
    isp                   = "BGP"
    release_with_instance = false
    bandwidth_package_id  = ""
  }
  primary_network_interface = {
    security_group_ids  = ["sg-1v9zvjkmx14w51j8e73xxxxx"]
    subnet_id           = "subnet-rrwqhg3qzxfkv0x57gxxxxx"
    ipv_6_address_count = 1
  }
  secondary_network_interfaces = [
    {
      security_group_ids = ["sg-1v9zvjkmx14w51j8e73xxxxx"]
    subnet_id = "subnet-rrwqhg3qzxfkv0x57gxxxxx" }
  ]
  placement = {
    affinity                  = "Default"
    dedicated_host_cluster_id = ""
    tenancy                   = "Default"
    dedicated_host_id         = ""
  }
  system_volume = {
    size                 = 50
    delete_with_instance = true
    volume_type          = "ESSD_FlexPL"
  }
}