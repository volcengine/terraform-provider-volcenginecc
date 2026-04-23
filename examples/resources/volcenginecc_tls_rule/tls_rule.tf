resource "volcenginecc_tls_rule" "TLSRuleDemo" {
  container_rule = {
    container_name_regex = "ccapi-test"
    env_tag = [{
      key = "k1"
      val = "v1"
    }]
    exclude_container_env_regex = [{
      key = "k2"
      val = "v2"
    }]
    exclude_container_label_regex = [{
      key = "k2"
      val = "v2"
    }]
    include_container_env_regex = [{
      key = "k1"
      val = "v1"
    }]
    include_container_label_regex = [{
      key = "k1"
      val = "v1"
    }]
    kubernetes_rule = {
      annotation_tag = [{
        key = "k1"
        val = "v1"
      }]
      enable_all_label_tag = true
      exclude_pod_annotation_regex = [{
        key = "k2"
        val = "v2"
      }]
      exclude_pod_label_regex = [{
        key = "k2"
        val = "v2"
      }]
      include_pod_annotation_regex = [{
        key = "k1"
        val = "v1"
      }]
      include_pod_label_regex = [{
        key = "k1"
        val = "v1"
      }]
      label_tag = [{
        key = "k1"
        val = "v1"
      }]
      namespace_name_regex = "ccapi"
      pod_name_regex       = "ccapi"
      workload_name_regex  = "ccapi"
      workload_type        = "Deployment"
    }
    stream = ""
  }
  exclude_paths = [{
    type  = "Path"
    value = "/test"
  }]
  extract_rule = {
    begin_regex       = ""
    delimiter         = " "
    enable_nanosecond = true
    filter_key_regex = [{
      key   = "k1"
      regex = "v1"
    }]
    keys      = ["k1"]
    log_regex = ""
    log_template = {
      format = ""
      type   = ""
    }
    quote                   = ""
    time_extract_regex      = "dsfasdf"
    time_format             = "%Y-%m-%d %H:%M:%S.%f"
    time_key                = "time_key"
    time_sample             = "2006-01-02T15:04:05Z07:00"
    time_zone               = "Etc/GMT+12"
    un_match_log_key        = "LogParseFailed"
    un_match_up_load_switch = true
  }
  host_group_infos = [{
    host_group_id = "38feb3b9-xxxxx-de52cxxxxx"
  }]
  input_type = 2
  log_sample = "asdfasf"
  log_type   = "delimiter_log"
  paths      = ["/test"]
  rule_name  = "dx-rule-1"
  topic_id   = "b881e3cd-xxxxx-966f-fe98xxxxxxx"
  user_define_rule = {
    advanced = {
      close_eof                       = false
      close_inactive                  = 10
      close_removed                   = false
      close_renamed                   = false
      close_timeout                   = 0
      no_line_terminator_eof_max_time = 0
    }
    enable_host_group_label = true
    enable_hostname         = true
    enable_raw_log          = true
    fields = [{
      key = "k1"
      val = "v1"
    }]
    host_group_label_key = "host_group_label"
    hostname_key         = "hostname"
    ignore_older         = 72
    multi_collects_type  = "RuleID"
    parse_path_rule = {
      keys        = ["instance-id", "pod-name"]
      path_sample = "/var/logs/instanceid_any_podname/test.log"
      regex       = "/var/logs/([a-z]*)_any_([a-z]*)/test.log"
    }
    plugin = {
      processors = "{}"
    }
    raw_log_key = "raw"
    shard_hash_key = {
      hash_key = "2342134"
    }
    tail_files   = true
    tail_size_kb = 10240
  }
}