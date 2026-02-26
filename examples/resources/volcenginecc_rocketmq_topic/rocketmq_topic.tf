resource "volcenginecc_rocketmq_topic" "RocketMQTopicDemo" {
  instance_id  = "rocketmq-cnngxxxxx3c2d"
  topic_name   = "ccapi-test-test-2"
  message_type = 0
  description  = "this is a test"
  queue_number = 6
  access_policies = [
    {
      access_key = "TPl1xxxxxxblh5Gpxa79"
    authority = "ALL" },
  ]
}