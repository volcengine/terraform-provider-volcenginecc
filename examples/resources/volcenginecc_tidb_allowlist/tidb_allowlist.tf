resource "volcenginecc_tidb_allowlist" "primary_tidb_allowlist_full" {
  instance_id             = "替换为 TiDB 实例 ID"
  vpc_endpoint_id         = "替换为接入侧 PrivateLink 终端节点 ID"
  authorized_network_name = "ccapi-tidb-allowlist-full"
}