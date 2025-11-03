resource "volcenginecc_vefaas_kafka_trigger" "VefaasTriggerDemo" {
  function_id            = "djb27tnr****"
  name                   = "VefaasTriggerDemo"
  maximum_retry_attempts = 74
  starting_position      = "Earliest"
  enabled                = false
  kafka_credentials = {
    mechanism = "PLAIN"
    username  = "VefaasTriggerDemo"
    password  = "*****"
  }
  mq_instance_id                    = "kafka-cnng8cm4q2fncrvs****"
  topic_name                        = "VefaasTriggerDemo"
  description                       = "VefaasTriggerDemo description"
  batch_flush_duration_milliseconds = 8637
  batch_size                        = 709
}