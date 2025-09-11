resource "volcenginecc_storageebs_volume" "VolumeDemo" {

  volume_name = "EBS-VolumeDemo"
  volume_type = "ESSD_PL0"
  size = 10
  zone_id = "cn-beijing-x"
}