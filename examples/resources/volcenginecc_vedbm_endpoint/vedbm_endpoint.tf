resource "volcenginecc_vedbm_endpoint" "VEDBMEndpointDemo" {
  public_addresses = {
    eip_id = "eip-ij2xxxxx74o8cubosgei"
  }
  auto_add_new_nodes          = true
  consist_level               = "Eventual"
  description                 = "this is a test"
  distributed_transaction     = false
  endpoint_name               = "ccapi-test-1001"
  endpoint_type               = "Custom"
  instance_id                 = "vedbm-jxxxxttjdcea"
  master_accept_read_requests = false
  node_ids                    = ["vedbm-jxxxxttjdcea-1"]
  read_write_mode             = "ReadOnly"
}