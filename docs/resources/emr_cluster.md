---
page_title: "volcenginecc_emr_cluster Resource - terraform-provider-volcenginecc"
subcategory: "EMR"
description: |-
  E-MapReduce（EMR）是开源 Hadoop 生态的企业级大数据分析系统，完全兼容开源，为您提供 Hadoop、Spark、Hive、Hudi、Iceberg 等生态组件集成和管理。
---

# volcenginecc_emr_cluster (Resource)

E-MapReduce（EMR）是开源 Hadoop 生态的企业级大数据分析系统，完全兼容开源，为您提供 Hadoop、Spark、Hive、Hudi、Iceberg 等生态组件集成和管理。

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

- `charge_type` (String) 付费类型，PRE表示包月，POST表示按量计费。
- `cluster_name` (String) 集群名称。
- `cluster_type` (String) 集群类型。
- `node_attribute` (Attributes) 集群全局的节点信息。 (see [below for nested schema](#nestedatt--node_attribute))
- `release_version` (String) 集群版本。
- `security_group_id` (String) 集群全局安全组ID，所有节点组下的ecs都会加入该安全组。
- `vpc_id` (String) Vpc ID。

### Optional

- `application_extras` (Attributes List) 集群服务的扩展信息列表，包括服务的自定义配置项、服务组件的自定义部署拓扑设置，以及服务的元数据连接配置信息。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--application_extras))
- `application_names` (Set of String) 集群安装的服务名称列表。创建字段。
- `bootstrap_scripts` (Attributes Set) 集群的引导脚本列表。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--bootstrap_scripts))
- `charge_pre_config` (Attributes) 包月的配置参数，当chargeType=PRE时，必选。 (see [below for nested schema](#nestedatt--charge_pre_config))
- `deploy_mode` (String) 部署模式。SIMPLE表示简单模式，HIGH_AVAILABLE表示高可用模式。
- `history_server_mode` (String) HistoryServer模式，LOCAL将活动数据存储于集群内，PHS将活动数据存储于集群外。
- `node_group_attributes` (Attributes Set) 节点组属性列表。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--node_group_attributes))
- `project_name` (String) 资源所属项目，默认为default。一个资源只能归属于一个项目。只能包含字母、数字、下划线“_”、点“.”和中划线“-”。长度限制在64个字符以内。
- `security_mode` (String) 安全模式。
- `tags` (Attributes Set) 标签列表。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `account_id` (Number) 账号ID。
- `applications` (Attributes Set) 集群安装的服务列表。只读字段。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--applications))
- `cluster_domain_names` (Set of String) 集群dns域名后缀列表。
- `cluster_id` (String) 集群ID。
- `cluster_state` (String) 集群状态。
- `created_time` (Number) 集群创建时间。
- `creator_id` (Number) 创建者ID。
- `creator_name` (String) 创建者名称。
- `ecs_image_id` (String) ECS镜像ID。
- `expire_time` (Number) 集群过期时间。
- `id` (String) Uniquely identifies the resource.
- `ready_time` (Number) 集群创建完成时间。
- `state_change_reason` (Attributes) 状态变更原因。 (see [below for nested schema](#nestedatt--state_change_reason))
- `terminate_time` (Number) 集群终止时间。

<a id="nestedatt--node_attribute"></a>
### Nested Schema for `node_attribute`

Optional:

- `ecs_iam_role` (String) ECS实例角色。
- `zone_id` (String) 可用区ID。


<a id="nestedatt--application_extras"></a>
### Nested Schema for `application_extras`

Optional:

- `application_component_layouts` (Attributes Set) 服务组件的自定义部署拓扑列表。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--application_extras--application_component_layouts))
- `application_configs` (Attributes List) 服务的自定义配置参数列表。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--application_extras--application_configs))
- `application_name` (String) 应用名称。
- `connection_id` (String) 元数据连接id。
- `connection_type` (String) 元数据连接类型。BUILT_IN_MYSQL：内置数据库。EXTERNAL_MYSQL：外置数据库。HIVE_METASTORE：HMS。

<a id="nestedatt--application_extras--application_component_layouts"></a>
### Nested Schema for `application_extras.application_component_layouts`

Optional:

- `component_name` (String) 组件名称。
- `effective_scope` (Attributes) 组件的布局范围。 (see [below for nested schema](#nestedatt--application_extras--application_component_layouts--effective_scope))

<a id="nestedatt--application_extras--application_component_layouts--effective_scope"></a>
### Nested Schema for `application_extras.application_component_layouts.effective_scope`

Optional:

- `component_names` (Set of String) 组件名列表，当EffectiveType=COMPONENT_NAME，必选。
- `effective_type` (String) 生效类型。CLUSTER，NODE_GROUP_NAME，NODE_GROUP_ID，NODE_GROUP_TYPE，NODE_NAME，NODE_ID，COMPONENT_NAME。
- `node_group_ids` (Set of String) 节点组ID列表，EffectiveType=NODE_GROUP_ID时，必选。
- `node_group_names` (Set of String) 节点组名称列表，EffectiveType=NODE_GROUP_NAME时，必选。
- `node_group_types` (Set of String) 节点组类型列表，EffectiveType=NODE_GROUP_TYPE时，必选。目前包括MASTER、CORE、TASK。
- `node_ids` (Set of String) 节点ID列表，EffectiveType=NODE_ID时，必选。
- `node_names` (Set of String) 节点名列表，EffectiveType=NODE_NAME时，必选。



<a id="nestedatt--application_extras--application_configs"></a>
### Nested Schema for `application_extras.application_configs`

Optional:

- `component_instance_name` (String) 组件实例名称。
- `component_name` (String) 组件名称。
- `config_file_name` (String) 配置文件名。
- `config_item_key` (String) 配置项名称。
- `config_item_value` (String) 配置项值。
- `deleted` (Boolean) 是否删除。
- `effective_scope` (Attributes) 影响组件。 (see [below for nested schema](#nestedatt--application_extras--application_configs--effective_scope))

<a id="nestedatt--application_extras--application_configs--effective_scope"></a>
### Nested Schema for `application_extras.application_configs.effective_scope`

Optional:

- `component_names` (Set of String) 组件名列表，当EffectiveType=COMPONENT_NAME，必选。
- `effective_type` (String) 生效类型。CLUSTER，NODE_GROUP_NAME，NODE_GROUP_ID，NODE_GROUP_TYPE，NODE_NAME，NODE_ID，COMPONENT_NAME。
- `node_group_ids` (Set of String) 节点组ID列表，EffectiveType=NODE_GROUP_ID时，必选。
- `node_group_names` (Set of String) 节点组名称列表，EffectiveType=NODE_GROUP_NAME时，必选。
- `node_group_types` (Set of String) 节点组类型列表，EffectiveType=NODE_GROUP_TYPE时，必选。目前包括MASTER、CORE、TASK。
- `node_ids` (Set of String) 节点ID列表，EffectiveType=NODE_ID时，必选。
- `node_names` (Set of String) 节点名列表，EffectiveType=NODE_NAME时，必选。




<a id="nestedatt--bootstrap_scripts"></a>
### Nested Schema for `bootstrap_scripts`

Optional:

- `effective_scope` (Attributes) 脚本执行范围。 (see [below for nested schema](#nestedatt--bootstrap_scripts--effective_scope))
- `execution_fail_strategy` (String) 执行失败策略。取值范围：FAILED_CONTINUE：失败后继续执行其他任务。FAILED_BLOCK：失败后中断，不再继续执行后续任务。当脚本为BOOTSTRAP时，会导致创建集群、扩容节点组操作中断并失败。默认值：FAILED_BLOCK。
- `execution_moment` (String) 脚本的执行时机。仅scriptType=BOOTSTRAP时生效。BEFORE_APPLICATION_INSTALL：应用安装前。AFTER_APPLICATION_STARTED：应用启动后。默认值：BEFORE_APP_INSTALL
- `priority` (String) 脚本执行优先级。取值范围：1~1000。默认值1。
- `script_args` (String) 脚本参数。
- `script_name` (String) 脚本名称。必填。长度为1~128个字符，必须以大小字母或中文开头，不能以 http:： 和 https:： 开头。可以包含中文、英文、数字、下划线（_）、或者短划线（-）。
- `script_path` (String) 脚本所在TOS路径。必填。以 tos:： 开头。
- `script_type` (String) 脚本类型。NORMAL：普通脚本。BOOTSTRAP：引导脚本。

<a id="nestedatt--bootstrap_scripts--effective_scope"></a>
### Nested Schema for `bootstrap_scripts.effective_scope`

Optional:

- `component_names` (Set of String) 组件名列表，当EffectiveType=COMPONENT_NAME，必选。
- `effective_type` (String) 生效类型。CLUSTER，NODE_GROUP_NAME，NODE_GROUP_ID，NODE_GROUP_TYPE，NODE_NAME，NODE_ID，COMPONENT_NAME。
- `node_group_ids` (Set of String) 节点组ID列表，EffectiveType=NODE_GROUP_ID时，必选。
- `node_group_names` (Set of String) 节点组名称列表，EffectiveType=NODE_GROUP_NAME时，必选。
- `node_group_types` (Set of String) 节点组类型列表，EffectiveType=NODE_GROUP_TYPE时，必选。目前包括MASTER、CORE、TASK。
- `node_ids` (Set of String) 节点ID列表，EffectiveType=NODE_ID时，必选。
- `node_names` (Set of String) 节点名列表，EffectiveType=NODE_NAME时，必选。



<a id="nestedatt--charge_pre_config"></a>
### Nested Schema for `charge_pre_config`

Optional:

- `charge_period` (Number) chargeType=PRE默认值=1，包月的购买时长单位。
- `charge_period_unit` (String) chargeType=PRE时，默认值=Month，包月的购买时长单位，取值范围：Month：月。Year：年。

Read-Only:

- `auto_renew` (Boolean) 是否开启自动续费。取值范围：true：开启。false：不开启。
- `auto_renew_period` (Number) 自动续费触发时的续费时长，当AutoRenew=true时，默认值=1。
- `auto_renew_period_unit` (String) 自动续费触发时的续费时长单位，当AutoRenew=true时，默认值=Month。取值范围：Month：月。Year：年。
- `charge_type` (String) 付费类型，取值范围：PRE，POST。


<a id="nestedatt--node_group_attributes"></a>
### Nested Schema for `node_group_attributes`

Optional:

- `bandwidth` (Number) 公网带宽。默认值 8M。后续如果用户侧调整了，emr侧需要同步该信息。
- `charge_type` (String) 节点组付费类型。为空时，复用集群的chargeType。Master、Core组必须复用集群维度的付费类型。当集群的chargeType为PRE时，task节点组的chargeType允许设置为POST；当集群的chargeType为POST时，节点组的chargeType默认为POST，节点组上的此参数设置无效。
- `data_disks` (Attributes Set) 数据盘配置。当前单个节点组内只支持同规格的数据盘。即List的长度限制为1。
 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。 (see [below for nested schema](#nestedatt--node_group_attributes--data_disks))
- `ecs_instance_types` (Set of String) 节点组的ecs机型列表。当前只支持设置1个机型。即List的长度限制为1。
- `ecs_key_pair_name` (String) ecs的密钥对名称。
- `ecs_password` (String) Ecs root账号的密码。
- `node_count` (Number) 节点组当前期望购买的节点数量。
- `node_group_name` (String) 长度为1~128个字符，不能以 http:： 和 https:： 开头。可以包含中文、英文、数字、下划线（_）、或者短划线（-）。
- `node_group_type` (String) 节点组类型。
- `subnet_ids` (Set of String) 子网Id列表，目前只能传递一个参数，且各节点组的子网Id都是相同的。
- `system_disk` (Attributes) 系统盘配置。 (see [below for nested schema](#nestedatt--node_group_attributes--system_disk))
- `with_public_ip` (Boolean) 是否挂载公网ip。
- `zone_id` (String) 可用区ID。

<a id="nestedatt--node_group_attributes--data_disks"></a>
### Nested Schema for `node_group_attributes.data_disks`

Optional:

- `count` (Number) 磁盘块数，默认值4，最大15，最小1。
- `size` (Number) 磁盘大小，默认值80GB，最小60GB，最大2048GB，单位GB。
- `volume_type` (String) 磁盘类型。ESSD_PL0 ：极速型SSD_PL0。ESSD_PL1 ：极速型SSD_PL1。ESSD_PL2 ：极速型SSD_PL2。ESSD_PL3 ：极速型SSD_PL3。ESSD_FLEXPL ：极速型SSD_FlexPL。ULTRA_DISK ：高效云盘。PTSSD ：性能型SSD。SSD ：通用型SSD。EHDD ：高效云盘。ZENYA_SSD ：Zenya。LOCAL_HDD ：大数据型HDD。LOCAL_SSD ：本地SSD型。LOCAL_SSD_SRIOV ：本地SSD型SRIOV


<a id="nestedatt--node_group_attributes--system_disk"></a>
### Nested Schema for `node_group_attributes.system_disk`

Optional:

- `size` (Number) 磁盘大小。
- `volume_type` (String) 磁盘类型。



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) 用户标签的标签键。
- `value` (String) 用户标签的标签值。


<a id="nestedatt--applications"></a>
### Nested Schema for `applications`

Read-Only:

- `application_config_home` (String) 应用配置路径。
- `application_home` (String) 应用安装路径。
- `application_name` (String) 应用名称。
- `application_state` (String) 服务状态。NORMAL：正常；WARNING：告警；STOPPED：已停止；INIT：初始化中；INSTALLING：安装中；INSTALLED：已安装；STARTING：启动中；STARTED：已启动；STOPPING：停止中；UNINSTALLING：卸载中；UNINSTALLED：已卸载；EXCEPTION：异常。
- `application_version` (String) 应用版本。
- `group` (String) 应用用户组。
- `support_client` (Boolean) 是否支持客户端。
- `user` (String) 应用用户。


<a id="nestedatt--state_change_reason"></a>
### Nested Schema for `state_change_reason`

Read-Only:

- `code` (String) 状态更新码。
- `reason` (String) 状态更新原因。

## Import

Import is supported using the following syntax:

```shell
$ terraform import volcenginecc_emr_cluster.example "cluster_id"
```
