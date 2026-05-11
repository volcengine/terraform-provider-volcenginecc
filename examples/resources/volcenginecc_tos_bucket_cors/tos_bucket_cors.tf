resource "volcenginecc_tos_bucket_cors" "TOSBucketCorsDemo" {
  bucket_name = "ccapi-test"
  cors_rules = [{
    allowed_origins = ["http://test1.example.com:8080", "http://test2.example.com:8088"]
    allowed_methods = ["PUT", "DELETE"]
    expose_headers  = ["x-tos-version-id"]
    max_age_seconds = 7200
    allowed_headers = ["Content-Type", "Authorization"]
    response_vary   = false
    }, {
    allowed_origins = ["https://*.example.org"]
    allowed_methods = ["HEAD"]
    expose_headers  = ["x-tos-request-id"]
    max_age_seconds = 1800
    allowed_headers = ["*"]
    response_vary   = true
  }]
}