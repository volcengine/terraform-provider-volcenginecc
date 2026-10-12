resource "volcenginecc_waf_custom_bot" "example" {
  accurate = {
    accurate_rules = [{
      http_obj     = "request.header.User-Agent"
      obj_type     = 14
      opretar      = 6
      property     = 0
      value_string = "curl"
      }, {
      http_obj     = "request.method"
      obj_type     = 2
      opretar      = 2
      property     = 0
      value_string = "GET"
    }]
    logic = 1
  }
  action                      = "js"
  bot_type                    = "ccapiBotFull"
  description                 = "ccapi custom bot full"
  enable                      = 1
  host                        = "www.testwaf.com"
  verification_exemption_time = 300

}
