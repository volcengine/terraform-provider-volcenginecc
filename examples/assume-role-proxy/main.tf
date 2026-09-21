terraform {
  required_providers {
    volcenginecc = {
      source = "volcengine/volcenginecc"
    }
  }
}

provider "volcenginecc" {
  region = "cn-beijing"
  endpoints = {
    cloudcontrolapi = "cloudcontrol.cn-beijing.volcengineapi.com"
  }
  assume_role = {
    assume_role_trn = var.role_trn
    proxy_url       = var.assume_role_proxy_url
  }
}

resource "volcenginecc_vpc_vpc" "regression" {
  cidr_block   = "192.168.0.0/24"
  vpc_name     = var.vpc_name
  description  = "Temporary AssumeRole proxy regression VPC"
  project_name = "default"
  tags         = [{ key = "purpose", value = "assume-role-proxy-regression" }]
}

output "vpc_id" {
  value = volcenginecc_vpc_vpc.regression.id
}
