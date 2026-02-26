resource "volcenginecc_ecs_deployment_set" "EcsDeploymentSetDemo" {
  description                 = "this is a test DeploymentSet"
  deployment_set_name         = "test-deployment-set"
  granularity                 = "host"
  instance_ids                = ["i-yedvixxxxxva4izkjtl"]
  deployment_set_group_number = 1
  strategy                    = "Availability"
}