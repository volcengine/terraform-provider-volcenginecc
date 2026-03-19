resource "volcenginecc_emr_cluster" "EMRClusterDemo" {
  project_name    = "default"
  cluster_name    = "ccapi-tftest-1"
  cluster_type    = "Hadoop"
  release_version = "3.19.0"
  application_names = [
    "HDFS",
    "YARN",
    "SPARK3"
  ]
  deploy_mode         = "SIMPLE"
  security_mode       = "SIMPLE"
  charge_type         = "POST"
  vpc_id              = "vpc-rrco37ovjq4gv0xxxxxxxxx"
  security_group_id   = "sg-1v9zvjkmx14w51j8xxxxxxxx"
  history_server_mode = "LOCAL"
  node_attribute = {
    zone_id      = "cn-beijing-a"
    ecs_iam_role = "VEECSforEMRRole"
  }
  node_group_attributes = [
    {
      zone_id            = "cn-beijing-a"
      subnet_ids         = ["subnet-rrwqhg3qzxfkv0xxxxxxxx"]
      node_group_name    = "MasterGroup"
      ecs_key_pair_name  = "pln-test"
      ecs_password       = ""
      bandwidth          = 8
      charge_type        = "POST"
      node_group_type    = "MASTER"
      with_public_ip     = false
      ecs_instance_types = ["ecs.r1ie.xlarge"]
      node_count         = 1
      system_disk = {
        volume_type = "ESSD_FlexPL"
        size        = 120
      }
      data_disks = [
        {
          volume_type = "ESSD_FlexPL"
          size        = 80
        count = 1 }
    ] },
    {
      zone_id            = "cn-beijing-a"
      subnet_ids         = ["subnet-rrwqhg3qzxfkv0x5xxxxxx"]
      node_group_name    = "CoreGroup"
      ecs_key_pair_name  = "pln-test"
      ecs_password       = ""
      bandwidth          = 8
      charge_type        = "POST"
      node_group_type    = "CORE"
      with_public_ip     = false
      ecs_instance_types = ["ecs.r1ie.xlarge"]
      node_count         = 2
      system_disk = {
        volume_type = "ESSD_FlexPL"
        size        = 80
      }
      data_disks = [
        {
          volume_type = "ESSD_FlexPL"
          size        = 80
        count = 4 }
    ] },
    {
      zone_id            = "cn-beijing-a"
      subnet_ids         = ["subnet-rrwqhg3qzxfkv0xxxxxxxx"]
      node_group_name    = "TaskGroup-1"
      ecs_key_pair_name  = "pln-test"
      ecs_password       = ""
      bandwidth          = 8
      charge_type        = "POST"
      node_group_type    = "TASK"
      with_public_ip     = false
      ecs_instance_types = ["ecs.r1ie.xlarge"]
      node_count         = 1
      system_disk = {
        volume_type = "ESSD_FlexPL"
        size        = 80
      }
      data_disks = [
        {
          volume_type = "ESSD_FlexPL"
          size        = 80
        count = 4 }
      ]
    }
  ]
  application_extras = [
    {
      application_name = "HIVE"
      connection_type  = "BUILT_IN_MYSQL"
      connection_id    = ""
      application_configs = [
        {
          config_file_name        = "hive-site.xml"
          config_item_key         = "hive.metastore.warehouse.dir"
          config_item_value       = "tos://ccapi-test-tos-1/managed"
          deleted                 = false
          component_instance_name = ""
          component_name          = ""
          effective_scope = {
            effective_type   = "CLUSTER"
            node_group_ids   = []
            node_group_names = []
            node_group_types = []
            node_names       = []
            node_ids         = []
            component_names  = []
          }
        },
        {
          config_file_name        = "hive-site.xml"
          config_item_key         = "hive.metastore.warehouse.external.dir"
          config_item_value       = "tos://ccapi-test-tos-1/external"
          deleted                 = false
          component_instance_name = ""
          component_name          = ""
          effective_scope = {
            effective_type   = "CLUSTER"
            node_group_ids   = []
            node_group_names = []
            node_group_types = []
            node_names       = []
            node_ids         = []
            component_names  = []
          }
        },
        {
          config_file_name        = "hive-site.xml"
          config_item_key         = "hive.metastore.warehouse.dir"
          config_item_value       = "tos://ccapi-test-tos-1/managed"
          deleted                 = false
          component_instance_name = ""
          component_name          = ""
          effective_scope = {
            effective_type   = "CLUSTER"
            node_group_ids   = []
            node_group_names = []
            node_group_types = []
            node_names       = []
            node_ids         = []
            component_names  = []
          }
        },
        {
          config_file_name        = "hive-site.xml"
          config_item_key         = "hive.metastore.warehouse.external.dir"
          config_item_value       = "tos://ccapi-test-tos-1/external"
          deleted                 = false
          component_instance_name = ""
          component_name          = ""
          effective_scope = {
            effective_type   = "CLUSTER"
            node_group_ids   = []
            node_group_names = []
            node_group_types = []
            node_names       = []
            node_ids         = []
            component_names  = []
          }
        }
      ]
      application_component_layouts = [
        {
          component_name = ""
          effective_scope = {
            effective_type   = "CLUSTER"
            node_group_ids   = []
            node_group_names = []
            node_group_types = []
            node_names       = []
            node_ids         = []
            component_names  = []
          }
        }
      ]
    }
  ]
  bootstrap_scripts = [
    {
      script_name = "ccapi-test-script"
      script_type = "BOOTSTRAP"
      script_path = "tos://ccapi-test-tos-1/managed/"
      script_args = ""
      priority    = "1"
      effective_scope = {
        effective_type   = "CLUSTER"
        component_names  = []
        node_group_ids   = []
        node_group_names = []
        node_group_types = []
        node_ids         = []
        node_names       = []
      }
      execution_moment        = "AFTER_APPLICATION_STARTED"
      execution_fail_strategy = "FAILED_BLOCK"
    }
  ]
  tags = [
    {
      "key"   = "env"
      "value" = "test"
    }
  ]
  lifecycle {
    ignore_changes = [
      application_names,
      application_extras,
      bootstrap_scripts,
      node_group_attributes
    ]
  }
}