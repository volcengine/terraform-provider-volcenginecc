resource "volcenginecc_vepfs_fileset" "VEPFSFilesetDemo" {
  fileset_name           = "ccapi-test"
  file_system_id         = "vepfs-cnbj81448xxxxx"
  fileset_path_on_create = "/test1/"
}