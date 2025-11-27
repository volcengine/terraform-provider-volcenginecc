resource "volcenginecc_vke_cluster" "VKEClusterDemo" {
  project_name              = "default"
  name                      = "VKEClusterDemo"
  description               = "VKEClusterDemo"
  delete_protection_enabled = true
  cluster_config = {
    subnet_ids = [
      "subnet-***"
    ]
    api_server_public_access_enabled = true
    api_server_public_access_config = {
      public_access_network_config = {
        billing_type = 3
        bandwidth    = 1
        isp          = "BGP"
      }
    }
    resource_public_access_default_enabled = true
  }
  pods_config = {
    pod_network_mode = "VpcCniShared"
    vpc_cni_config = {
      subnet_ids = [
        "subnet-***"
      ]
      trunk_eni_enabled = false
    }
  }
  services_config = {
    service_cidrsv_4 = ["172.22.xxx.0/22"]
  }
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  kubernetes_version_create = "1.30"
}