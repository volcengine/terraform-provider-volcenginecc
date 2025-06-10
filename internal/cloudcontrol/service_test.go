package cloudcontrol

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/test-go/testify/suite"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

type CloudControlSDKTestSuite struct {
	suite.Suite
	client *CloudControl
}

func (s *CloudControlSDKTestSuite) TestCreateResource() {
	body := `
{
	"Name": "resource-ccapi-test-ws",
	"Description": "created by ccapi",
	"Runtime": "python3.12/v1",
	"ExclusiveMode": false,
	"RequestTimeout": 10,
	"MaxConcurrency": 10,
	"MemoryMB": 1024,
	"Source": "ccapi-test:iam-user.zip",
	"SourceType": "tos",
	"ProjectName": "default",
	"Envs": [
		{
			"Key": "key1",
			"Value": "value1"
		},
		{
			"Key": "key2",
			"Value": "value2"
		}
	],
	"VpcConfig": {
		"VpcId": "vpc-***********",
		"SubnetIds": [
			"subnet-**********"
		],
		"EnableVpc": true,
		"SecurityGroupIds": [
			"sg-2c00bg8naq3gg2dx0efxfcdnv"
		],
		"EnableSharedInternetAccess": false
	},
	"TlsConfig": {
		"EnableLog": false
	},
	"NasStorage": {
		"EnableNas": false
	},
	"TosMountConfig": {
		"EnableTos": true,
		"Credentials":{
			"AccessKeyId":"**************",
			"SecretAccessKey":"***********"
		},
		"MountPoints": [
			{
				"Endpoint": "https://tos-cn-boe.ivolces.com",
				"ReadOnly": true,
				"BucketName": "ccapi-test",
				"BucketPath": "/",
				"LocalMountPath": "/mnt/tos"
			}
		]
	}
}
	`

	state := make(map[string]any)

	err := json.Unmarshal([]byte(body), &state)
	s.NoError(err)

	input := &CreateResourceInput{
		TypeName:    volcengine.String("Volcengine::VEFAAS::Function"),
		TargetState: &state,
	}

	output, err := s.client.CreateResourceWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func (s *CloudControlSDKTestSuite) TestGetResource() {
	input := &GetResourceInput{
		TypeName:   volcengine.String("Volcengine::VEFAAS::Function"),
		Identifier: volcengine.String("3k5n73q2"),
	}

	output, err := s.client.GetResourceWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func (s *CloudControlSDKTestSuite) TestListResource() {
	input := &ListResourceInput{
		TypeName:   volcengine.String("Volcengine::VEFAAS::Function"),
		MaxResults: 4,
	}

	output, err := s.client.ListResourceWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func (s *CloudControlSDKTestSuite) TestDeleteResource() {
	input := &DeleteResourceInput{
		TypeName:   volcengine.String("Volcengine::VEFAAS::Function"),
		Identifier: volcengine.String("3k5n73q2"),
	}

	output, err := s.client.DeleteResourceWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func (s *CloudControlSDKTestSuite) TestUpdateResource() {
	input := &UpdateResourceInput{
		TypeName:   volcengine.String("Volcengine::VEFAAS::Function"),
		Identifier: volcengine.String("vi4kn9lr"),
		PatchDocument: []any{
			map[string]any{
				"op":    "replace",
				"path":  "/Description",
				"value": "updated by ccapi sdk",
			},
		},
	}

	output, err := s.client.UpdateResourceWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func (s *CloudControlSDKTestSuite) TestGetTask() {
	input := &GetTaskInput{
		TaskID: volcengine.String("task-0433123c-5c46-43f5-8597-af05b98b5f76"),
	}

	output, err := s.client.GetTaskWithContext(context.Background(), input)
	s.NoError(err)
	s.NotNil(output)
}

func TestCloudControlSDKTestSuite(t *testing.T) {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, _ := session.NewSession(config)

	suite.Run(t, &CloudControlSDKTestSuite{
		client: New(sess),
	})
}
