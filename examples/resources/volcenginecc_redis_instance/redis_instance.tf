resource "volcenginecc_redis_instance" "RedisInstanceDemo" {
  instance_name       = "RedisInstance"
  sharded_cluster     = 1
  password            = "********"
  node_number         = 2
  shard_capacity      = 512
  shard_number        = 2
  engine_version      = "5.0"
  vpc_id              = "vpc-13f8xxxx"
  subnet_id           = "vpc_subnet-xxxx"
  deletion_protection = "disabled"
  charge_type         = "PostPaid"
  port                = 6381
  project_name        = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  multi_az = "enabled"
  configure_nodes = [
    {
    az = "cn-beijing-a" },
    {
    az = "cn-beijing-b" }
  ]
}