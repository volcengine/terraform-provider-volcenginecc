resource "volcenginecc_rdsmysql_endpoint" "RdsMysqlEndpointDemo" {
  auto_add_new_nodes               = false
  connection_mode                  = "Proxy"
  connection_pool_type             = "Transaction"
  description                      = "this is a description"
  enable_connection_persistent     = true
  endpoint_name                    = "ccapi-test-1"
  endpoint_type                    = "Custom"
  idle_connection_reclaim          = true
  implicit_trans_split             = true
  instance_id                      = "mysql-2e229xxxxxd5"
  master_node_routing              = false
  master_protector_timeout         = 60
  multi_statements_mode            = "Strict"
  nodes                            = "Primary,mysql-2e229xxxxxdd5-r442d,mysql-2e2xxxxxdd5-r8d0f"
  overload_protection              = true
  read_only_node_distribution_type = "RoundRobinCustom"
  read_only_node_max_delay_time    = 30
  read_only_node_weights = [
    {
      node_id   = ""
      node_type = "Primary"
    weight = 100 },
    {
      node_id   = "mysql-2e2xxxxx4dd5-r442d"
      node_type = "ReadOnly"
    weight = 200 },
    {
      node_id   = "mysql-2e22xxxxxdd5-r8d0f"
      node_type = "ReadOnly"
    weight = 300 }
  ]
  read_write_mode     = "ReadWrite"
  read_write_spliting = true
}