resource "volcenginecc_redis_instance" "RedisInstanceDemo" {
  configure_nodes = [
    {
      az = "cn-beijing-a"
    },
    {
      az = "cn-beijing-b"
    }
  ]
  multi_az = "enabled"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  project_name        = "default"
  vpc_id              = "vpc-xxxxx"
  subnet_id           = "subnet-xxxxx"
  deletion_protection = "enabled"
  port                = 9999
  auto_renew          = false
  charge_type         = "PostPaid"
  engine_version      = "6.0"
  shard_capacity      = 512
  shard_number        = 2
  node_number         = 2
  allow_list_ids      = ["acl-cnlfwmfaqdefxxxxx"]
  password            = "********"
  sharded_cluster     = 1
  instance_name       = "ccapi-auto-test"
  no_auth_mode        = "open"
  parameter_group_id  = "DefaultParamGroupId-6.0"
  continuous_backup   = true
}