resource "volcenginecc_tls_alarm_webhook_integration_type" "Example" {
  webhook_name   = "ccapi-test-1002replace"
  webhook_secret = ""
  webhook_type   = "custom"
  webhook_url    = "https://open.feishu.cn/open-apis/bot/v2/hook/79a2f8dxxxxxxx"
  webhook_headers = [
    {
      key   = "Content-Type"
      value = "application/json"
    },
    {
      key   = "env"
      value = "pre"
    }
  ]
  webhook_method = "PUT"
}