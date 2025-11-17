resource "volcenginecc_privatezone_resolver_rule" "PrivateZoneResolverRuleDemo" {
  line = "电信"
  name = "PrivateZoneResolverRuleDemo"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  type = "LINE"
  vp_cs = [
    {
      region = "cn-beijing"
    vpc_id = "vpc-3nqt4kq87xn28931eclihh90****" },
    {
      region = "cn-beijing"
    vpc_id = "vpc-1a10aeq5vr2tc8nvepkauwljx****" }
  ]
  vpc_trns = ["trn:vpc:cn-beijing:********:vpc/vpc-3nqt4kq87xn2893xxxxx", "trn:vpc:cn-beijing:********:vpc/vpc-1a10aeq5vr2tc8nvepxxxxx"]
}