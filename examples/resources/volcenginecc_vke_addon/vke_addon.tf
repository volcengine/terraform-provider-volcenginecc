resource "volcenginecc_vke_addon" "VKEAddonDemo" {
  cluster_id       = "cd35mtki***"
  deploy_mode      = "Unmanaged"
  deploy_node_type = "Node"
  name             = "csi-nas"
  version          = "v1.2.7"
  config           = "{\"CsiNasDriver\":{\"Resources\":{\"Requests\":{\"Cpu\":\"0.01\",\"Memory\":\"20Mi\"},\"Limits\":{\"Cpu\":\"0.9\",\"Memory\":\"1900Mi\"}}},\"CsiProvisioner\":{\"Resources\":{\"Requests\":{\"Cpu\":\"0.01\",\"Memory\":\"20Mi\"},\"Limits\":{\"Cpu\":\"0.5\",\"Memory\":\"4Gi\"}}},\"LivenessProbe\":{\"Resources\":{\"Requests\":{\"Cpu\":\"0.01\",\"Memory\":\"20Mi\"},\"Limits\":{\"Cpu\":\"0.1\",\"Memory\":\"100Mi\"}}}}"
}