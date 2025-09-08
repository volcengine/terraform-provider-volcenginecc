resource "volcenginecc_iam_accesskey" "TestTerraform" {
  user_name = "s2222"
  # 禁用的时候传
  status    = "inactive"
}