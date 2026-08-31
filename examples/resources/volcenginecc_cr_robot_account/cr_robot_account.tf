resource "volcenginecc_cr_robot_account" "example" {
  registry    = "cr_instance.id"
  username    = "crrobot@ccapi-full"
  description = "ccapi robotaccount full"
  duration    = 90
  permissions = [
    {
      resource = "namespace-xx"
      actions  = ["PullRepository", "PushRepository"]
    },
    {
      resource = "project-xxx"
      actions  = ["CreateRepository", "PullRepository"]
    }
  ]
}
