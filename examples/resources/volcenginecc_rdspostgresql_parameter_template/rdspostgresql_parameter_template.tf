resource "volcenginecc_rdspostgresql_parameter_template" "RDSPostgreSQLParameterTemplatDdemo" {
  template_name         = "全部测试参数模板"
  template_type         = "PostgreSQL"
  template_type_version = "PostgreSQL_13"
  template_desc         = "全部测试参数模板,描述场景"
  template_params = [
    {
      name  = "auto_explain.log_analyze"
      value = "on"
    },
    {
      name  = "auto_explain.log_buffers"
      value = "off"
    }
  ]
}