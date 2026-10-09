resource "volcenginecc_tidb_branch" "example" {
  cluster_id       = "cluster_id-xxx"
  display_name     = "ccapi-tidb-branch-full"
  parent_id        = "parent_id"
  parent_timestamp = "YYYY-MM-DD'T'hh:mm:ss'Z'"
  password         = "Example@Test1234"

}
