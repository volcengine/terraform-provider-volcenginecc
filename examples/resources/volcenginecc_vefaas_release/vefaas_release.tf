resource "volcenginecc_vefaas_release" "VEFAASReleaseDemo" {
  function_id           = "xxxxxxxx"
  revision_number       = 0
  target_traffic_weight = 100
  rolling_step          = 100
  description           = "ccapi-test"
  max_instance          = 10
}