resource "volcenginecc_rocketmq_allow_list" "RocketMQAllowListDemo" {
  allow_list_type = "IPv4"
  allow_list      = "192.xxx.0.0/24"
  allow_list_name = "ccapi-test"
  allow_list_desc = "this is a description"
}