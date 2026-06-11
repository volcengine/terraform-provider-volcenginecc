resource "volcenginecc_gtm_rule" "GTMRuleDemo" {
  disable                  = false
  effective_pool_set_index = 0
  gtm_id                   = "gtm_id_xxxxx"
  line                     = "default"
  name                     = "Protocol-https"
  policy_type              = "policy_type_xxxxx"
  pool_set_mode            = "manual"
  pool_sets = [{
    name            = "主地址池集"
    active_addr_thr = 1
    pools = [{
      pool_id = "pool_id_xxxxx"
      weight  = 1
    }]
  }]
  probe = {
    advised_node_count = 8
    disable            = false
    failed_count       = 3
    protocol           = "https"
    port               = 443
    url                = "/content/pages"
    host               = "www.example.com"
    http_method        = "GET"
    http_usability_codes = [{
      operator = "interval"
      codes    = [100, 200]
    }]
    ping_count        = 0
    ping_loss_percent = 0
    interval          = 60
    is_manual_nodes   = false
    nodes             = ["China-North lq Bytedance", "China-North lf Bytedance", "China-North hl Bytedance"]
    timeout           = 5
  }
  remark                 = "test for create"
  use_rule_probe_config  = true
  use_policy_probe_nodes = false
  weight                 = 1
}