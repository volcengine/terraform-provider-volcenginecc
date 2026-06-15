resource "volcenginecc_iam_allowed_ip_address" "IAMAllowedIpAddressDemo" {
  user_id = "2109xxxxx"
  ip_list = [
    {
      ip          = "192.168.1.100",
      description = "test1"
    },
    {
      ip          = "192.168.2.100",
      description = "test2"
    },
    {
      ip          = "192.168.3.100",
      description = "test3"
    }
  ]
}