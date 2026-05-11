resource "volcenginecc_tos_bucket_notification" "TOSBucketNotificationDemo" {
  bucket_name = "ccapi-test"
  notification_rules = [{
    destination = {
      kafka = [{
        instance_id = "kafka-cnnxxxxxm26rh"
        region      = "cn-beijing"
        role        = "trn:iam::21xxxxxxx:role/TOSNotiKafkaRole"
        topic       = "topic-1"
        user        = "user-1"
      }]
      rocket_mq = [{
        access_key_id = "IpGw4i6xxxxxKrkdwRXZ"
        instance_id   = "rocketmq-cnngaxxxxx22ab8"
        role          = "trn:iam::21xxxxxxx:role/TOSNotiRocketMQRole"
        topic         = "topic-1"
      }]
      ve_faa_s = [{
        function_id = "o1xxxxx"
      }]
    }
    events = ["tos:ObjectCreated:Put", "tos:LifecycleExpiration:Delete"]
    filter = {
      tos_key = {
        filter_rules = [{
          name  = "prefix"
          value = "preifx"
          }, {
          name  = "suffix"
          value = "suffix"
        }]
      }
    }
    rule_id = "test-0"
  }]
}