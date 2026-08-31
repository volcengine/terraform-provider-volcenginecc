resource "volcenginecc_iam_api_key" "example" {
  user_name            = "user-xxxx"
  name                 = "ccapi-apikey-full"
  permissions_boundary = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"iam:*\"],\"Resource\":[\"*\"]}]}"
}
