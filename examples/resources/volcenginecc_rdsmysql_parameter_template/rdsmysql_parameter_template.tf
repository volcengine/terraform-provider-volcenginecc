resource "volcenginecc_rdsmysql_parameter_template" "RDSMySQLParameterTemplateDemo" {
  template_name         = "test-parametertemplate"
  template_type         = "Mysql"
  template_type_version = "MySQL_5_7"
  template_params = [
    {
      name          = "auto_increment_increment"
      running_value = "33"
    },
    {
      name          = "automatic_sp_privileges"
      running_value = "on"
    }
  ]
  template_desc = "测试参数模板"
  engine_type   = "InnoDB"
  project_name  = "default"
}