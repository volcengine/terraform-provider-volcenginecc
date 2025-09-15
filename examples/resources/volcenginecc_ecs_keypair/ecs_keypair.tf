resource "volcenginecc_ecs_keypair" "ECSKeypairDemo" {
  key_pair_name = "ECSKeypairDemo"
  project_name  = "default"
  description   = "ECSKeypairDemo Description"
  tags = [
    {
      key   = "env"
      value = "test"
    }
  ]
}