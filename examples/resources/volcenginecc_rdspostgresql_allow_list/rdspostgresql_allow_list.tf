resource "volcenginecc_rdspostgresql_allow_list" "RdsPostgresqlAllowListDemo" {
  allow_list          = ["1.2.3.4/32", "5.6.7.8/32"]
  allow_list_category = "Ordinary"
  allow_list_desc     = "test"
  allow_list_name     = "ccapi-test-1"
  allow_list_type     = "IPv4"
}