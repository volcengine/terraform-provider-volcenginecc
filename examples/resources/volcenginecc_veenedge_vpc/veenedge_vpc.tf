resource "volcenginecc_veenedge_vpc" "Example" {
  desc    = "测试VPC实例"
  project = "default"
  subnets = [{
    desc = "测试子网1"
    cidr = "10.0.0.0/16"
    name = "test-subnet-1"
    }, {
    desc = "测试子网2"
    cidr = "10.1.0.0/16"
    name = "test-subnet-2"
  }]
  cluster_name = "bdcdn-sxcu"
  vpc_name     = "ccapi-dx-2"
  segment      = "10.0.0.0/8"
  tags = [{
    key   = "env"
    value = "test"
  }]
}