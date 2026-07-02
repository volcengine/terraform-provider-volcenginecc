resource "volcenginecc_waf_host_group" "Example" {
  description  = "测试域名组描述"
  project_name = "default"
  host_list    = ["www.testwaf.com"]
  name         = "test-host-group-full"
}