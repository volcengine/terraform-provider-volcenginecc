resource "volcenginecc_tls_etl" "TLSEtlDemo" {
  dsl_type        = "NORMAL"
  description     = "ccapi-test"
  enable          = true
  from_time       = 1780302267
  name            = "ccapi-test-etltask"
  script          = "# æŒ‰kvæ‹†åˆ†contentå­—æ®µ\r\next_kv(\"content\", pair_sep=\"&\", kv_sep=\"=\", prefix=\"\", suffix=\"\", mode=\"overwrite\")\r\n# åˆ é™¤contentå­—æ®µ\r\nf_drop(\"content\")"
  source_topic_id = "b881e3cd-3c45-42e7-966f-fe98xxxxxx"
  target_resources = [
    {
      alias  = "test-dst12"
      region = "cn-shanghai"
    topic_id = "7ded0540-a2e6-48ea-9667-e0xxxxxx" },
    {
      alias  = "test-dst13"
      region = "cn-beijing"
    topic_id = "c7474a2b-60bf-4d1d-bb5e-664xxxxxx" }
  ]
  task_type = "Resident"
  to_time   = 1780475069
}