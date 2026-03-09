resource "volcenginecc_kafka_allowlist" "KafkaAllowListDemo" {
  allow_list      = "127.0.0.2"
  allow_list_desc = "test"
  allow_list_name = "ccapi-text"
  associated_instances = [
    {
    instance_id = "kafka-cnng9x9s***" },
    {
    instance_id = "kafka-cnnghf99**" }
  ]
  apply_instance_num = 2

}