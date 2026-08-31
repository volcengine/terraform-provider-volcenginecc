resource "volcenginecc_ecs_dedicated_host_cluster" "example" {
  dedicated_host_cluster_name = "ccapi-ddhc-full"
  description                 = "ccapi dedicated host cluster full"
  zone_id                     = "cn-beijing-a"
}
