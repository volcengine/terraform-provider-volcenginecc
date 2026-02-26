resource "volcenginecc_rocketmq_group" "RocketMQGroupDemo" {
  instance_id             = "rocketmq-cnnxxxxxc2d"
  group_id                = "ccapi-test-test-1"
  consume_message_orderly = true
  description             = "test"
  group_type              = "TCP"
  retry_max_times         = 2
}