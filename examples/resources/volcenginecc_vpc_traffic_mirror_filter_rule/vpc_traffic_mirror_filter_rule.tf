resource "volcenginecc_vpc_traffic_mirror_filter_rule" "VPCTrafficMirrorFilterRuleDemo" {
  description              = "test"
  destination_cidr_block   = "10.*.*.0/24"
  destination_port_range   = "80/80"
  policy                   = "accept"
  priority                 = 1
  protocol                 = "tcp"
  source_cidr_block        = "10.*.*.0/24"
  source_port_range        = "80/80"
  traffic_direction        = "ingress"
  traffic_mirror_filter_id = "tmf-3nqp5bt6a3dog931exxxxx"
}

