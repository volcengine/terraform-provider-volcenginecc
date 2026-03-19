resource "volcenginecc_tls_topic" "TlsTopicDemo" {
  ttl             = 187
  hot_ttl         = 8
  cold_ttl        = 79
  archive_ttl     = 100
  shard_count     = 2
  auto_split      = true
  max_split_shard = 256
  tags = [
    {
      key = "key1"
    value = "v1" }
  ]
  time_key       = "time"
  time_format    = "%Y-%m-%d %H:%M:%S"
  log_public_ip  = false
  topic_name     = "ccapi-test"
  description    = "test"
  project_id     = "c6fef4c1-041f-434e-b0f4-d5e9*****"
  enable_hot_ttl = false
  allow_consume  = false
}