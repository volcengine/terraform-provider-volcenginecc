resource "volcenginecc_gtm_gtm" "GTMGTMDemo" {
  domain       = "www.example4.com"
  access_mode  = "cname"
  cname        = "gtm-cname.example4.com.bdns-gtm-pressure.com"
  project_name = "default"
  remark       = "gtm test"
  spec_name    = "ultimate"
  ttl          = 600
  policy_type  = "perf"
  alarm_id     = "71xxxx26d79-2b08-4033-b025-71xxxx26d79"
}