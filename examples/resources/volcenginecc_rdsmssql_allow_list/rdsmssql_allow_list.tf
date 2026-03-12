resource "volcenginecc_rdsmssql_allow_list" "RDSMsSQLAllowlistDemo" {
  project_name        = "default"
  allow_list_name     = "ccapi-test-all"
  allow_list_desc     = "ccapi-test"
  allow_list_type     = "IPv4"
  allow_list_category = "Ordinary"
  allow_list          = "192.168.0.0/24,192.168.0.1,192.168.0.2"
  user_allow_list     = "192.168.0.0/24,192.168.0.1,192.168.0.2"
  associated_instances = [
    {
      instance_id = "mssql-9b195******"
    }
  ]
}