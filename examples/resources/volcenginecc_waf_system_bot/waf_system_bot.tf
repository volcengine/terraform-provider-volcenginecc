resource "volcenginecc_waf_system_bot" "example" {
  host                        = "www.testwaf.com"
  rule_tag                    = "J000000000025"
  bot_type                    = "search_engine_bot"
  action                      = "js"
  project_name                = "default"
  verification_exemption_time = 300
}
