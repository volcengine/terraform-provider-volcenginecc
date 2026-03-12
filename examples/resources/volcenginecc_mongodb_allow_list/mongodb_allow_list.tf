resource "volcenginecc_mongodb_allow_list" "MongoDBAllowListDemo" {
  allow_list_name     = "mongodb_allowlist-1"
  allow_list_type     = "IPv4"
  project_name        = "default"
  allow_list_desc     = "test"
  allow_list_category = "Ordinary"
  allow_list = [
    "0.0.0.0/0",
    "127.0.0.1",
    "192.168.1.0/24"
  ]
  associated_instances = [
    { instance_id = "mongo-shard-d838exxxxx" },
    { instance_id = "mongo-replica-63axxxxx" }
  ]
}