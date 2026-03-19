resource "volcenginecc_vedbm_database" "VEDBMDatabaseDemo" {
  db_name            = "testdb-tf"
  instance_id        = "vedbm-ls2ehotj2***"
  character_set_name = "utf8mb4"
  db_desc            = "desctest"
  databases_privileges = [
    {
      account_name             = "test",
      account_privilege        = "Custom",
      account_privilege_detail = ["SELECT", "UPDATE", "INSERT"]
    }
  ]

}