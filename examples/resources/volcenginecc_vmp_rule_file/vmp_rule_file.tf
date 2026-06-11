resource "volcenginecc_vmp_rulefile" "VMPRuleFileDemo" {
  description         = "这是一个测试规则文件"
  target_workspace_id = "3ba7844b-e7fc-4688-a869-****"
  content             = "groups:\n- name: example\n  rules:\n  - record: example_metric\n    expr: sum(up)"
  workspace_id        = "3ba7844b-e7fc-4688-a869-****"
  name                = "test-rule-file"
}