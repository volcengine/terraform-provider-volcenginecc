resource "volcenginecc_kafka_group" "KafkaGroupDemo" {
  instance_id = "instance_id_xxxxxx"
  group_id    = "ccapi-test-2"
  description = "test"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}