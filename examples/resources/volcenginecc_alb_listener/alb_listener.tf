resource "volcenginecc_alb_listener" "ALBListenerDemo" {
  load_balancer_id = "alb-bdazuxj87hts8dv40noxxxxx"
  listener_name    = "ALBListenerDemo"
  protocol         = "HTTPS"
  port             = 80
  enabled          = "on"
  server_group_id  = "rsp-1pff7rhpf5e68845wfah2xxxx"
  acl_status       = "on"
  acl_type         = "black"
  acl_ids = [
    "acl-xp8avgtjnmrk54ov5fyoxxxx"
  ]
  customized_cfg_id = "ccfg-xoblsk63beo054ov5el5xxxx"
  description       = "ALBListenerDemo description"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  certificate_source         = "cert_center"
  cert_center_certificate_id = "cert-a126f867a19141618ff316a52a50xxxx"
  ca_certificate_source      = "pca_root"
  pca_root_ca_certificate_id = "pca_root_xxxxxx"
  enable_http_2              = "on"
  enable_quic                = "off"
  domain_extensions = [
    {
      cert_center_certificate_id = "cert-7718b7e7afa94e4db38ec9092cxxxxxx"
      certificate_source         = "cert_center"
      domain                     = ""
    }
  ]
}