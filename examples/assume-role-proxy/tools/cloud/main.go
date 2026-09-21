package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/iam"
	"github.com/volcengine/volcengine-go-sdk/service/sts"
	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: cloud ACTION [JSON or @FILE]")
	}
	ak, sk := os.Getenv("VOLCENGINE_ACCESS_KEY"), os.Getenv("VOLCENGINE_SECRET_KEY")
	if ak == "" || sk == "" {
		return fmt.Errorf("VOLCENGINE_ACCESS_KEY and VOLCENGINE_SECRET_KEY are required")
	}
	region := os.Getenv("VOLCENGINE_REGION")
	if region == "" {
		region = "cn-beijing"
	}
	sess, err := session.NewSession(volcengine.NewConfig().WithRegion(region).WithCredentials(credentials.NewStaticCredentials(ak, sk, os.Getenv("VOLCENGINE_SESSION_TOKEN"))).WithMaxRetries(1))
	if err != nil {
		return err
	}
	input := map[string]interface{}{}
	if len(os.Args) > 2 {
		data := []byte(os.Args[2])
		if len(data) > 0 && data[0] == '@' {
			data, err = os.ReadFile(string(data[1:]))
			if err != nil {
				return err
			}
		}
		if err = json.Unmarshal(data, &input); err != nil {
			return err
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client := iam.New(sess, volcengine.NewConfig().WithEndpoint("iam.volcengineapi.com"))
	var out interface{}
	switch os.Args[1] {
	case "identity":
		out, err = sts.New(sess, volcengine.NewConfig().WithEndpoint("sts.volcengineapi.com")).GetCallerIdentityWithContext(ctx, &sts.GetCallerIdentityInput{})
	case "list-roles":
		out, err = client.ListRolesCommonWithContext(ctx, &input)
	case "get-role":
		out, err = client.GetRoleCommonWithContext(ctx, &input)
	case "create-role":
		out, err = client.CreateRoleCommonWithContext(ctx, &input)
	case "delete-role":
		out, err = client.DeleteRoleCommonWithContext(ctx, &input)
	case "get-policy":
		out, err = client.GetPolicyCommonWithContext(ctx, &input)
	case "create-policy":
		out, err = client.CreatePolicyCommonWithContext(ctx, &input)
	case "delete-policy":
		out, err = client.DeletePolicyCommonWithContext(ctx, &input)
	case "attach-policy":
		out, err = client.AttachRolePolicyCommonWithContext(ctx, &input)
	case "detach-policy":
		out, err = client.DetachRolePolicyCommonWithContext(ctx, &input)
	case "list-role-policies":
		out, err = client.ListAttachedRolePoliciesCommonWithContext(ctx, &input)
	case "describe-vpcs":
		out, err = vpc.New(sess).DescribeVpcsCommonWithContext(ctx, &input)
	default:
		return fmt.Errorf("unsupported action")
	}
	if err != nil {
		return err
	}
	return json.NewEncoder(os.Stdout).Encode(out)
}
