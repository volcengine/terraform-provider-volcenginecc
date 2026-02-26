resource "volcenginecc_rdsmysql_allow_list" "RdsMysqlAllowListDemo" {
  user_allow_list     = ["1.2.3.4"]
  allow_list_category = "Default"
  allow_list_desc     = "this is a test"
  allow_list_name     = "ccapi-test-2"
  allow_list_type     = "IPv4"
  project_name        = "default"
  security_group_bind_infos = [
    {
      bind_mode         = "AssociateEcsIp"
      ip_list           = null
      security_group_id = "sg-1a10axxxxxvepkdgqgnu"
    security_group_name = "Default" },
    {
      bind_mode         = "IngressDirectionIp"
      ip_list           = ["100.70.0.0/16", "100.72.0.0/16", "0.0.0.0/0"]
      security_group_id = "sg-3nqt4kwxxxxx931ebkntmrc"
    security_group_name = "Default" }
  ]
}