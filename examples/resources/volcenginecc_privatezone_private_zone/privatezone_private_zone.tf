resource "volcenginecc_privatezone_private_zone" "PrivateZoneDemo" {
  zone_name = "new.example.com"
  vpcs = [{
    vpc_id = "vpc-btgxxxxx0b2tnxxxxx"
    region = "cn-beijing"
  }]
  project_name   = "default"
  line_mode      = 1
  recursion_mode = true
  tags = [{
    value = "test"
    key   = "env"
  }]
  remark = "测试域名"
}