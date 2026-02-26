resource "volcenginecc_alb_customized_cfg" "ALBCustomizedCfgDemo" {
  customized_cfg_name    = "ccapi-test"
  description            = "testdesc"
  customized_cfg_content = "ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256;\r\nssl_protocols TLSv1.2 TLSv1.1 TLSv1;\r\nclient_max_body_size 60M;\r\nkeepalive_timeout 77s;\r\nproxy_connect_timeout 4s;\r\nproxy_request_buffering off;\r\nproxy_ignore_client_abort off;\r\nproxy_read_timeout 60s;\r\nproxy_send_timeout 60s;\r\nclient_header_timeout 60s;\r\nclient_body_timeout 60s;\r\nsend_timeout 60s;\r\nssl_verify_depth 3;\r\n"
  project_name           = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}