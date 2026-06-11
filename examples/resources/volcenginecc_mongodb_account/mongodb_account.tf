resource "volcenginecc_mongodb_account" "MongoDBAccountDemo" {
  account_desc     = "this is a test"
  account_name     = "ccapi-test"
  account_password = "QWExxxxzxc"
  account_privileges = [{
    db_name   = "admin"
    role_name = "read"
    }, {
    db_name   = "admin"
    role_name = "userAdmin"
  }]
  auth_db     = "admin"
  instance_id = "mongo-replica-xxxxxxxx"
}