resource "volcenginecc_kafka_instance" "KafkaInstanceDemo" {
  charge_info = {
    charge_type = "PostPaid"
  }
  compute_spec         = "kafka.20xrate.hw"
  eip_id               = "eip-****"
  instance_description = "CCAPI-TEST"
  instance_name        = "CCAPI-TEST"
  subnet_id            = "subnet-****"
  ip_white_list        = ["acl-****"]
  parameters = jsonencode(
    {
      "MessageMaxByte" : "10",
      "LogRetentionHours" : "72",
      "OffsetRetentionMinutes" : "4320",
      "MessageTimestampType" : "CreateTime",
      "AutoDeleteGroup" : "true",
      "LogAutoDeletion" : "true",
      "LogAutoDeletionPercent" : "90",
      "TransactionDisabled" : "false"
    }
  )
  partition_number = 350
  storage_space    = 300
  storage_type     = "ESSD_PL0"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  version      = "2.8.2"
  vpc_id       = "vpc-****"
  zone_id      = "cn-beijing-a"
  project_name = "default"
}