resource "volcenginecc_directconnect_virtual_interface" "DirectConnectVirtualInterfaceDemo" {
  local_ipv_6_ip               = "2408:xxxx:cc:400::1/64"
  description                  = "生产环境虚拟接口"
  direct_connect_gateway_id    = "dcg-****"
  direct_connect_connection_id = "dcc-****"
  peer_ip                      = "192.168.100.2/30"
  enable_bfd_echo              = false
  enable_nqa                   = false
  bfd_detect_multiplier        = 3
  peer_ipv_6_ip                = "2408:xxxx:cc:400::2/64"
  route_type                   = "BGP"
  virtual_interface_name       = "prod-virtual-interface"
  local_ip                     = "192.168.100.1/30"
  enable_bfd                   = true
  bandwidth                    = 50
  vlan_id                      = 2111
  tags = [
    {
      value = "env"
      key   = "test"
    }
  ]
  bfd_detect_interval = 1000
}