resource "volcenginecc_emr_cluster_user" "EMRClusterUserDemo" {
  cluster_id       = "emr-xxxxxxxxxx"
  user_name        = "emrclusteruserdemo"
  password         = "UserDemo1234"
  user_group_names = ["users", "ccapi"]
  description      = "EMRClusterUserDemo"
}