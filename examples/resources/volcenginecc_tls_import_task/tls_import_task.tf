resource "volcenginecc_tls_import_task" "TLSImportTaskDemo" {
  description = "ccapi-test-kafka"
  import_source_info = {
    kafka_source_info = {
      host                = "kafka-cnngsl83xxxxx.kafka.cn-beijing.ivolces.com:9092"
      group               = "group1"
      topic               = "topic1"
      encode              = "UTF-8"
      password            = ""
      protocol            = ""
      username            = ""
      mechanism           = ""
      instance_id         = "kafka-cnngsl83e6xxxxx"
      is_need_auth        = false
      initial_offset      = 0
      time_source_default = 1
    }
  }
  project_id  = "4f1af9e7-34af-4ce5-866e-xxxxxxx"
  source_type = "kafka"
  target_info = {
    region     = "cn-beijing"
    log_type   = "json_log"
    log_sample = ""
    extract_rule = {
      extract_rule = {
        begin_regex = ""
        delimiter   = ""
        log_regex   = ""
        log_template = {
          format = ""
          type   = ""
        }
        quote                   = ""
        time_format             = "%Y-%m-%d %H:%M:%S,%f"
        time_key                = "time"
        time_sample             = ""
        un_match_log_key        = "LogParseFailed"
        un_match_up_load_switch = true
        filter_key_regex = [
          {
            key   = "user_id"
            regex = "^[0-9]+$"
          }
        ]
      }
      skip_line_count    = 0
      time_extract_regex = "[0-9]{0,2}\\/[0-9a-zA-Z]+\\/[0-9:,]+"
      time_zone          = "Asia/Shanghai"
    }
  }
  task_name = "ccapi-test-kafka-1001"
  topic_id  = "b75fffd8-1986-460c-9cca-xxxxxxx"
}