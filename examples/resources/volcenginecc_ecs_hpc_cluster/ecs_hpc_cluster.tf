resource "volcenginecc_ecs_hpc_cluster" "ECSHpcClusterDemo" {
  name         = "ECSHpcClusterDemo"
  zone_id      = "cn-beijing-a"
  description  = "ECSHpcClusterDemo description"
  project_name = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}