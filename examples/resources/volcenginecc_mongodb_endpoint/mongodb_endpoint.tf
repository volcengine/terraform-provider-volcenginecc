resource "volcenginecc_mongodb_endpoint" "MongoDBEndpointDemo" {
  eip_ids      = ["eip-****", "eip-****"]
  instance_id  = "mongo-replica-****"
  network_type = "Public"
}