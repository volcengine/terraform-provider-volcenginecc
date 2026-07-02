resource "volcenginecc_directconnect_connection" "Example" {
  bandwidth                      = 500
  customer_contact_email         = "zhaoliu@example.com"
  customer_contact_phone         = "136xxxx6000"
  customer_name                  = "赵六"
  description                    = "其他运营商物理专线"
  direct_connect_access_point_id = "ap-cn-beijing-cy-A"
  direct_connect_connection_name = "其他运营商专线"
  line_operator                  = "ChinaTelecom"
  peer_location                  = "深圳市南山区某IDC"
  port_spec                      = "1G"
  port_type                      = "1000Base"
  project_name                   = "default"
  tags = [{
    key   = "env"
    value = "test"
  }]
}