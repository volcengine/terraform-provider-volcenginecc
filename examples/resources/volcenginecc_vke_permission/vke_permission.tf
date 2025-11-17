resource "volcenginecc_vke_permission" "VKEPermissionDemo" {
  role_domain    = "namespace"
  cluster_id     = "cd48m3cb1b2ba7l6ebgp0xxxxx"
  namespace      = "kube-public"
  role_name      = "vke:visitor"
  is_custom_role = false
  grantee_id     = 59433888
  grantee_type   = "User"
}