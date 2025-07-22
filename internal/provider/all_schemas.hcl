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