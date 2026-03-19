resource "volcenginecc_cdn_domain" "CDNDomainDemo" {
  area_access_rule = {
    rule_type = ""
    switch    = false
  }
  cache = [
    {
      cache_action = {
        action         = "cache"
        default_policy = "force_cache"
        ttl            = 2
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "path"
            operator = "match"
            type     = "url"
            value    = "/index.shtml;/index.html;/index.php;/index.aspx;/index.htm;/;/C"
          }
        ]
        connective = "OR"
      }
    },
    {
      cache_action = {
        action         = "cache"
        default_policy = "default"
        ttl            = 0
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "filetype"
            operator = "match"
            type     = "url"
            value    = "php;jsp;asp;aspx"
          }
        ]
        connective = "OR"
      }
    },
    {
      cache_action = {
        action         = "cache"
        default_policy = "default"
        ttl            = 2592000
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "directory"
            operator = "match"
            type     = "url"
            value    = "/"
          }
        ]
        connective = "OR"
      }
    }
  ]
  cache_key = [
    {
      cache_key_action = {
        cache_key_components = [
          {
            action      = "exclude"
            ignore_case = false
            object      = "queryString"
            subobject   = "*"
          }
        ]
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "path"
            operator = "match"
            type     = "url"
            value    = "/A"
          }
        ]
        connective = "OR"
      }
    },
    {
      cache_key_action = {
        cache_key_components = [
          {
            action      = "includePart"
            ignore_case = true
            object      = "queryString"
            subobject   = "A"
          }
        ]
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "filetype"
            operator = "match"
            type     = "url"
            value    = "jpg;jpeg;png;gif;webp;bmp;ico;tiff;htm;shtml;html;css;js;xml;json;bin;zip;rar;ipa;apk;sis;xap;msi;exe;cab;7z;txt;wmv;mp3;wma;ogg;flv;mp4;avi;mpg;mpeg;f4v;hlv;rmvb;rm;3pg;img;m3u8;ts;swf;A"
          }
        ]
        connective = "OR"
      }
    },
    {
      cache_key_action = {
        cache_key_components = [
          {
            action      = "include"
            ignore_case = false
            object      = "queryString"
            subobject   = "*"
          }
        ]
      }
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "directory"
            operator = "match"
            type     = "url"
            value    = "/"
          }
        ]
        connective = "OR"
      }
    }
  ]
  compression = {
    compression_rules = [
      {
        compression_action = {
          compression_format = "all"
          compression_target = "*"
          compression_type   = ["gzip", "br"]
          max_file_size_kb   = 2048
          min_file_size_kb   = 0
        }
        condition = {
          condition_rule = [
            {
              name     = ""
              object   = "path"
              operator = "match"
              type     = "url"
              value    = "/F"
            }
          ]
          connective = "OR"
        }
      },
      {
        compression_action = {
          compression_format = "all"
          compression_target = "*"
          compression_type   = ["gzip", "br"]
          max_file_size_kb   = 2048
          min_file_size_kb   = 0
        }

        condition = {
          condition_rule = [
            {
              name     = ""
              object   = "directory"
              operator = "match"
              type     = "url"
              value    = "/E/"
            }
          ]
          connective = "OR"
        }
      },
      {
        compression_action = {
          compression_format = "all"
          compression_target = "*"
          compression_type   = ["gzip", "br"]
          max_file_size_kb   = 2048
          min_file_size_kb   = 0
        }
        condition = {
          condition_rule = [
            {
              name     = ""
              object   = "filetype"
              operator = "match"
              type     = "url"
              value    = "D"
            }
          ]
          connective = "OR"
        }
      },
      {
        compression_action = {
          compression_format = "default"
          compression_target = "*"
          compression_type   = ["gzip"]
          max_file_size_kb   = 2048
          min_file_size_kb   = 0
        }
      }
    ]
    switch = true
  }
  conditional_origin = {
    origin_rules = [
      {
        actions = {
          origin_lines = [
            {
              address       = "192.168.0.4"
              http_port     = "80"
              https_port    = "443"
              instance_type = "ip"
              origin_host   = "192.168.0.5"
            },
            {
              address       = "192.168.0.5"
              http_port     = "80"
              https_port    = "443"
              instance_type = "ip"
              origin_host   = "192.168.0.6"
            }
          ]
        }
        condition = {
          condition_groups = [
            {
              condition = {
                object   = "path"
                operator = "equal"
                value    = ["/a"]
              }
            },
            {
              condition = {
                object   = "directory"
                operator = "equal"
                value    = ["/b/"]
              }
            },
            {
              condition = {
                object   = "filetype"
                operator = "equal"
                value    = ["c"]
              }
            },
            {
              condition = {
                object   = "client_ip"
                operator = "equal"
                value    = ["192.168.0.3"]
              }
            },
            {
              condition = {
                object   = "client_ip_operator"
                operator = "belong"
                value    = ["CT", "CM", "CU", "CMTIETONG", "CBN", "CERNET", "DRPENG"]
              }
            }
          ]
          connective = "and"
          is_group   = true
        }
      }
    ]
    switch = true
  }
  custom_error_page     = { switch = false }
  customize_access_rule = { switch = false }
  domain                = "ccapi-test-web.com"
  download_speed_limit  = { switch = false }
  follow_redirect       = true
  https = {
    disable_http = false
    forced_redirect = {
      enable_forced_redirect = false
      status_code            = "301"
    }
    http2 = false
    hsts = {
      switch = false
      ttl    = 0
    }
    ocsp   = false
    switch = false
  }
  http_forced_redirect = {
    enable_forced_redirect = false
    status_code            = "301"
  }
  i_pv_6 = { switch = false }
  method_denied_rule = {
    methods = ""
    switch  = false
  }
  multi_range = { switch = false }
  offline_cache = {
    object      = ""
    status_code = ""
    switch      = false
  }
  origin = [
    {
      origin_action = {
        origin_lines = [
          {
            address               = "192.168.0.1"
            http_port             = "80"
            https_port            = "443"
            instance_type         = "ip"
            origin_host           = "a.com"
            origin_type           = "primary"
            private_bucket_access = false
            weight                = "1"
          },
          {
            address               = "b.com"
            http_port             = "80"
            https_port            = "443"
            instance_type         = "domain"
            origin_host           = "c.com"
            origin_type           = "primary"
            private_bucket_access = false
            weight                = "1"
          },
          {
            address               = "ccapi-test-red.tos-cn-beijing.volces.com"
            http_port             = ""
            https_port            = ""
            instance_type         = "tos"
            origin_host           = "ccapi-test-red.tos-cn-beijing.volces.com"
            origin_type           = "primary"
            private_bucket_access = true
            private_bucket_auth = {
              auth_type = "tos"
              switch    = true
            }
            weight = "1"
          },
          {
            address               = "192.168.0.2"
            http_port             = "80"
            https_port            = "443"
            instance_type         = "ip"
            origin_host           = "d.com"
            origin_type           = "backup"
            private_bucket_access = false
            weight                = "1"
          }
        ]
      }
    }
  ]
  origin_arg = [
    {
      condition = {
        condition_rule = [
          {
            name     = ""
            object   = "directory"
            operator = "match"
            type     = "url"
            value    = "/"
          }
        ]
        connective = "OR"
      }
      origin_arg_action = {
        origin_arg_components = [
          {
            action    = "include"
            object    = "queryString"
            subobject = "*"
          }
        ]
      }
    }
  ]
  origin_cert_check = { switch = false }
  origin_host       = ""
  origin_i_pv_6     = "followclient"
  origin_protocol   = "followclient"
  origin_range      = true
  origin_retry = {
    status_code = ""
    switch      = false
  }
  origin_rewrite = { switch = false }
  origin_sni = {
    sni_domain = ""
    switch     = false
  }
  page_optimization = {
    optimization_type = [
      "html", "js", "css"
    ]
    switch = true
  }
  project             = "default"
  redirection_rewrite = { switch = false }
  remote_auth         = { switch = false }
  request_block_rule  = { switch = false }
  rewrite_hls = {
    sign_name = ""
    switch    = false
  }
  service_region = "outside_chinese_mainland"
  service_type   = "web"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  timeout       = { switch = false }
  url_normalize = { switch = false }
  video_drag    = { switch = false }
}