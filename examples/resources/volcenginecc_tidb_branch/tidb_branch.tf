resource "volcenginecc_tidb_branch" "example" {
  instance_id      = "替换为 TiDB 实例 ID"
  branch_name      = "ccapi-tidb-branch-full"
  parent_id        = "parent_id"
  parent_timestamp = "YYYY-MM-DD'T'hh:mm:ss'Z'"
  password         = "Example@Test1234"

}
