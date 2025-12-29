resource "volcenginecc_rocketmq_instance" "RocketMQInstanceDemo" {
  allow_list_ids = [
    "acl-5380876c62044b658c3c7da4xxxx",
    "acl-3e998a9ef43e48eda1d07e2exxxx"
  ]
  ip_version_type = "IPv4"
  enable_ssl      = true
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  version            = "5.x"
  zone_id            = "cn-beijing-a,cn-beijing-c,cn-beijing-d"
  compute_spec       = "rocketmq.x2.2k"
  storage_space      = 300
  vpc_id             = "vpc-1a1vgeo93ycxxxxxxxjnuw"
  subnet_id          = "subnet-ij9s4h4xxxxxxx95wx4p"
  file_reserved_time = 72
  instance_name      = "RocketMQInstanceDemo"
  charge_detail = {
    charge_type = "PostPaid"
  }
  ssl_mode             = "permissive"
  network_types        = "PrivateNetwork"
  project_name         = "default"
  instance_description = "RocketMQInstanceDemo description"
  eip_id               = "eip-bt6jb362txxxxx2zbpbo"
}