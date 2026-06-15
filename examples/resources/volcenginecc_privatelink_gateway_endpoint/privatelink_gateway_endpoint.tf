resource "volcenginecc_privatelink_gateway_endpoint" "PrivateLinkGatewayEndpointDemo" {
  description   = "ccapi-test"
  vpc_id        = "vpc-2ntqlxxxxxxxxxrlvfu7z"
  endpoint_name = "test-endpoint"
  project_name  = "default"
  vpc_policy    = "{\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":\"*\",\"Resource\":\"*\",\"Condition\":null}]}"
  tags = [{
    key   = "env"
    value = "test"
  }]
  service_id = "gwepsvc-3rxxxxxxsk2ilz3f62"
}