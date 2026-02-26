resource "volcenginecc_vpc_traffic_mirror_target" "VPCTrafficMirrorTargetDemo" {
  traffic_mirror_target_name = "test-terraformtest"
  description                = "this is a test"
  instance_id                = "eni-2yggxh4o692bk6asvxxxxx"
  instance_type              = "NetworkInterface"
  project_name               = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}