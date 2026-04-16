resource "volcenginecc_tls_consumer_group" "TLSConsumerGroupDemo" {
  project_id = "c6fef4c1-041f-43*****"
  topic_id_list = [
    "bead2d9c*****"
  ]
  consumer_group_name = "test-gruopname"
  heartbeat_ttl       = 10
  ordered_consume     = true
}