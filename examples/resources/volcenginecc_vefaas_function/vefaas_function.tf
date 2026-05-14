resource "volcenginecc_vefaas_function" "VefaasFunctionDemo" {
  name            = "VefaasFunction1001"
  runtime         = "native/v1"
  exclusive_mode  = false
  max_concurrency = 10
  memory_mb       = 512
  request_timeout = 30
  port            = 8000
  vpc_config = {
    enable_vpc                    = true
    vpc_id                        = "vpc-rrco37ovjq4gxxxxxxx"
    subnet_ids                    = ["subnet-rrwqhg3qzxxxxxxxx"]
    security_group_ids            = ["sg-rrco3fkzzy0wxxxxxxxxxxxx"]
    enable_shared_internet_access = true
  }
  command        = "./run.sh"
  enable_apmplus = false
  project_name   = "default"
  nas_storage = {
    enable_nas = true
    nas_configs = [
      {
        remote_path      = "/"
        local_mount_path = "/mnt/nas"
        mount_point_id   = "mount-xxxxxxx"
        file_system_id   = "enas-cnbja0bxxxxxxx"
        gid              = 1000
        uid              = 1000
      }
    ]
  }
  tos_mount_config = {
    enable_tos = true
    credentials = {
      access_key_id     = "AK**************"
      secret_access_key = "*****************=="
    }
    mount_points = [
      {
        bucket_path      = "/"
        local_mount_path = "/mnt/tos"
        read_only        = false
        bucket_name      = "ccapi-register-test"
      endpoint = "http://tos-cn-beijing.ivolces.com" }
    ]
  }
  envs = [
    {
      key = "env"
    value = "test" }
  ]
  async_task_config = {
    enable_async_task = true
    max_retry         = 2
    destination_config = {
      on_failure = {
        destination = "https://ccapi-failure.com"
      }
      on_success = {
        destination = "https://ccapi-success.com"
      }
    }
  }
  role = "trn:iam::*******:role/xxxxxxx"
  tls_config = {
    enable_log     = true
    tls_project_id = "9248d829-21b0-43e2-a2f6-xxxxxxxxx"
    tls_topic_id   = "84baf7cf-0f60-44e0-a5f6-xxxxxxxxx"
  }
  source                    = "https://vefaas-prod-xxxxxxxxxPAYLOAD&X-Tos-Credential=AKxxxxxxxxxxxxFtos-cn-beijing.volces.com%2Ftos%2Frequest&X-Tos-Date=20260327T080542Z&X-Tos-Expires=1800&X-Tos-SignedHeaders=host&X-Tos-Signature=5a60f816def2be59baxxxxxxxxxxxxx"
  source_type               = "tos"
  enable_dependency_install = true
}