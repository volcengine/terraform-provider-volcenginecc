resource "volcenginecc_cen_inter_region_bandwidth" "CENInterRegionBandwidthDemo" {
  local_region_id          = "cn-beijing"
  peer_region_id           = "cn-shanghai"
  bandwidth                = 1
  cen_id                   = "cen-2v73nw1h8a03k6x7exxxxx"
  cen_bandwidth_package_id = "cbp-mikp555wa4u85smt1xxxxxx"
}