resource "volcenginecc_dns_record" "DNSRecordDemo" {
  zid    = "xxxxxx"
  host   = "test"
  type   = "CNAME"
  value  = "test1.example1.com"
  line   = "default"
  remark = "test"
  ttl    = 601
  weight = 2
  enable = false
}