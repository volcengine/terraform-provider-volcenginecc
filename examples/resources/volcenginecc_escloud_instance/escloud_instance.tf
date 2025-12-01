resource "volcenginecc_escloud_instance" "ESCloudInstanceDemo" {
  instance_configuration = {
    vpc = {
      vpc_id   = "vpc-1a1vgeo93xxxg8nvepjyxxxxx"
      vpc_name = "ESCloudInstanceDemo-vpc"
    }
    subnet = {
      subnet_id   = "subnet-3nrb3atxxxf40931eb4xxxxx"
      subnet_name = "ESCloudInstanceDemo-Subnet"
    }
    zone_id            = "cn-beijing-d"
    version            = "V7_10"
    region_id          = "cn-beijing"
    charge_type        = "PostPaid"
    enable_https       = true
    project_name       = "default"
    instance_name      = "ESCloudInstanceDemo"
    enable_pure_master = true
    node_specs_assigns = [
      {
        type               = "Hot"
        number             = 1
        storage_size       = 100
        storage_spec_name  = "es.volume.essd.pl0"
        resource_spec_name = "es.x2.medium"
        extra_performance = {
          throughput = 0
        }
      },
      {
        type               = "Master"
        number             = 3
        storage_size       = 20
        storage_spec_name  = "es.volume.essd.pl0"
        resource_spec_name = "es.x2.medium"
        extra_performance = {
          throughput = 0
        }
      },
      {
        type               = "Kibana"
        number             = 1
        storage_size       = 0
        storage_spec_name  = ""
        resource_spec_name = "kibana.x2.small"
        extra_performance = {
          throughput = 0
        }
      }
    ]
    configuration_code  = "es.standard"
    deletion_protection = true
    network_specs = [
      {
        type      = "Elasticsearch"
        bandwidth = 1
        is_open   = true
      spec_name = "es.eip.bgp_fixed_bandwidth" },
      {
        type      = "Kibana"
        bandwidth = 1
        is_open   = true
      spec_name = "es.eip.bgp_fixed_bandwidth" }
    ]
    admin_password = "********"
    tags = [
      {
        key = "env"
      value = "test" }
    ]
  }
}