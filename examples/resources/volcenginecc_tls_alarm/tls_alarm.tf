resource "volcenginecc_tls_alarm" "TLSAlarmDemo" {
  alarm_name = "ccapi-test"
  project_id = "project_id_xxxx"
  status     = false
  query_requests = [
    {
      end_time_offset        = -1
      end_time_offset_unit   = ""
      query                  = "raw:\"error\" | SELECT __time__, hostname, raw ORDER BY __time__ DESC LIMIT 5"
      start_time_offset      = -11
      start_time_offset_unit = ""
      time_span_type         = ""
      topic_id               = "topic_id_xxxx"
      topic_name             = "topic_name_xxxx"
      truncated_time         = ""
    },
    {
      end_time_offset        = -5
      end_time_offset_unit   = ""
      query                  = "test"
      start_time_offset      = -15
      start_time_offset_unit = ""
      time_span_type         = ""
      topic_id               = "topic_id_xxxx"
      topic_name             = "topic_name_xxxx"
      truncated_time         = ""
    },
    {
      end_time_offset        = -60
      end_time_offset_unit   = ""
      query                  = "demo"
      start_time_offset      = -120
      start_time_offset_unit = ""
      time_span_type         = ""
      topic_id               = "topic_id_xxxx"
      topic_name             = "topic_name_xxxx"
      truncated_time         = ""
    }
  ]
  request_cycle = {
    cron_tab       = "0 0/1 * * *"
    time           = 0
    type           = "CronTab"
    cron_time_zone = "Asia/Shanghai"
  }
  trigger_period = 10
  alarm_period   = 1440
  alarm_period_detail = {
    email           = 1440
    phone           = 1440
    sms             = 1440
    general_webhook = 1440
  }
  alarm_notify_groups = [
    {
      alarm_notify_group_id = "notify_group_id_xxxx"
    },
    {
      alarm_notify_group_id = "notify_group_id_xxxx"
    }
  ]
  user_define_msg = "{{toJson(QueryResult)}}{{UserName}}{{ProjectName}}{{Region}}{{Alarm}}{{AlarmID}}{{Topics|join:\",\"}}{{toJson(Results[0].RawResultsCount)}}"
  join_configurations = [
    {
      set_operation_type = "CrossJoin"
      condition          = ""
    },
    {
      set_operation_type = "InnerJoin"
      condition          = "$1. >= $3."
    }
  ]
  send_resolved = false
  trigger_conditions = [
    {
      condition       = ""
      count_condition = "__count__ == 2"
      no_data         = false
      severity        = "warning"
    },
    {
      condition       = "$1."
      count_condition = ""
      no_data         = false
      severity        = "critical"
    },
    {
      condition       = "$2."
      count_condition = "__count__ == 4"
      no_data         = false
      severity        = "warning"
    },
    {
      condition       = ""
      count_condition = ""
      no_data         = true
      severity        = "critical"
    }
  ]
}