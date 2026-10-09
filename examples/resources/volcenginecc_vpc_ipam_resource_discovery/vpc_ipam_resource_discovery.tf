resource "volcenginecc_vpc_ipam_resource_discovery" "example" {
  ipam_resource_discovery_name = "ccapi-ipam-rd-renamed"
  description                  = "ccapi ipam resource discovery updated description"
  operating_regions            = ["cn-beijing"]
  tags = [{
    key   = "ccapi-test-key-1"
    value = "ccapi-test-value-1"
    }, {
    key   = "ccapi-test-key-2"
    value = "ccapi-test-value-2"
  }]

}
