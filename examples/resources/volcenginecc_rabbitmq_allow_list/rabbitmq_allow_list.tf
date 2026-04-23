resource "volcenginecc_rabbitmq_allowlist" "RabbitMQAllowListDemo" {
  allow_list_type = "IPv4"
  allow_list      = "192.x.0.0/24"
  allow_list_name = "ccapi-test"
  allow_list_desc = "test-desc"
  associated_instances = [
    {
      instance_id = "rbtmq-a69238e***"
    }
  ]
}