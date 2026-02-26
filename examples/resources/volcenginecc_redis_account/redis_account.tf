resource "volcenginecc_redis_account" "RedisAccountDemo" {
  instance_id  = "redis-mlrfiqivjttxxxxx"
  account_name = "ccapi_test_1"
  description  = "this is a test"
  password     = "******"
  role_name    = "Administrator"
}