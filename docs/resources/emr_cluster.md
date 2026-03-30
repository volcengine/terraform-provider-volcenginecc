---
page_title: "volcenginecc_emr_cluster Resource - terraform-provider-volcenginecc"
subcategory: "EMR"
description: |-
  E-MapReduce (EMR) is an enterprise-grade big data analytics system based on the open-source Hadoop ecosystem. It is fully compatible with open source and provides integration and management for ecosystem components such as Hadoop, Spark, Hive, Hudi, and Iceberg.
---

# volcenginecc_emr_cluster (Resource)

E-MapReduce (EMR) is an enterprise-grade big data analytics system based on the open-source Hadoop ecosystem. It is fully compatible with open source and provides integration and management for ecosystem components such as Hadoop, Spark, Hive, Hudi, and Iceberg.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `charge_type` (String) Payment type. PRE means monthly subscription, POST means pay-as-you-go.
- `cluster_name` (String) Cluster name.
- `cluster_type` (String) Cluster type.
- `node_attribute` (Attributes) Cluster global node information. (see [below for nested schema](#nestedatt--node_attribute))
- `release_version` (String) Cluster version.
- `security_group_id` (String) Cluster global security group ID. All ECS in node groups will join this security group.
- `vpc_id` (String) Vpc ID。

### Optional

- `application_extras` (Attributes List) Cluster service extension information list, including custom configuration items for services, custom deployment topology settings for service components, and metadata connection configuration information for services.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--application_extras))
- `application_names` (Set of String) List of service names installed in the cluster. Creation-related field.
- `bootstrap_scripts` (Attributes Set) Cluster bootstrap script list.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--bootstrap_scripts))
- `charge_pre_config` (Attributes) Monthly subscription configuration parameters. Required when chargeType=PRE. (see [below for nested schema](#nestedatt--charge_pre_config))
- `deploy_mode` (String) Deployment mode. SIMPLE means simple mode, HIGH_AVAILABLE means high availability mode.
- `history_server_mode` (String) HistoryServer mode: LOCAL stores active data within the cluster, PHS stores active data outside the cluster.
- `node_group_attributes` (Attributes Set) Node group property list.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--node_group_attributes))
- `project_name` (String) Project to which the resource belongs. Default is 'default'. Each resource can belong to only one project. Only letters, numbers, underscores '_', dots '.', and hyphens '-' are allowed. Maximum length is 64 characters.
- `security_mode` (String) Security mode.
- `tags` (Attributes Set) Tag list.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `account_id` (Number) Account ID.
- `applications` (Attributes Set) Installed service list for cluster. Read-only field.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--applications))
- `cluster_domain_names` (Set of String) Cluster DNS domain suffix list.
- `cluster_id` (String) Cluster ID.
- `cluster_state` (String) Cluster status.
- `created_time` (Number) Cluster creation time.
- `creator_id` (Number) Creator ID.
- `creator_name` (String) Creator name.
- `ecs_image_id` (String) ECS image ID.
- `expire_time` (Number) Cluster expiration time.
- `id` (String) Uniquely identifies the resource.
- `ready_time` (Number) Cluster creation completion time.
- `state_change_reason` (Attributes) Status change reason. (see [below for nested schema](#nestedatt--state_change_reason))
- `terminate_time` (Number) Cluster termination time.

<a id="nestedatt--node_attribute"></a>
### Nested Schema for `node_attribute`

Optional:

- `ecs_iam_role` (String) ECS instance role.
- `zone_id` (String) Zone ID.


<a id="nestedatt--application_extras"></a>
### Nested Schema for `application_extras`

Optional:

- `application_component_layouts` (Attributes Set) Custom deployment topology list for service components.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--application_extras--application_component_layouts))
- `application_configs` (Attributes List) Custom configuration parameter list for services.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--application_extras--application_configs))
- `application_name` (String) Application name.
- `connection_id` (String) Metadata connection ID.
- `connection_type` (String) Metadata connection type. BUILT_IN_MYSQL: built-in database. EXTERNAL_MYSQL: external database. HIVE_METASTORE: HMS.

<a id="nestedatt--application_extras--application_component_layouts"></a>
### Nested Schema for `application_extras.application_component_layouts`

Optional:

- `component_name` (String) Component name.
- `effective_scope` (Attributes) Component layout scope. (see [below for nested schema](#nestedatt--application_extras--application_component_layouts--effective_scope))

<a id="nestedatt--application_extras--application_component_layouts--effective_scope"></a>
### Nested Schema for `application_extras.application_component_layouts.effective_scope`

Optional:

- `component_names` (Set of String) Component name list. Required when EffectiveType=COMPONENT_NAME.
- `effective_type` (String) Effective type. CLUSTER, NODE_GROUP_NAME, NODE_GROUP_ID, NODE_GROUP_TYPE, NODE_NAME, NODE_ID, COMPONENT_NAME.
- `node_group_ids` (Set of String) Node group ID list. Required when EffectiveType=NODE_GROUP_ID.
- `node_group_names` (Set of String) Node group name list. Required when EffectiveType=NODE_GROUP_NAME.
- `node_group_types` (Set of String) Node group type list. Required when EffectiveType=NODE_GROUP_TYPE. Currently includes MASTER, CORE, TASK.
- `node_ids` (Set of String) Node ID list. Required when EffectiveType=NODE_ID.
- `node_names` (Set of String) Node name list. Required when EffectiveType=NODE_NAME.



<a id="nestedatt--application_extras--application_configs"></a>
### Nested Schema for `application_extras.application_configs`

Optional:

- `component_instance_name` (String) Component instance name.
- `component_name` (String) Component name.
- `config_file_name` (String) Configuration file name.
- `config_item_key` (String) Configuration item name.
- `config_item_value` (String) Configuration item value.
- `deleted` (Boolean) Whether to delete.
- `effective_scope` (Attributes) Affected component. (see [below for nested schema](#nestedatt--application_extras--application_configs--effective_scope))

<a id="nestedatt--application_extras--application_configs--effective_scope"></a>
### Nested Schema for `application_extras.application_configs.effective_scope`

Optional:

- `component_names` (Set of String) Component name list. Required when EffectiveType=COMPONENT_NAME.
- `effective_type` (String) Effective type. CLUSTER, NODE_GROUP_NAME, NODE_GROUP_ID, NODE_GROUP_TYPE, NODE_NAME, NODE_ID, COMPONENT_NAME.
- `node_group_ids` (Set of String) Node group ID list. Required when EffectiveType=NODE_GROUP_ID.
- `node_group_names` (Set of String) Node group name list. Required when EffectiveType=NODE_GROUP_NAME.
- `node_group_types` (Set of String) Node group type list. Required when EffectiveType=NODE_GROUP_TYPE. Currently includes MASTER, CORE, TASK.
- `node_ids` (Set of String) Node ID list. Required when EffectiveType=NODE_ID.
- `node_names` (Set of String) Node name list. Required when EffectiveType=NODE_NAME.




<a id="nestedatt--bootstrap_scripts"></a>
### Nested Schema for `bootstrap_scripts`

Optional:

- `effective_scope` (Attributes) Script execution scope. (see [below for nested schema](#nestedatt--bootstrap_scripts--effective_scope))
- `execution_fail_strategy` (String) Execution failure policy. Range: FAILED_CONTINUE: Continue with other tasks after failure. FAILED_BLOCK: Stop and do not execute subsequent tasks after failure. If the script is BOOTSTRAP, this will interrupt and fail cluster creation or node group expansion. Default: FAILED_BLOCK.
- `execution_moment` (String) Script execution timing. Only effective when scriptType=BOOTSTRAP. BEFORE_APPLICATION_INSTALL: before application installation. AFTER_APPLICATION_STARTED: after application startup. Default: BEFORE_APP_INSTALL.
- `priority` (String) Script execution priority. Range: 1~1000. Default: 1.
- `script_args` (String) Script parameters.
- `script_name` (String) Script name (required). Length: 1–128 characters. Must start with an uppercase or lowercase letter or Chinese character. Cannot start with http: or https:. Can include Chinese, English, numbers, underscores (_), or hyphens (-).
- `script_path` (String) Script TOS path. Required. Must start with 'tos:'.
- `script_type` (String) Script type. NORMAL: normal script. BOOTSTRAP: bootstrap script.

<a id="nestedatt--bootstrap_scripts--effective_scope"></a>
### Nested Schema for `bootstrap_scripts.effective_scope`

Optional:

- `component_names` (Set of String) Component name list. Required when EffectiveType=COMPONENT_NAME.
- `effective_type` (String) Effective type. CLUSTER, NODE_GROUP_NAME, NODE_GROUP_ID, NODE_GROUP_TYPE, NODE_NAME, NODE_ID, COMPONENT_NAME.
- `node_group_ids` (Set of String) Node group ID list. Required when EffectiveType=NODE_GROUP_ID.
- `node_group_names` (Set of String) Node group name list. Required when EffectiveType=NODE_GROUP_NAME.
- `node_group_types` (Set of String) Node group type list. Required when EffectiveType=NODE_GROUP_TYPE. Currently includes MASTER, CORE, TASK.
- `node_ids` (Set of String) Node ID list. Required when EffectiveType=NODE_ID.
- `node_names` (Set of String) Node name list. Required when EffectiveType=NODE_NAME.



<a id="nestedatt--charge_pre_config"></a>
### Nested Schema for `charge_pre_config`

Optional:

- `charge_period` (Number) When chargeType=PRE, default value=1. Unit for monthly subscription duration.
- `charge_period_unit` (String) When chargeType=PRE, default value=Month. Unit for monthly subscription duration. Range: Month: month. Year: year.

Read-Only:

- `auto_renew` (Boolean) Enable auto-renewal. Options: true (enabled), false (disabled).
- `auto_renew_period` (Number) Renewal duration when auto-renewal is triggered. When AutoRenew=true, the default value is 1.
- `auto_renew_period_unit` (String) Renewal duration unit when auto-renewal is triggered. When AutoRenew=true, default is Month. Options: Month, Year.
- `charge_type` (String) Payment type. Options: PRE, POST.


<a id="nestedatt--node_group_attributes"></a>
### Nested Schema for `node_group_attributes`

Optional:

- `bandwidth` (Number) Public bandwidth. Default is 8M. If the user changes it later, EMR needs to synchronize this information.
- `charge_type` (String) Node group billing type. If empty, reuses the cluster's chargeType. Master and Core groups must reuse the cluster-level billing type. When the cluster's chargeType is PRE, the task node group's chargeType can be set to POST. When the cluster's chargeType is POST, the node group's chargeType defaults to POST, and this parameter setting on the node group is invalid.
- `data_disks` (Attributes Set) Data disk configuration. Only disks of the same specification are supported within a single node group. The list length is limited to 1.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--node_group_attributes--data_disks))
- `ecs_instance_types` (Set of String) Node group's ECS instance type list. Only one instance type can be set. The list length is limited to 1.
- `ecs_key_pair_name` (String) ECS key pair name.
- `ecs_password` (String) Password for ECS root account.
- `node_count` (Number) Current expected number of nodes to purchase for the node group.
- `node_group_name` (String) Length: 1–128 characters. Cannot start with http: or https:. Can include Chinese, English, numbers, underscores (_), or hyphens (-).
- `node_group_type` (String) Node group type.
- `subnet_ids` (Set of String) Subnet ID list. Currently, only one parameter can be passed, and all node groups must use the same subnet ID.
- `system_disk` (Attributes) System disk configuration. (see [below for nested schema](#nestedatt--node_group_attributes--system_disk))
- `with_public_ip` (Boolean) Whether to attach public IP.
- `zone_id` (String) Availability zone ID.

<a id="nestedatt--node_group_attributes--data_disks"></a>
### Nested Schema for `node_group_attributes.data_disks`

Optional:

- `count` (Number) Number of disk blocks. Default is 4, maximum is 15, minimum is 1.
- `size` (Number) Disk size. Default is 80GB, minimum is 60GB, maximum is 2048GB, unit: GB.
- `volume_type` (String) Disk type. ESSD_PL0: Ultra SSD_PL0. ESSD_PL1: Ultra SSD_PL1. ESSD_PL2: Ultra SSD_PL2. ESSD_PL3: Ultra SSD_PL3. ESSD_FLEXPL: Ultra SSD_FlexPL. ULTRA_DISK: Efficient cloud disk. PTSSD: Performance SSD. SSD: General SSD. EHDD: Efficient cloud disk. ZENYA_SSD: Zenya. LOCAL_HDD: Big data HDD. LOCAL_SSD: Local SSD. LOCAL_SSD_SRIOV: Local SSD SRIOV.


<a id="nestedatt--node_group_attributes--system_disk"></a>
### Nested Schema for `node_group_attributes.system_disk`

Optional:

- `size` (Number) Disk size.
- `volume_type` (String) Disk type.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) User tag key.
- `value` (String) User tag value.


<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `application_config_home` (String) Application configuration path.
- `application_home` (String) Application installation path.
- `application_name` (String) Application name.
- `application_state` (String) Service status. NORMAL: normal; WARNING: alert; STOPPED: stopped; INIT: initializing; INSTALLING: installing; INSTALLED: installed; STARTING: starting; STARTED: started; STOPPING: stopping; UNINSTALLING: uninstalling; UNINSTALLED: uninstalled; EXCEPTION: exception.
- `application_version` (String) Application version.
- `group` (String) Application user group.
- `support_client` (Boolean) Whether client is supported.
- `user` (String) Application user.


<a id="nestedatt--state_change_reason"></a>
### Nested Schema for `state_change_reason`

Read-Only:

- `code` (String) Status update code.
- `reason` (String) Status update reason.

## Import

Import is supported using the following syntax:

```shell
$ terraform import volcenginecc_emr_cluster.example "cluster_id"
```
