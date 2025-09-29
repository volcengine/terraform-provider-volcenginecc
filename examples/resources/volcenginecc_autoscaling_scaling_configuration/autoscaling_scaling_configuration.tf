resource "volcenginecc_autoscaling_scaling_configuration" "AutoScalingDemo" {
  instance_name              = "AutoScalingDemo"
  scaling_configuration_name = "AutoScalingDemo"
  host_name                  = "AutoScalingDemo"
  scaling_group_id           = "scg-ye43d97gsvkxgganxxxx"
  security_group_ids         = ["sg-rrco3fkzzy0wv0x589wxxxx"]
  eip = {
    bandwidth            = 1
    billing_type         = "PostPaidByBandwidth"
    isp                  = "BGP"
    bandwidth_package_id = "bwp-btgl56gbli4g5h0b2v7ixxxx"
  }
  ipv_6_address_count = 1
  spot_strategy       = "SpotWithPriceLimit"
  instance_type_overrides = [
    {
      instance_type = "ecs.g4il.large"
    price_limit = 0.5 }
  ]
  image_id                      = "image-aagd56zrvqjtdripxxxx"
  security_enhancement_strategy = "Active"
  volumes = [
    {
      delete_with_instance = true
      size                 = 40
    volume_type = "ESSD_FlexPL" }
  ]
  project_name         = "default"
  key_pair_name        = "test"
  password             = "test"
  instance_description = "Web server configuration"
  zone_id              = "cn-test"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}