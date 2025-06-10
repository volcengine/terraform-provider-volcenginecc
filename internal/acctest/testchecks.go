// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"context"
	"fmt"
	"time"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	tfcloudcontrol "github.com/volcengine/terraform-provider-volcenginecc/internal/service/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
)

const (
	deleteResourceTimeout = 120 * time.Minute
)

// CheckDestroy returns a TestCheckFunc that tests whether a resource has been destroyed in Volcengine.
func (td TestData) CheckDestroy() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		for resourceName, resourceState := range state.RootModule().Resources {
			if resourceState.Type != td.TerraformResourceType {
				continue
			}

			if resourceName != td.ResourceName {
				continue
			}

			return td.checkExists(false)(state)
		}

		return nil
	}
}

// CheckExistsInVolcengine returns a TestCheckFunc that tests whether a resource exists in Volcengine.
func (td TestData) CheckExistsInVolcengine() resource.TestCheckFunc {
	return td.checkExists(true)
}

// DeleteResource returns a TestCheckFunc that deletes a resource in Volcengine.
func (td TestData) DeleteResource() resource.TestCheckFunc {
	return func(state *terraform.State) error {
		resourceName := td.ResourceName
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id := rs.Primary.ID

		if id == "" {
			return fmt.Errorf("no ID is set")
		}

		provider, err := getCloudControlAPIProvider(td.provider)

		if err != nil {
			return err
		}

		ctx := getTestContext()

		return tfcloudcontrol.DeleteResource(ctx, provider.CloudControlAPIClient(ctx), "", "", td.CloudControlResourceType, id)
	}
}

func (td TestData) checkExists(shouldExist bool) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		provider, err := getCloudControlAPIProvider(td.provider)

		if err != nil {
			return err
		}

		ctx := getTestContext()

		return existsFunc(shouldExist)(
			ctx,
			provider.CloudControlAPIClient(ctx),
			"",
			td.CloudControlResourceType,
			td.TerraformResourceType,
			td.ResourceName,
		)(state)
	}
}

func existsFunc(shouldExist bool) func(context.Context, *cloudcontrol.CloudControl, string, string, string, string) resource.TestCheckFunc {
	return func(ctx context.Context, conn *cloudcontrol.CloudControl, roleARN, cfTypeName, tfTypeName, resourceName string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			rs, ok := state.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("not found: %s", resourceName)
			}

			id := rs.Primary.ID

			if id == "" {
				return fmt.Errorf("no ID is set")
			}

			_, err := tfcloudcontrol.FindResourceByTypeNameAndID(ctx, conn, "", cfTypeName, id)

			// TODO
			// TODO Some resource can still be found but are logically deleted.
			// TODO

			if !shouldExist {
				if err == nil {
					return fmt.Errorf("(%s/%s) resource (%s) still exists", cfTypeName, tfTypeName, id)
				}

				if tfresource.NotFound(err) {
					return nil
				}
			}

			if err != nil {
				return fmt.Errorf("reading (%s/%s) resource (%s): %w", cfTypeName, tfTypeName, id, err)
			}

			return nil
		}
	}
}

func getCloudControlAPIProvider(provider fwprovider.Provider) (tfcloudcontrol.Provider, error) {
	if v, ok := provider.(interface{ ProviderData() any }); ok {
		if v, ok := v.ProviderData().(tfcloudcontrol.Provider); ok {
			return v, nil
		}
	}

	return nil, fmt.Errorf("unable to convert %T to Cloud Control API Provider", provider)
}

func getTestContext() context.Context {
	return tfsdklog.NewRootProviderLogger(context.TODO(), tflog.WithLevelFromEnv("TF_LOG"), tflog.WithoutLocation())
}
