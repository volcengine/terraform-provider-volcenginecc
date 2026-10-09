resource "volcenginecc_waf_bot_analyse_protect_rule" "example" {
  action_type          = 1
  enable               = 1
  field                = "CLIENTIP"
  host                 = "www.testwaf.com"
  name                 = "ccapi-bot-analyse-required"
  path                 = "/ccapi-required"
  path_threshold       = 10
  rule_priority        = 0
  single_threshold     = 3
  statistical_duration = 30
  statistical_type     = "1"

}
