resource "volcenginecc_redis_allow_list" "RedisAllowListDemo" {
  allow_list          = "127.0.0.1,0.0.0.0/0,192.168.0.0/24"
  allow_list_name     = "ccapi_test"
  allow_list_desc     = "this is a test"
  allow_list_category = "Ordinary"
  project_name        = "default"
  security_group_bind_infos = [
    {
      bind_mode = "AssociateEcsIp"
    security_group_id = "sg-w06pxjgvtds0865yxxxxxx" },
    {
      bind_mode = "AssociateEcsIp"
    security_group_id = "sg-rrco3fkzzy0wv0x5xxxxxx" },
    {
      bind_mode = "IngressDirectionIp"
    security_group_id = "sg-w07syi789s74865yxxxxxx" }
  ]
  instance_ids = [
    "redis-cnlfwpqoohxxxxxx",
    "redis-cnlfenhsypxxxxxx"
  ]
}