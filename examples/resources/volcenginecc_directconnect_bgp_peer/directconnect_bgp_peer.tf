resource "volcenginecc_directconnect_bgp_peer" "DirectConnectBgpPeerDemo" {
  auth_key             = "QWExxxxx"
  bgp_peer_name        = "ccapi-test"
  description          = "test"
  ip_version           = "IPv4"
  remote_asn           = 4294960000
  virtual_interface_id = "dcv-33cjs5xxxxxxxxx"
}