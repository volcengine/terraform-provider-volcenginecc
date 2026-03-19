resource "volcenginecc_tls_index" "TLSIndexDemo" {
  enable_auto_index = true
  full_text = {
    case_sensitive = true
    delimiter = chomp(<<-EOT
  , '";$#!=()[]{}?@&<>/:\n\t\r\\%*+-._^`|~QWE
EOT
    )
    include_chinese = true
  }
  key_value = [
    {
      key = "a"
      value = {
        auto_index_flag = false
        case_sensitive  = true
        delimiter = chomp(<<-EOT
  , '";$#!=()[]{}?@&<>/:\n\t\r\\%*+-._^`|~QWE
EOT
        )
        include_chinese = true
        index_all       = false
        sql_flag        = true
        value_type      = "text"
        index_sql_all   = false
      }
    },
    {
      key = "b"
      value = {
        auto_index_flag = false
        case_sensitive  = false
        delimiter       = ""
        include_chinese = false
        index_all       = false
        sql_flag        = true
        value_type      = "long"
        index_sql_all   = false
      }
    },
    {
      key = "c"
      value = {
        auto_index_flag = false
        case_sensitive  = false
        delimiter       = ""
        include_chinese = false
        index_all       = false
        sql_flag        = true
        value_type      = "double"
        index_sql_all   = false
      }
    },
    {
      key = "d"
      value = {
        auto_index_flag = false
        case_sensitive  = true
        delimiter = chomp(<<-EOT
  , '";$#!=()[]{}?@&<>/:\n\t\r\\%*+-._^`|~QWE
EOT
        )
        include_chinese = true
        index_all       = false
        index_sql_all   = true
        json_keys = [
          {
            key = "d1"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "text"
              index_sql_all = false
            }
          },
          {
            key = "d2"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "long"
              index_sql_all = false
            }
          },
          {
            key = "d3"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "double"
              index_sql_all = false
            }
          }
        ]
        sql_flag   = true
        value_type = "json"
      }
    },
    {
      key = "e"
      value = {
        auto_index_flag = false
        case_sensitive  = true
        delimiter = chomp(<<-EOT
  , '";$#!=()[]{}?@&<>/:\n\t\r\\%*+-._^`|~QWE
EOT
        )
        include_chinese = true
        index_all       = true
        index_sql_all   = false
        json_keys = [
          {
            key = "e1"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "text"
              index_sql_all = false
            }
          }
        ]
        sql_flag   = true
        value_type = "json"
      }
    }
  ]
  max_text_len = 20480
  topic_id     = "7126b1fb-b68a-47a5-8c99-1a226601ed1a"
  user_inner_key_value = [
    {
      key = "__content__"
      value = {
        auto_index_flag = false
        case_sensitive  = true
        delimiter = chomp(<<-EOT
  , '";$#!=()[]{}?@&<>/:\n\t\r\\%*+-._^`|~QWE
EOT
        )
        include_chinese = true
        index_all       = true
        index_sql_all   = false
        json_keys = [
          {
            key = "f1"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "text"
              index_sql_all = false
            }
          },
          {
            key = "f2"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "long"
              index_sql_all = false
            }
          },
          {
            key = "f3"
            value = {
              index_all     = false
              sql_flag      = true
              value_type    = "double"
              index_sql_all = false
            }
          }
        ]
        sql_flag   = true
        value_type = "json"
      }
    }
  ]
}