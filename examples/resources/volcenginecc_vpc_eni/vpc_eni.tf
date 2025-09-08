resource "volcenginecc_vpc_eni" "eniDemo" {
  network_interface_name = "eniDemo"
  subnet_id              = volcenginecc_vpc_subnet.TestTerraform_vpc_subnet.id
  security_group_ids     = [
    volcenginecc_vpc_securitygroup.TestTerraform_vpc_security_group.id
  ]
  device_id          = volcenginecc_ecs_instance.TestTerraform_ecs_instance.id
  project_name       = "default"
  primary_ip_address = {
    private_ip_address = "192.168.xxx.10"
  }
  private_ip_sets = [
    {
      private_ip_address = "192.168.xxx.102"
    },
    {
      private_ip_address    = "192.168.xxx.101"
      associated_elastic_ip = {
        allocation_id = "eip-17xf500wwdb7kv98gwdxxxxxx"
      }
    }
  ]
  i_pv_6_sets : ["2408:1000:ab69:fe68:4099:9e0c:5c3c:xxx"]
  tags = [
    {
      key   = "ENV"
      value = "test"
    }
  ]
}