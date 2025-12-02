resource "volcenginecc_apig_gateway" "APIGGatewayDemo" {
  name       = "APIGGatewayDemo"
  type       = "standard"
  comments   = "APIGGatewayDemo comments"
  vpc_id     = "vpc-13f8k4dwdsydc3n6nu5rxxxxx"
  subnet_ids = ["subnet-***", "subnet-***"]
  resource_spec = {
    instance_spec_code          = "1c2g"
    replicas                    = 2
    clb_spec_code               = "small_1"
    public_network_bandwidth    = 1
    public_network_billing_type = "bandwidth"
    network_type = {
      enable_public_network  = true
      enable_private_network = true
    }
  }
  monitor_spec = {
    enable       = true
    workspace_id = "***"
  }
  log_spec = {
    enable     = true
    project_id = "***"
    topic_id   = "***"
  }
  trace_spec = {
    enable     = true
    trace_type = "tls"
    tls_trace_spec = {
      project_id  = "***"
      iam_user_ak = "***"
      iam_user_sk = "***"
    }
  }
}