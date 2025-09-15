# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

defaults {
  schema_cache_directory     = "../service/cloudcontrol/schemas"
  terraform_type_name_prefix = "volcenginecc"
}

meta_schema {
  path = "../service/cloudcontrol/meta-schemas/provider.definition.schema.v1.json"
}

# CloudControl resource types schemas are available for use with the Cloud Control API.
resource_schema "volcengine_ecs_image" {
  cloudcontrol_type_name = "Volcengine::ECS::Image"
}

resource_schema "volcengine_iam_role" {
  cloudcontrol_type_name = "Volcengine::IAM::Role"

}

resource_schema "volcengine_iam_user" {
  cloudcontrol_type_name = "Volcengine::IAM::User"
}

resource_schema "volcengine_tos_bucket" {
  cloudcontrol_type_name = "Volcengine::TOS::Bucket"
}

resource_schema "volcengine_vpc_vpc" {
  cloudcontrol_type_name = "Volcengine::VPC::VPC"

}

resource_schema "volcengine_iam_accesskey" {
  cloudcontrol_type_name = "Volcengine::IAM::Accesskey"
}

resource_schema "volcengine_iam_group" {
  cloudcontrol_type_name = "Volcengine::IAM::Group"
}

resource_schema "volcengine_iam_policy" {
  cloudcontrol_type_name = "Volcengine::IAM::Policy"
}

resource_schema "volcengine_vpc_eip" {
  cloudcontrol_type_name = "Volcengine::VPC::EIP"
}

resource_schema "volcengine_vpc_eni" {
  cloudcontrol_type_name = "Volcengine::VPC::ENI"
}

resource_schema "volcengine_vpc_route_table" {
  cloudcontrol_type_name = "Volcengine::VPC::RouteTable"
}

resource_schema "volcengine_vpc_security_group" {
  cloudcontrol_type_name = "Volcengine::VPC::SecurityGroup"
}

resource_schema "volcengine_vpc_subnet" {
  cloudcontrol_type_name = "Volcengine::VPC::Subnet"
}

resource_schema "volcengine_ecs_command" {
  cloudcontrol_type_name = "Volcengine::ECS::Command"
}

resource_schema "volcengine_ecs_invocation" {
  cloudcontrol_type_name = "Volcengine::ECS::Invocation"
}

resource_schema "volcengine_vmp_workspace" {
  cloudcontrol_type_name = "Volcengine::VMP::Workspace"
}

resource_schema "volcengine_transitrouter_transit_router" {
  cloudcontrol_type_name = "Volcengine::TransitRouter::TransitRouter"
}

resource_schema "volcengine_clb_nlb" {
  cloudcontrol_type_name = "Volcengine::CLB::NLB"
}

resource_schema "volcengine_clb_nlb_server_group" {
  cloudcontrol_type_name = "Volcengine::CLB::NLBServerGroup"
}

resource_schema "volcengine_clb_clb" {
  cloudcontrol_type_name = "Volcengine::CLB::CLB"
}

resource_schema "volcengine_cr_repository" {
  cloudcontrol_type_name = "Volcengine::CR::Repository"
}

resource_schema "volcengine_filenas_instance" {
  cloudcontrol_type_name = "Volcengine::FileNAS::Instance"
}

resource_schema "volcengine_storageebs_volume" {
  cloudcontrol_type_name = "Volcengine::StorageEBS::Volume"
}

resource_schema "volcengine_ark_endpoint" {
  cloudcontrol_type_name = "Volcengine::ARK::Endpoint"
}

# 9月15日发布
resource_schema "volcengine_ecs_keypair" {
  cloudcontrol_type_name = "Volcengine::ECS::Keypair"
}

resource_schema "volcengine_ecs_hpc_cluster" {
  cloudcontrol_type_name = "Volcengine::ECS::HpcCluster"
}

resource_schema "volcengine_clb_nlb_listener" {
  cloudcontrol_type_name = "Volcengine::CLB::NLBListener"
}

resource_schema "volcengine_clb_server_group" {
  cloudcontrol_type_name = "Volcengine::CLB::ServerGroup"
}

resource_schema "volcengine_vpc_bandwidth_package" {
  cloudcontrol_type_name = "Volcengine::VPC::BandwidthPackage"
}

resource_schema "volcengine_alb_certificate" {
  cloudcontrol_type_name = "Volcengine::ALB::Certificate"
}