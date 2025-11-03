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

resource_schema "volcengine_transitrouter_transit_router_route_entry" {
  cloudcontrol_type_name = "Volcengine::TransitRouter::TransitRouterRouteEntry"
}

resource_schema "volcengine_cr_name_space" {
  cloudcontrol_type_name = "Volcengine::CR::NameSpace"
}

resource_schema "volcengine_clb_rule" {
  cloudcontrol_type_name = "Volcengine::CLB::Rule"
}

resource_schema "volcengine_clb_certificate" {
  cloudcontrol_type_name = "Volcengine::CLB::Certificate"
}

resource_schema "volcengine_clb_acl" {
  cloudcontrol_type_name = "Volcengine::CLB::ACL"
}

resource_schema "volcengine_autoscaling_scaling_configuration" {
  cloudcontrol_type_name = "Volcengine::AutoScaling::ScalingConfiguration"
}

resource_schema "volcengine_alb_listener" {
  cloudcontrol_type_name = "Volcengine::ALB::Listener"
}

resource_schema "volcengine_vpc_prefix_list" {
  cloudcontrol_type_name = "Volcengine::VPC::PrefixList"
}

resource_schema "volcengine_alb_acl" {
  cloudcontrol_type_name = "Volcengine::ALB::ACL"
}

resource_schema "volcengine_alb_server_group" {
  cloudcontrol_type_name = "Volcengine::ALB::ServerGroup"
}

resource_schema "volcengine_alb_load_balancer" {
  cloudcontrol_type_name = "Volcengine::ALB::LoadBalancer"
}

resource_schema "volcengine_rdsmysql_database" {
  cloudcontrol_type_name = "Volcengine::RDSMySQL::Database"
}

resource_schema "volcengine_redis_instance" {
  cloudcontrol_type_name = "Volcengine::Redis::Instance"
}


resource_schema "volcengine_clb_listener" {
  cloudcontrol_type_name = "Volcengine::CLB::Listener"
}

resource_schema "volcengine_ecs_instance" {
  cloudcontrol_type_name = "Volcengine::ECS::Instance"
}


resource_schema "volcengine_waf_domain" {
  cloudcontrol_type_name = "Volcengine::WAF::Domain"
}

resource_schema "volcengine_cr_registry" {
  cloudcontrol_type_name = "Volcengine::CR::Registry"
}

resource_schema "volcengine_directconnect_direct_connect_gateway" {
  cloudcontrol_type_name = "Volcengine::DirectConnect::DirectConnectGateway"
}

resource_schema "volcengine_transitrouter_transit_router_route_table" {
  cloudcontrol_type_name = "Volcengine::TransitRouter::TransitRouterRouteTable"
}

resource_schema "volcengine_vpc_network_acl" {
  cloudcontrol_type_name = "Volcengine::VPC::NetworkAcl"
}

resource_schema "volcengine_tls_topic" {
  cloudcontrol_type_name = "Volcengine::TLS::Topic"
}

resource_schema "volcengine_apig_upstream_source" {
  cloudcontrol_type_name = "Volcengine::APIG::UpstreamSource"
}

resource_schema "volcengine_vke_addon" {
  cloudcontrol_type_name = "Volcengine::VKE::Addon"
}

resource_schema "volcengine_apig_gateway_service" {
  cloudcontrol_type_name = "Volcengine::APIG::GatewayService"
}

resource_schema "volcengine_vefaas_sandbox" {
  cloudcontrol_type_name = "Volcengine::VEFAAS::Sandbox"
}

resource_schema "volcengine_vefaas_kafka_trigger" {
  cloudcontrol_type_name = "Volcengine::VEFAAS::KafkaTrigger"
}




