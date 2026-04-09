resource "volcenginecc_vefaas_timer" "VEFAASTimerDemo" {
  function_id        = "o1zxxx"
  name               = "ccapi-test-timer"
  description        = "ccapi test timer"
  enabled            = false
  crontab            = "*/30 * * * *"
  payload            = "{\"body\": \"create event payload\"}"
  enable_concurrency = false
  retries            = 5
}