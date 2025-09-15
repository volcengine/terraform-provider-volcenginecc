resource "volcenginecc_clb_server_group" "CLBServerGroupDemo" {
  load_balancer_id   = "clb-mioj49on0zcw5smt1b8rxxxx"
  server_group_name  = "CLBServerGroupDemo"
  type               = "instance"
  description        = "CLBServerGroupDemo description"
  address_ip_version = "ipv4"
  any_port_enabled   = "off"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
  servers = [
    {
      instance_id      = "i-ye498lwge85i3z3kxxxx"
      type             = "ecs"
      weight           = 50
      ip               = "192.168.*.13"
      any_port_enabled = "off"
      port             = 8080
    description = "test" },
    {
      instance_id      = "i-ye48ymyy9s5i3z4pxxxx"
      type             = "ecs"
      weight           = 22
      ip               = "192.168.*.52"
      any_port_enabled = "off"
      port             = 91
    description = "test" }
  ]
}