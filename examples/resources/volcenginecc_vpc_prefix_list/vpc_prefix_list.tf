resource "volcenginecc_vpc_prefix_list" "PrefixlistDemo" {
  description      = "PrefixlistDemo Example"
  ip_version       = "IPv4"
  max_entries      = 20
  prefix_list_name = "my-prefix-list"
  project_name     = "default"
  prefix_list_entries = [
    {
      cidr        = "192.168.0.0/*"
      description = "privite description"
    }
  ]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}