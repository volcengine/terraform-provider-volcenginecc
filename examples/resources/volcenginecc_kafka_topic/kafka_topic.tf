resource "volcenginecc_kafka_topic" "KafkaTopicDemo" {
  access_policies = [
    {
      user_name     = "KafkaTopicDemo-A"
      access_policy = "Pub"
    },
    {
      user_name     = "KafkaTopicDemo-B"
      access_policy = "Sub"
    },
    {
      user_name     = "KafkaTopicDemo-C"
      access_policy = "PubSub"
    }
  ]
  all_authority    = true
  cleanup_policy   = ["delete", "compact"]
  description      = "KafkaTopicDemo"
  instance_id      = "kafka-c****f"
  parameters       = "{\"LogRetentionHours\":\"3\",\"MessageMaxByte\":\"2\",\"MinInsyncReplicaNumber\":\"1\"}"
  partition_number = 1
  replica_number   = 1
  topic_name       = "KafkaTopicDemo"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
}