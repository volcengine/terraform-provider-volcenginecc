resource "volcenginecc_gtm_policy" "GTMPolicyDemo" {
  gtm_id      = "003cd3af-dc55-xxxxx-xxxxxx-41470d406367"
  policy_type = "perf"
  targets = [{
    pool_id = "ad65b9fb-e8f5-xxxxx-xxxxxx-bc4f2741cd66"
  }]
  alarm_only   = false
  routing_mode = "geo-lb"
}