resource "volcenginecc_vedbm_allow_list" "VEDBMAllowListDemo" {
  allow_list_type = "IPv4"
  allow_list      = "192.168.0.0/27,192.168.1.0/26"
  allow_list_name = "test"
  allow_list_desc = "test"
  project_name    = "default"
  associated_instances = [
    {
    instance_id = "vedbm-7chs5mtj2***" },
    {
    instance_id = "vedbm-a5nwhmtj***" }
  ]
  modify_mode = "Cover"
}