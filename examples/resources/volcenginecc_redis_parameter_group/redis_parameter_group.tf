resource "volcenginecc_redis_parametergroup" "RedisParameterGroupDemo" {
  engine_version = "6.0"
  description    = "用于测试的Redis 6.0自定义参数组"
  param_values = [
    {
      value = "allkeys-lru"
      name  = "maxmemory-policy"
    },
    {
      value = "300"
      name  = "proxy-client-idle-timeout"
    }
  ]
  name = "test-redis6-param-group"
}