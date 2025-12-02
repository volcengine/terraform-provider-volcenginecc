resource "volcenginecc_vke_node_pool" "VKENodePoolDemo" {
  cluster_id = "cd4iklanmxxxb9ixxxxx"
  name       = "VKENodePoolDemo"
  kubernetes_config = {
    labels = [
      {
        key = "labels-key"
      value = "labels-value" }
    ]
    taints = [
      {
        key   = "taints-key"
        value = "taints-value"
      effect = "NoSchedule" }
    ]
    cordon      = false
    name_prefix = "name-prefix"
    kubelet_config = {
      feature_gates = {
        qo_s_resource_manager = true
      }
      topology_manager_scope  = "container"
      topology_manager_policy = "none"
      max_pods                = 110
      kube_api_qps            = 5
      kube_api_burst          = 10
      kube_reserved = [
        {
          name = "memory"
        quantity = "200m" }
      ]
      system_reserved = [
        {
          name = "memory"
        quantity = "200m" }
      ]
      registry_pull_qps     = 5
      registry_burst        = 10
      serialize_image_pulls = true
      cpu_manager_policy    = "none"
    }
    auto_sync_disabled = false
    name_suffix        = "name-suffix"
    name_use_hostname  = false
  }
  node_config = {
    instance_type_ids = ["ecs.g3il.large"]
    subnet_ids        = ["subnet-3nr6sws8sxxx931ebscxxxxx"]
    image_id          = "image-ybqi99sxxx8rx7mxxxxx"
    system_volume = {
      size = 40
      type = "ESSD_PL0"
    }
    data_volumes = [
      {
        size        = 20
        type        = "ESSD_PL0"
        mount_point = "/dev/vdb"
      snapshot_id = "snap-3wpmsnixxx55inqxxxxx" }
    ]
    initialize_script = "YmFzaCBteV9maWxxxxxx"
    security = {
      security_group_ids  = ["sg-3hitecg7d6xxx3nkipkyxxxxx"]
      security_strategies = ["Hids"]
      login = {
        password = "*******"
      }
    }
    additional_container_storage_enabled = true
    instance_charge_type                 = "PostPaid"
    name_prefix                          = "name-prefix"
    tags = [
      {
        key = "env"
      value = "test" }
    ]
    spot_strategy = "SpotAsPriceGo"
    instances_distribution = {
      capacity_rebalance                       = true
      compensate_with_on_demand                = true
      on_demand_base_capacity                  = 0
      on_demand_percentage_above_base_capacity = 0
    }
    project_name          = "default"
    public_access_enabled = false
    public_access_config = {
      billing_type = 2
      isp          = "BGP"
      bandwidth    = 5
    }
    hostname                    = "node-host"
    instance_name               = "node"
    network_traffic_mode        = "Standard"
    deployment_set_id           = "dps-ydzccfzqjxxxx8kxxxxx"
    deployment_set_group_number = 1
  }
  auto_scaling = {
    max_replicas     = 10
    min_replicas     = 0
    enabled          = true
    desired_replicas = 0
    priority         = 10
    subnet_policy    = "ZoneBalance"
  }
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  management = {
    enabled = true
    remedy_config = {
      enabled   = true
      remedy_id = "R202511251441xxxxuv230vsiu92xxxxx"
    }
  }
}