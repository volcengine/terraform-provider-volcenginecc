resource "volcenginecc_vke_kubeconfig" "VKEKubeconfigDemo" {
  cluster_id     = "cd****"
  type           = "Private"
  valid_duration = 2
}