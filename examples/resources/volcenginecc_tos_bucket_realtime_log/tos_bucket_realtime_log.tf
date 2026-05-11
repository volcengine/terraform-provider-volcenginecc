resource "volcenginecc_tos_bucket_realtime_log" "TOSBucketRealtimeLogDemo" {
  bucket_name = "ccapi-test"
  real_time_log = {
    role              = "TOSLogArchiveTLSRole"
    use_service_topic = true
    tls_project_id    = "569b6ea5-xxxxxx-90d47ff07774"
    tls_topic_id      = "22fca26e-xxxxxx-a9bb6d3fb9bd"
  }
}