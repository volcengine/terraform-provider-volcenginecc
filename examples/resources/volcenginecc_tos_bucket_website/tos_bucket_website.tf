resource "volcenginecc_tos_bucket_website" "TOSBucketWebsiteDemo" {
  bucket = "ccapi-test"
  index_document = {
    suffix            = "index.html"
    forbidden_sub_dir = false
  }
  error_document = {
    key = "error.html"
  }
  routing_rules = [
    {
      condition = {
        http_error_code_returned_equals = 404
        key_prefix_equals               = "red/"
      }
      redirect = {
        host_name               = "example.com"
        http_redirect_code      = 302
        protocol                = "https"
        replace_key_prefix_with = "redirect/"
      }
    }
  ]
}