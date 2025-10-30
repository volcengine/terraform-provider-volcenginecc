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
      key = "env"
    value = "test" }
  ]
  time_key       = "time"
  time_format    = "%Y-%m-%d %H:%M:%S"
  log_public_ip  = false
  topic_name     = "test"
  description    = "test"
  project_id     = "44a425f0-a6ef-4a****"
  enable_hot_ttl = false
}