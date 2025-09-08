resource "volcenginecc_ecs_image" "ImageDemo" {
  boot_mode = "UEFI"
  description = "ImageDemo Example"
  image_name = "image-demo"
  instance_id = "i-ydzhj1el8gr9cxxdnxxxx"
  kernel = "Linux"
  license_type = "BYOL"
  os_name = "CentOS"
  os_type = "Linux"
  platform = "CentOS"
  platform_version = "8.3"
  project_name = "default"
  share_permission = ["2000000***"]
  tags = [
    {
      key = "env"
      value = "test"
    }
  ]
}