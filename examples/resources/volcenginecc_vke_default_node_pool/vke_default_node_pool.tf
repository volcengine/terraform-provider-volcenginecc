resource "volcenginecc_vke_default_node_pool" "VkeDefaultNodePoolDemo" {
  cluster_id = "cd60pbhxxxxxxc4f31a0"
  kubernetes_config = {
    auto_sync_disabled = false
    cordon             = false
    labels = [
      {
        key = "env"
      value = "test" }
    ]
    name_prefix       = "name-prefix"
    name_suffix       = "name-suffix"
    name_use_hostname = false
    taints = [
      {
        key   = "taint-key"
        value = "taint-value"
      effect = "NoSchedule" }
    ]
  }
  node_config = {
    initialize_script = "YmFzaCBteV9maWxlLnNo"
    name_prefix       = "name-prefix"
    security = {
      login = {
        password = "RHgxMTIyMzM/"
      }
      security_group_ids  = ["sg-1c0e5jxxxxxx5e8j70agemnk"]
      security_strategies = ["Hids"]
    }
    tags = [
      {
        key = "env"
      value = "test" }
    ]
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}