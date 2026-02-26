resource "volcenginecc_privatezone_record" "PrivateZoneRecordDemo" {
  host   = "ccapi-test-1"
  line   = "default"
  remark = "test"
  ttl    = 600
  type   = "AAAA"
  value  = "ff03:0:0:0:0:0:0:c1"
  weight = 1
  zid    = 403215
  enable = true
}