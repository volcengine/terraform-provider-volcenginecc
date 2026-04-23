resource "volcenginecc_tls_alarm_notify_group" "TLSAlarmNotifyGroupDemo" {
  alarm_notify_group_name = "ccapi-test"
  iam_project_name        = "default"
  notice_rules = [
    {
      rule_node = "{\"Type\":\"Operation\",\"Value\":[\"OR\"],\"Children\":[{\"Type\":\"Operation\",\"Value\":[\"AND\"],\"Children\":[{\"Type\":\"Operation\",\"Value\":[\"OR\"],\"Children\":[{\"Type\":\"Condition\",\"Value\":[\"Severity\",\"in\",\"[\\\"notice\\\"]\"]},{\"Type\":\"Condition\",\"Value\":[\"Duration\",\">=\",\"5\"]}]},{\"Type\":\"Condition\",\"Value\":[\"NotifyTime\",\"in\",\"[\\\"1776669310\\\",\\\"1777014911\\\"]\"]}]},{\"Type\":\"Condition\",\"Value\":[\"Severity\",\"in\",\"[\\\"warning\\\",\\\"critical\\\"]\"]}]}"
      receiver_infos = [
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxx", "xxxxxxx"]
          receiver_channels      = ["Email", "Sms", "Phone", "GeneralWebhook"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "aa6d01cd-5cf4-449d-a944-xxxxxxx"
          alarm_webhook_integration_name = "test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["Lark"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "abc47ee4-9ad3-4c25-803f-xxxxxxxxx"
          alarm_webhook_integration_name = "飞书test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "UserGroup"
          receiver_names         = ["xxxxxxxxx"]
          receiver_channels      = ["Email", "Phone", "GeneralWebhook"]
          start_time             = "00:30:00"
          end_time               = "23:30:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "aa6d01cd-5cf4-449d-a944-xxxxxxxxx"
          alarm_webhook_integration_name = "test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        }
      ]
      has_next = true
    has_end_node = true },
    {
      rule_node = "{\"Type\":\"Operation\",\"Value\":[\"AND\"],\"Children\":[{\"Type\":\"Operation\",\"Value\":[\"OR\"],\"Children\":[{\"Type\":\"Operation\",\"Value\":[\"AND\"],\"Children\":[{\"Type\":\"Condition\",\"Value\":[\"Severity\",\"in\",\"[\\\"notice\\\"]\"]},{\"Type\":\"Condition\",\"Value\":[\"Alarm\",\"in\",\"[\\\"test1\\\"]\"]}]},{\"Type\":\"Condition\",\"Value\":[\"AlarmID\",\"in\",\"[\\\"66fff5a0-af57-43d1-a40d-f2d32605aa61\\\"]\"]}]},{\"Type\":\"Operation\",\"Value\":[\"AND\"],\"Children\":[{\"Type\":\"Condition\",\"Value\":[\"Duration\",\"==\",\"5\"]},{\"Type\":\"Condition\",\"Value\":[\"NotifyTime\",\"in\",\"[\\\"1776653008\\\",\\\"1776998606\\\"]\"]}]}]}"
      receiver_infos = [
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["Email", "Sms", "Phone"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            },
            {
              key   = "env"
              value = "test"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "aa6d01cd-5cf4-449d-a944-xxxxxxxxx"
          alarm_webhook_integration_name = "test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx", "xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "UserGroup"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["Email", "Sms", "Phone"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "aa6d01cd-5cf4-449d-a944-xxxxxxxxx"
          alarm_webhook_integration_name = "test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx", "xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["Lark"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "abc47ee4-9ad3-4c25-803f-xxxxxxxxx"
          alarm_webhook_integration_name = "飞书test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx", "xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["DingTalk"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "3eaaef0a-134a-4d62-bf48-xxxxxxxxx"
          alarm_webhook_integration_name = "钉钉test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx", "xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        },
        {
          receiver_type          = "User"
          receiver_names         = ["xxxxxxxxx", "xxxxxxxxx"]
          receiver_channels      = ["WeChat"]
          start_time             = "00:00:00"
          end_time               = "23:59:59"
          general_webhook_url    = ""
          general_webhook_method = ""
          general_webhook_headers = [
            {
              key   = "Content-Type"
              value = "application/json"
            }
          ]
          general_webhook_body           = ""
          alarm_content_template_id      = "default-template"
          alarm_webhook_integration_id   = "c0454fd5-6597-4fc9-ad53-xxxxxxxxx"
          alarm_webhook_integration_name = "企业微信test"
          alarm_webhook_is_at_all        = false
          alarm_webhook_at_users         = ["xxxxxxxxx", "xxxxxxxxx"]
          alarm_webhook_at_groups        = ["xxxxxxxxx"]
        }
      ]
      has_next = false
    has_end_node = false }
  ]
}