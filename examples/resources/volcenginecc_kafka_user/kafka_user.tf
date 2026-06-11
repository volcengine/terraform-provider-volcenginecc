resource "volcenginecc_kafka_user" "KafkaUserDemo" {
  instance_id   = "kafka-cnngavomxxxxxx"
  user_name     = "user1123"
  user_password = "********"
  all_authority = true
  description   = "vip"
}