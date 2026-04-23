---
page_title: "volcenginecc_tls_alarm_notify_group Resource - terraform-provider-volcenginecc"
subcategory: "TLS"
description: |-
  The logging service notification group is a collection of alarm notification channels used to specify the action strategy after an alarm is triggered. You can add notification rules, and the logging service will intelligently assign different alarm notification channels based on these rules.
---

# volcenginecc_tls_alarm_notify_group (Resource)

The logging service notification group is a collection of alarm notification channels used to specify the action strategy after an alarm is triggered. You can add notification rules, and the logging service will intelligently assign different alarm notification channels based on these rules.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alarm_notify_group_name` (String) Alarm notification group name. Refer to the resource naming rules.

### Optional

- `iam_project_name` (String) Name of the IAM project to which the alarm group belongs. If not specified, the logging service adds the alarm group to the IAM project named default.
- `notice_rules` (Attributes Set) Alarm notification group configuration. Note: If the NoticeRules parameter is configured, leave the NotifyType and Receivers parameters empty. If NoticeRules is empty, you must configure the NotifyType and Receivers parameters. When modifying, do not change NoticeRules, NotifyType, and Receivers at the same time, as some fields may not take effect.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--notice_rules))
- `notify_type` (Set of String) Type of alarm notification. Trigger: alarm triggered. Recovery: alarm recovered.
- `receivers` (Attributes Set) IAM user list to receive alarms. You can set 1–10 IAM users.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--receivers))

### Read-Only

- `alarm_notify_group_id` (String) Alarm notification group ID.
- `created_time` (String) Time when the alarm notification group was created.
- `id` (String) Uniquely identifies the resource.
- `updated_time` (String) Time when the alarm notification group was modified.

<a id="nestedatt--notice_rules"></a>
### Nested Schema for `notice_rules`

Optional:

- `has_end_node` (Boolean) Whether there is an end node afterwards.
- `has_next` (Boolean) Condition for whether to proceed to the next level.
- `receiver_infos` (Attributes Set) Notification channel information.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--notice_rules--receiver_infos))
- `rule_node` (String) Rule node. JSON format.

<a id="nestedatt--notice_rules--receiver_infos"></a>
### Nested Schema for `notice_rules.receiver_infos`

Optional:

- `alarm_content_template_id` (String) Alarm content template ID.
- `alarm_webhook_at_groups` (Set of String) User group name to notify when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `alarm_webhook_at_users` (Set of String) Username to notify when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `alarm_webhook_integration_id` (String) Alarm webhook integration configuration ID.
- `alarm_webhook_integration_name` (String) Name of the alarm Webhook integration configuration.
- `alarm_webhook_is_at_all` (Boolean) Whether to notify everyone when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `end_time` (String) End time for receiving alarm notifications. Uses 24-hour format: HH:mm:ss, with a valid range of 00:00:00–23:59:59. StartTime cannot be greater than EndTime.
- `general_webhook_body` (String) Custom WebHook request body. It is recommended to set the request body according to the callback interface requirements of the corresponding service.
- `general_webhook_headers` (Attributes Set) Custom callback request headers for the interface.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--notice_rules--receiver_infos--general_webhook_headers))
- `general_webhook_method` (String) Custom callback method for the interface. Only POST or PUT is supported.
- `general_webhook_url` (String) Custom callback URL for the interface.
- `receiver_channels` (Set of String) Notification channels. Supports one or more channels. Options: Email, Sms, Phone, GeneralWebhook, Lark, DingTalk, WeChat.
- `receiver_names` (Set of String) IAM user or user group name.
- `receiver_type` (String) Recipient type. Options: User: IAM user; UserGroup: IAM user group.
- `start_time` (String) Alarm notification start time. Uses 24-hour format (HH:mm:ss), valid range is 00:00:00–23:59:59. StartTime cannot be later than EndTime.

<a id="nestedatt--notice_rules--receiver_infos--general_webhook_headers"></a>
### Nested Schema for `notice_rules.receiver_infos.general_webhook_headers`

Optional:

- `key` (String) Custom request header key.
- `value` (String) Custom request header value.




<a id="nestedatt--receivers"></a>
### Nested Schema for `receivers`

Optional:

- `alarm_content_template_id` (String) Alarm content template ID.
- `alarm_webhook_at_groups` (Set of String) User group name to notify when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `alarm_webhook_at_users` (Set of String) Username to notify when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `alarm_webhook_integration_id` (String) Alarm webhook integration configuration ID.
- `alarm_webhook_integration_name` (String) Name of the alarm Webhook integration configuration.
- `alarm_webhook_is_at_all` (Boolean) Whether to notify everyone when sending notifications to Feishu, DingTalk, or WeCom via Webhook integration.
- `end_time` (String) End time for receiving alarm notifications. Uses 24-hour format: HH:mm:ss, with a valid range of 00:00:00–23:59:59. StartTime cannot be greater than EndTime.
- `general_webhook_body` (String) Custom WebHook request body. It is recommended to set the request body according to the callback interface requirements of the corresponding service.
- `general_webhook_headers` (Attributes Set) Custom callback request headers for the interface.
 Important Note: When using SetNestedAttribute, you must fully define all attributes of its nested structure. Incomplete definitions may cause Terraform to detect unexpected differences during plan comparison, triggering unnecessary resource updates and affecting resource stability and predictability. (see [below for nested schema](#nestedatt--receivers--general_webhook_headers))
- `general_webhook_method` (String) Custom callback method for the interface. Only POST or PUT is supported.
- `general_webhook_url` (String) Custom callback URL for the interface.
- `receiver_channels` (Set of String) Notification channels. Supports one or more channels. Options: Email, Sms, Phone, GeneralWebhook, Lark, DingTalk, WeChat.
- `receiver_names` (Set of String) IAM user or user group name.
- `receiver_type` (String) Recipient type. Options: User: IAM user; UserGroup: IAM user group.
- `start_time` (String) Alarm notification start time. Uses 24-hour format (HH:mm:ss), valid range is 00:00:00–23:59:59. StartTime cannot be later than EndTime.

<a id="nestedatt--receivers--general_webhook_headers"></a>
### Nested Schema for `receivers.general_webhook_headers`

Optional:

- `key` (String) Custom request header key.
- `value` (String) Custom request header value.

## Import

Import is supported using the following syntax:

```shell
$ terraform import volcenginecc_tls_alarm_notify_group.example "alarm_notify_group_id"
```
