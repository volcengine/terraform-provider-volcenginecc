resource "volcenginecc_ecs_command" "CommandDemo" {
  command_content = "IyEvYmluL2Jhc2gKCm1rZGlyIHt7ZGxxxxx"
  description = "CommandDemo Example"
  enable_parameter = true
  name = "commandtest"
  parameter_definitions = [
    {
      type = "Digit"
      name = "dirname"
      required = true
      default_value = "10"
      min_length = 0
      max_length = 0
      min_value = "5"
      max_value = "100"
      decimal_precision = 0    }
  ]
  project_name = "default"
  tags = [
    {
      key = "env"
      value = "test"
    }
  ]
  timeout = 60
  type = "Shell"
  username = "rxxxx"
  working_dir = "/home"
}