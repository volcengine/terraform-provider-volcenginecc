resource "volcenginecc_rocketmq_access_key" "RocketMQAccessKeyDemo" {
  topic_permissions = [{
    permission = "SUB"
    topic_name = "topic1"
  }]
  instance_id   = "rocketmq-cnnxxxxxxxxxx"
  all_authority = "SUB"
  description   = "自定义SUB权限"
}