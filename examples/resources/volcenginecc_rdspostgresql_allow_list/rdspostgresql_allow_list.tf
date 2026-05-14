resource "volcenginecc_rdspostgresql_allow_list" "RdsPostgresqlAllowListDemo" {
  allow_list_category = "Default"
  allow_list_desc     = "test"
  allow_list_name     = "ccapi-dx-2"
  allow_list_type     = "IPv4"
  security_group_bind_infos = [{
    bind_mode         = "AssociateEcsIp"
    security_group_id = "sg-w06pxxxxx5yk9xgrubx"
    }, {
    bind_mode         = "IngressDirectionIp"
    ip_list           = ["100.70.0.0/16", "100.72.0.0/16", "100.73.0.0/16"]
    security_group_id = "sg-1v9zvjxxxxxj8e73vhtg8"
  }]
  user_allow_list = "1.2.4.0/24"
  associated_instances = [{
    instance_id = "postgres-cxxxx3b95"
  }]
}