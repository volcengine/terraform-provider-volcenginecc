resource "volcenginecc_iam_policy" "PolicyDemo" {
  policy_name     = "PolicyDemo"
  policy_type     = "Custom"
  description     = "PolicyDemo description"
  policy_document = "{\"Statement\":[{\"Action\":[\"clb:DescribeAclAttributes\",\"clb:DescribeHealthCheckLogProjectAttributes\",\"clb:DescribeListenerAttributes\",\"clb:DescribeListenerHealth\",\"clb:DescribeLoadBalancerAttributes\",\"clb:DescribeLoadBalancerHealth\",\"clb:DescribeLoadBalancersBilling\",\"clb:DescribeNLBListenerAttributes\",\"clb:DescribeNLBListenerCertificates\",\"clb:DescribeNLBListenerHealth\",\"clb:DescribeNLBListeners\",\"clb:DescribeNLBServerGroupAttributes\",\"clb:DescribeNLBServerGroups\",\"clb:DescribeNLBZones\",\"clb:DescribeNetworkLoadBalancerAttributes\",\"clb:DescribeNetworkLoadBalancers\",\"clb:DescribeServerGroupAttributes\",\"clb:DescribeZones\",\"clb:ListTagsForNLBResources\",\"clb:DescribeAcls\",\"clb:DescribeCertificates\",\"clb:DescribeHealthCheckLogTopicAttributes\",\"clb:DescribeListeners\",\"clb:DescribeLoadBalancerSpecs\",\"clb:DescribeLoadBalancers\",\"clb:DescribeRules\",\"clb:DescribeServerGroups\",\"clb:ListTagsForResources\",\"clb:TagNLBResources\",\"clb:TagResources\",\"clb:UntagNLBResources\",\"clb:UntagResources\"],\"Effect\":\"Allow\",\"Resource\":[\"*\"]}]}"
  policy_users = [
    {
      "name" : "test",
      "policy_scope" : [
        {
          "project_name" : "default",
          "policy_scope_type" : "test",
          "project_display_name" : "test"
        }
      ]
    }
  ]
  policy_roles = [
    {
      "name" : "roles",
      "policy_scope" : [
        {
          "project_name" : "default"
        }
      ]
    }
  ]
  policy_user_groups = [
    {
      "name" : "test",
      "policy_scope" : [
        {
          "project_name" : "default"
        }
      ]
    }
  ]
}