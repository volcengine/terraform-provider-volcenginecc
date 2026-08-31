resource "volcenginecc_redis_key_scan_job" "example" {
  instance_id             = "redis-xxxx"
  scan_key_num_per_second = 3000
  timeout_minutes         = 60
}
