resource "volcenginecc_tls_host_group" "TLSHostGroupDemo" {
  auto_update       = true
  host_group_name   = "ccapi-test-1"
  host_group_type   = "IP"
  host_ip_list      = ["192.168.1.1", "192.168.1.2"]
  service_logging   = false
  update_end_time   = "06:00"
  update_start_time = "00:00"
  iam_project_name  = "default"
}