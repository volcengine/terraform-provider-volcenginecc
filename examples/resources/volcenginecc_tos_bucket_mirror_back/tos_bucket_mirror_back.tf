resource "volcenginecc_tos_bucket_mirror_back" "TOSBucketMirrorBackDemo" {
  bucket = "ccapi-test"
  rules = [{
    condition = {
      http_code   = 404
      http_method = ["GET", "HEAD"]
      key_prefix  = "object-key-prefix"
      key_suffix  = "object-key-suffix"
      allow_host  = ["example.volcengine.com"]
    }
    id = "rule-003"
    redirect = {
      fetch_header_to_meta_data_rules = [{
        meta_data_suffix = "-meta"
        source_header    = "X-Custom-Header"
      }]
      fetch_source_on_redirect            = false
      fetch_source_on_redirect_with_query = false
      follow_redirect                     = true
      mirror_header = {
        pass_all = true
        pass     = ["aaa", "bbb"]
        remove   = ["xxx", "yyy"]
        set = [{
          key   = "X-Set-Header"
          value = "set-value"
        }]
      }
      pass_header_from_source      = ["X-Source-Header"]
      pass_query                   = true
      pass_status_code_from_source = [404, 500]
      private_source = {
        source_endpoint = {
          primary = [{
            bucket_name = "primary-bucket"
            endpoint    = "http://xxxxx.volces.com"
            credential_provider = {
              static_credential = {
                storage_vendor  = "S3"
                sk_encrypt_type = ""
                sk              = "sk-test-primary"
                ak              = "ak-test-primary"
              }
              role   = ""
              region = "cn-beijing"
            }
          }]
          follower = [{
            bucket_name = "follower-bucket"
            endpoint    = "http://xxxxx.volces.com"
            credential_provider = {
              static_credential = {
                storage_vendor  = "S3"
                sk_encrypt_type = ""
                sk              = "sk-test-follower"
                ak              = "ak-test-follower"
              }
              role   = ""
              region = "cn-beijing"
            }
          }]
        }
      }
      redirect_type = "Mirror"
      transform = {
        with_key_prefix = "addtional-key-prefix"
        with_key_suffix = "addtional-key-suffix"
        replace_key_prefix = {
          key_prefix   = "key-prefix"
          replace_with = "replace-with"
        }
      }
    }
  }]
}