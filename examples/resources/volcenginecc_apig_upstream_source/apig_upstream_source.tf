resource "volcenginecc_apig_upstream_source" "UpstreamSourceNacosDemo" {
  gateway_id  = "gd3s9vbk7npja181xxxxx"
  comments    = "upstreamSourceNacosDemo"
  source_type = "Nacos"
  source_spec = {
    nacos_source = {
      nacos_id = "nd3thmnjdl46p917xxxxx"
      auth_config = {
        basic = {
          username = "nacos"
          password = "******"
        }
      }
    }
  }
}