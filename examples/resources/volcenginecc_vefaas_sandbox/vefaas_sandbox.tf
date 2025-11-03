resource "volcenginecc_vefaas_sandbox" "VefaasSandboxDemo" {
  function_id     = "n5hcs3y5****"
  timeout         = 50
  memory_mb       = 2048
  cpu_milli       = 1000
  request_timeout = 30
  max_concurrency = 100
  metadata = [
    {
      key = "env"
    value = "test" }
  ]
  instance_image_info = {
    image    = "enterprise-public-cn-beijing.cr.volces.com/xxxxxx/all-in-one-sandbox:xxxxxx"
    command  = "/opt/gem/run.sh"
    port     = 8080
    image_id = "3ewzg8x5h1***"
  }
  instance_tos_mount_config = {
    enable = true
    tos_mount_points = [
      {
        bucket_path = "/mnt/tos"
      local_mount_path = "/mnt/tos" }
    ]
  }
  envs = [
    {
      key = "env"
    value = "test" }
  ]
}