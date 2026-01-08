resource "volcenginecc_bmq_group" "BMQGroupDemo" {
  description = "this is test group"
  group_name  = "cBMQGroupDemo"
  instance_id = "bmq-4ld4vpjzd32tq1gxxxxx"
  reset_info = {
    topic_id     = "5f81fcab96cb46c7955659fdxxxxx"
    reset_by     = "OFFSET"
    offset_type  = "LATEST"
    reset_value  = 4
    partition_id = 1
  }
}