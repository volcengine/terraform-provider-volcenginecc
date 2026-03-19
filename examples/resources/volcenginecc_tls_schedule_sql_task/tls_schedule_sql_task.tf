resource "volcenginecc_tls_schedule_sql_task" "TLSScheduleSqlTaskDemo" {
  description         = "testdesc"
  dest_region         = "cn-beijing"
  dest_topic_id       = "42fedbb3-2c2a-4822-bc30-9c5cc****"
  process_end_time    = 0
  process_sql_delay   = 60
  task_type           = 0
  process_start_time  = 1773072000
  process_time_window = "@m-15m,@m"
  query               = "* | SELECT * limit 10"
  request_cycle = {
    cron_tab       = "0 19 * * *"
    cron_time_zone = "GMT+08:00"
    time           = 2
    type           = "Fixed"
  }
  source_topic_id = "22fca26e-a776-4925-a3ca-a9bb6d***"
  status          = 1
  task_name       = "test"
}