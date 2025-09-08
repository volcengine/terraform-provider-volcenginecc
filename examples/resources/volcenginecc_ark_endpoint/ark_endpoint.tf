resource "volcenginecc_ark_endpoint" "EndpointDemo" {
  name            = "EndpointDemo"
  description     = "endpoint for tf demo"
  model_reference = {
    foundation_model = {
      name          = "doubao-1-5-thinking-***"
      model_version = "250428"
    }
  }

  tags = [
    {
      key   = "env"
      value = "test"
    }


  ]
  project_name = "default"
  rate_limit   = {
    tpm = 5
    rpm = 10
  }
}