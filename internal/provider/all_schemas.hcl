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

resource_schema "volcengine_iam_role" {
  cloudcontrol_type_name = "Volcengine::IAM::Role"
}

resource_schema "volcengine_ecs_image" {
  cloudcontrol_type_name = "Volcengine::ECS::Image"
}

resource_schema "volcengine_cr_repository" {
  cloudcontrol_type_name = "Volcengine::CR::Repository"
}

resource_schema "volcengine_vefaas_function" {
  cloudcontrol_type_name = "Volcengine::VEFAAS::Function"
}

resource_schema "volcengine_clb_clb" {
  cloudcontrol_type_name = "Volcengine::CLB::CLB"
}

resource_schema "volcengine_vpc_vpc" {
  cloudcontrol_type_name = "Volcengine::VPC::VPC"
}

resource_schema "volcengine_transitrouter_transitrouter" {
  cloudcontrol_type_name = "Volcengine::TransitRouter::TransitRouter"
}

resource_schema "volcengine_alb_loadbalancer" {
  cloudcontrol_type_name = "Volcengine::ALB::LoadBalancer"
}

resource_schema "volcengine_ebs_volume" {
  cloudcontrol_type_name = "Volcengine::EBS::Volume"
}

resource_schema "volcengine_iam_user" {
  cloudcontrol_type_name = "Volcengine::Iam::User"
}

resource_schema "volcengine_ark_endpoint" {
  cloudcontrol_type_name = "Volcengine::ARK::Endpoint"
}

resource_schema "volcengine_tos_bucket" {
  cloudcontrol_type_name = "Volcengine::TOS::Bucket"
}
resource_schema "volcengine_tos_bucket" {
  cloudcontrol_type_name = "Volcengine::TOS::Bucket"
}
