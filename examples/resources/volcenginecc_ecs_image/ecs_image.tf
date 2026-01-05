resource "volcenginecc_ecs_image" "ImageDemo" {
  description      = "ImageDemo Example"
  image_name       = "image-demo"
  instance_id      = "i-ydzhj1el8gr9cxxdnxxxx"
  project_name     = "default"
  share_permission = ["2000000***"]
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}