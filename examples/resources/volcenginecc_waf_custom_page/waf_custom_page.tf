resource "volcenginecc_waf_custom_page" "example" {
  client_ip    = "ALL"
  description  = "ccapi custom page required"
  host         = "www.testwaf.com"
  name         = "ccapi-custompage-required"
  page_mode    = 0
  policy       = 8
  url          = "/ccapi-required"
  body         = "blocked by ccapi required page"
  code         = "403"
  content_type = "text/html"
}
