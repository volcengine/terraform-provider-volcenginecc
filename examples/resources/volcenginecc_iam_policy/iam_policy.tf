resource "volcenginecc_iam_policy" "policyDemo" {
  policy_name     = "policy-demo"
  policy_type     = "Custom"
  description     = "policy-demo description"
  policy_document = "{\"Statement\":[{\"Action\":[\"clb:DescribeAclAttributes\",\"clb:DescribeHealthCheckLogProjectAttributes\",\"clb:DescribeListenerAttributes\",\"clb:DescribeListenerHealth\",\"clb:DescribeLoadBalancerAttributes\",\"clb:DescribeLoadBalancerHealth\",\"clb:DescribeLoadBalancersBilling\",\"clb:DescribeNLBListenerAttributes\",\"clb:DescribeNLBListenerCertificates\",\"clb:DescribeNLBListenerHealth\",\"clb:DescribeNLBListeners\",\"clb:DescribeNLBServerGroupAttributes\",\"clb:DescribeNLBServerGroups\",\"clb:DescribeNLBZones\",\"clb:DescribeNetworkLoadBalancerAttributes\",\"clb:DescribeNetworkLoadBalancers\",\"clb:DescribeServerGroupAttributes\",\"clb:DescribeZones\",\"clb:ListTagsForNLBResources\",\"clb:DescribeAcls\",\"clb:DescribeCertificates\",\"clb:DescribeHealthCheckLogTopicAttributes\",\"clb:DescribeListeners\",\"clb:DescribeLoadBalancerSpecs\",\"clb:DescribeLoadBalancers\",\"clb:DescribeRules\",\"clb:DescribeServerGroups\",\"clb:ListTagsForResources\",\"clb:TagNLBResources\",\"clb:TagResources\",\"clb:UntagNLBResources\",\"clb:UntagResources\"],\"Effect\":\"Allow\",\"Resource\":[\"*\"]}]}"
  policy_users    = [
    {
      "name" : "xxxx",
      "policy_scope" : [
        {
          "project_name" : "default",
          "policy_scope_type" : "xxxx",
          "project_display_name" : "xxxx"
        },
        {
          "project_name" : "xxxx"
        }
      ]
    },
    {
      "name" : "xxxx.xxx",
      "policy_scope" : [
        {
          "project_name" : "xxxx",
          "policy_scope_type" : "xxxx",
          "project_display_name" : "xxxx"
        },
        {
          "project_name" : "xxxx"
        }
      ]
    }
  ]
  policy_roles = [
    {
      "name" : "xxxxx",
      "policy_scope" : [
        {
          "project_name" : "xxxxx"
        }
      ]
    },
    {
      "name" : "xxxxx",
      "policy_scope" : [
        {
          "project_name" : "xxxx"
        }
      ]
    },
    {
      "name" : "xxxxx",
      "policy_scope" : [
        {
          "project_name" : "xxxx"
        },
        {
          "project_name" : "xxxx"
        }
      ]
    }
  ]
  policy_user_groups = [
    {
      "name" : "xxxx-test",
      "policy_scope" : [
        {
          "project_name" : "xxxx"
        }
      ]
    },
    {
      "name" : "xxxx-dev",
      "policy_scope" : [
        {
          "project_name" : "xxxx"
        },
        {
          "project_name" : "xxxx"
        }
      ]
    }
  ]
}