resource "volcenginecc_vpc_traffic_mirror_session" "VPCTrafficMirrorSessionDemo" {
  description          = "test-description"
  network_interface_id = "eni-1a1s0wz7xkm4g8nvexxxxx"
  packet_length        = 64
  priority             = 12
  project_name         = "default"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  traffic_mirror_filter_id    = "tmf-bu54sj0bl2bk5h0xxxxx"
  traffic_mirror_session_name = "test-1"
  traffic_mirror_target_id    = "tmt-ij32u0acvta874o8ctxxxxx"
  virtual_network_id          = 13
}