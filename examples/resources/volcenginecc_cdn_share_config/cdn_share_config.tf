resource "volcenginecc_cdn_share_config" "CDNShareConfigDemo" {
  config_name = "CDNShareConfigDemo"
  config_type = "allow_referer_access_rule"
  project     = "default"
  allow_referer_access_rule = {
    allow_empty = true
    common_type = {
      ignore_case = true
      rules       = ["139.x.x.1", "139.x.x.2"]
    }
  }
}