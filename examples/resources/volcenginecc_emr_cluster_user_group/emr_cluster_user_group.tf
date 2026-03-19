resource "volcenginecc_emr_cluster_user_group" "EMRClusterUserGroupDemo" {
  cluster_id      = "emr-xxxxxxxxxx"
  user_group_name = "ccapi-tf-users-2"
  description     = "EMRClusterUserGroupDemo"
  members         = ["emrclusteruserdemo1", "emrclusteruserdemo3"]
}