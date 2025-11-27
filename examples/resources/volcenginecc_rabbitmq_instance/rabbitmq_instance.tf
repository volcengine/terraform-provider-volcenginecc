resource "volcenginecc_rabbitmq_instance" "RabbitMQInstanceDemo" {
  zone_id              = "cn-beijing-a"
  user_name            = "RabbitMQInstanceDemo"
  compute_spec         = "rabbitmq.n1.x4.small"
  version              = "3.12"
  user_password        = "********"
  storage_space        = 100
  instance_description = "RabbitMQInstanceDemo"
  vpc_id               = "vpc-1a1vgeo93yxxx8nvepjxxxxx"
  charge_detail = {
    charge_type = "PostPaid"
  }
  subnet_id     = "subnet-ij9s4hxxxs3k74o8cuxxxxx"
  eip_id        = "eip-1c0qhbjo7xxxw5e8j70axxxxx"
  instance_name = "RabbitMQInstanceDemo"
  tags = [
    {
      key = "env"
    value = "test" }
  ]
  project_name = "default"
}