resource "volcenginecc_transitrouter_peer_attachment" "TransitRouterPeerAttachmentDemo" {
  transit_router_id                   = "tr-mjcxxxx"
  transit_router_attachment_name      = "ccapi-test-tf"
  description                         = "tf-test"
  peer_transit_router_id              = "tr-xxx"
  peer_transit_router_region_id       = "cn-xx"
  transit_router_bandwidth_package_id = "tbp-13f34rxxxx"
  bandwidth                           = 2
  tags {
    key   = "env"
    value = "test"
  }
}