resource "volcenginecc_vke_node_pool" "VKENodePoolDemo" {
  auto_scaling = {
    max_replicas     = 10
    min_replicas     = 0
    enabled          = true
    desired_replicas = 0
    priority         = 10
    subnet_policy    = "ZoneBalance"
  }
  cluster_id = "cd6gojxxxxxxxxxxx"
  kubernetes_config = {
    auto_sync_disabled = false
    containerd_config = {
      insecure_registries = [
      "example.com"]
      registry_proxy_configs = [
        {
          proxy_endpoints = [
          "https://example.com:8080"]
          registry = "example.com:8080"
        }
      ]
    }
    cordon = false
    kubelet_config = {
      cpu_manager_policy = "none"
      eviction_hard = [
        {
          key   = "memory.available"
          value = "15%"
        }
      ]
      kube_api_burst = 10
      kube_api_qps   = 5
      kube_reserved = [
        {
          name     = "memory"
          quantity = "200m"
        }
      ]
      max_pods              = 110
      registry_burst        = 10
      registry_pull_qps     = 5
      serialize_image_pulls = true
      system_reserved = [
        {
          name     = "memory"
          quantity = "200m"
        }
      ]
      topology_manager_policy = "none"
      topology_manager_scope  = "container"
    }
    labels = [
      {
        key   = "label-key"
        value = "label-value"
      }
    ]
    name_prefix       = "name-prefix"
    name_suffix       = "name-suffix"
    name_use_hostname = false
    taints = [
      {
        key    = "taint-key"
        value  = "taint-value"
        effect = "NoSchedule"
      }
    ]
  }
  management = {
    enabled = true
    remedy_config = {
      enabled   = true
      remedy_id = "R20260227xxxxxxxxxxx"
    }
  }
  name = "test"
  node_config = {
    instance_charge_type = "PostPaid"
    spot_strategy        = "SpotAsPriceGo"
    instances_distribution = {
      capacity_rebalance                       = true
      compensate_with_on_demand                = true
      on_demand_base_capacity                  = 0
      on_demand_percentage_above_base_capacity = 0
    }
    security = {
      security_group_ids = [
      "sg-1a14cxqxxxxxxxxxx"]
      security_strategies = [
      "Hids"]
      login = {
        ssh_key_pair_name = "MigrationKey-job-yecd7xxxxxxxxxx"
      }
    }
    data_volumes = [
      {
        mount_point = "/dev/vdb"
        size        = 100
        snapshot_id = "snap-3wvah8xxxxxxxxxx"
        type        = "ESSD_PL0"
      }
    ]
    system_volume = {
      size = 40
      type = "ESSD_PL0"
    }
    additional_container_storage_enabled = true
    deployment_set_group_number          = 0
    deployment_set_id                    = "dps-yedy0wxxxxxxxxxx"
    hostname                             = "node-host"
    image_id                             = "image-ybqi99xxxxxxxxxxx"
    initialize_script                    = "YmFzaCxxxxx9maxxxxx"
    instance_name                        = "node"
    instance_type_ids = [
    "ecs.g4il.large"]
    name_prefix          = "name-prefix"
    network_traffic_mode = "Standard"
    project_name         = "default"
    public_access_config = {
      billing_type = 2
      isp          = "BGP"
      bandwidth    = 5
    }
    public_access_enabled = false
    subnet_ids = [
    "subnet-ijifxxxxxo8cuxxxxx"]
    tags = [
      {
        key   = "key"
        value = "value"
      }
    ]
  }
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}
