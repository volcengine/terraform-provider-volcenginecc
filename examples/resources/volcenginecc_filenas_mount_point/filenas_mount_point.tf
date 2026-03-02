resource "volcenginecc_filenas_mount_point" "FileNASMountPointDemo" {
  file_system_id      = "enas-cnbja0f8*****"
  mount_point_name    = "test-1"
  permission_group_id = "pgroup-01bc1182"
  subnet_id           = "subnet-btepcsc5*****"
  vpc_id              = "vpc-3nr6adcn064u8931*****"
}