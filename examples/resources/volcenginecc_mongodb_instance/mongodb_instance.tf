resource "volcenginecc_mongodb_instance" "MongoDBInstanceDemo" {
  zone_id                        = "cn-beijing-a,cn-beijing-c,cn-beijing-d"
  vpc_id                         = "vpc-1a1vgeo9xxxcg8nvepjyxxxxx"
  subnet_id                      = "subnet-ij9s4hxxxs3k74o8cuxxxxx"
  db_engine                      = "MongoDB"
  db_engine_version              = "MongoDB_7_0"
  instance_type                  = "ShardedCluster"
  node_spec                      = "mongo.shard.2c4g"
  node_number                    = 3
  mongos_node_spec               = "mongo.mongos.2c4g"
  mongos_node_number             = 3
  shard_number                   = 2
  storage_space_gb               = 100
  config_server_node_spec        = "mongo.config.2c4g"
  config_server_storage_space_gb = 20
  super_account_name             = "****"
  super_account_password         = "*******"
  instance_name                  = "MongoDBInstanceDemo"
  instance_count                 = 1
  charge_type                    = "PostPaid"
  project_name                   = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  allow_list_ids = ["acl-c972e7b4ce4941a1a8d5xxxe57xxxxx", "acl-70dbb8ee8893467dbafxxxc964xxxxx"]
}