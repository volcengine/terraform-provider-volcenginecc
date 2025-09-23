// Copyright (c) HashiCorp, Inc.
// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0
// This file has been modified by Beijing Volcano Engine Technology Co., Ltd. on 2025

package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/common"
	baselogging "github.com/volcengine/terraform-provider-volcenginecc/internal/logging"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/registry"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

const (
	defaultMaxRetries         = 25
	defaultAssumeRoleDuration = 1 * time.Hour
)

// providerData is returned from the provider's Configure method and
// is passed to each resource and data source in their Configure methods.
type providerData struct {
	ccAPIClient *cloudcontrol.CloudControl
	logger      baselogging.Logger
	region      string
}

func (p *providerData) CloudControlAPIClient(_ context.Context) *cloudcontrol.CloudControl {
	return p.ccAPIClient
}

func (p *providerData) Region(_ context.Context) string {
	return p.region
}

func (p *providerData) RegisterLogger(ctx context.Context) context.Context {
	return baselogging.RegisterLogger(ctx, p.logger)
}

type VolcengineCCProvider struct {
	providerData *providerData // Used in acceptance tests.
}

func New() provider.Provider {
	return &VolcengineCCProvider{}
}

// ProviderData is used in acceptance testing to get access to configured API client etc.
func (p *VolcengineCCProvider) ProviderData() any {
	return p.providerData
}

func (p *VolcengineCCProvider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "volcenginecc"
	response.Version = Version
}

func (p *VolcengineCCProvider) Schema(ctx context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_key": schema.StringAttribute{
				Description: "The Access Key for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_ACCESS_KEY` environment variable",
				Optional:    true,
			},
			"secret_key": schema.StringAttribute{
				Description: "he Secret Key for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_SECRET_KEY` environment variable",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "The Region for Volcengine Provider. It must be provided, but it can also be sourced from the `VOLCENGINE_REGION` environment variable",
				Optional:    true,
			},
			"disable_ssl": schema.BoolAttribute{
				Description: "Disable SSL for Volcengine Provider",
				Optional:    true,
			},
			"customer_headers": schema.StringAttribute{
				Optional:    true,
				Description: "CUSTOMER HEADERS for Volcengine Provider. The customer_headers field uses commas (,) to separate multiple headers, and colons (:) to separate each header key from its corresponding value.",
			},
			"proxy_url": schema.StringAttribute{
				Optional:    true,
				Description: "PROXY URL for Volcengine Provider",
			},
			"assume_role": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"assume_role_trn": schema.StringAttribute{
						Description: "he TRN of the role to assume.",
						Required:    true,
					},
					"duration_seconds": schema.Int32Attribute{
						Description: "The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.",
						Optional:    true,
						Validators: []validator.Int32{
							int32validator.Between(900, 43200),
						},
					},
					"policy": schema.StringAttribute{
						Description: "A more restrictive policy when making the AssumeRole call",
						Optional:    true,
					},
				},
				Optional:    true,
				Description: "An `assume_role` block (documented below). Only one `assume_role` block may be in the configuration.",
			},
			"endpoints": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"cloudcontrolapi": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default Cloud Control API service endpoint URL",
					},
					"sts": schema.StringAttribute{
						Optional:    true,
						Description: "Use this to override the default STS service endpoint URL",
					},
				},
				Optional:    true,
				Description: "An `endpoints` block (documented below). Only one `endpoints` block may be in the configuration.",
			},
		},
	}
}

type configModel struct {
	AccessKey       types.String    `tfsdk:"access_key"`
	SecretKey       types.String    `tfsdk:"secret_key"`
	Region          types.String    `tfsdk:"region"`
	DisableSSL      types.Bool      `tfsdk:"disable_ssl"`
	CustomerHeaders types.String    `tfsdk:"customer_headers"`
	ProxyURL        types.String    `tfsdk:"proxy_url"`
	AssumeRole      *AssumeRoleData `tfsdk:"assume_role"`
	Endpoints       *endpointData   `tfsdk:"endpoints"`

	terraformVersion string
}
type AssumeRoleData struct {
	AssumeRoleTRN types.String `tfsdk:"assume_role_trn"`
	Duration      types.Int32  `tfsdk:"duration_seconds"`
	Policy        types.String `tfsdk:"policy"`
}

type endpointData struct {
	CloudControlAPI types.String `tfsdk:"cloudcontrolapi"`
	STS             types.String `tfsdk:"sts"`
}

func (p *VolcengineCCProvider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	var config configModel

	response.Diagnostics.Append(request.Config.Get(ctx, &config)...)
	if response.Diagnostics.HasError() {
		return
	}

	if !request.Config.Raw.IsFullyKnown() {
		response.Diagnostics.AddError("Unknown Value", "An attribute value is not yet known")
	}

	config.terraformVersion = request.TerraformVersion
	if config.AccessKey.IsNull() || config.AccessKey.IsUnknown() {
		config.AccessKey = types.StringValue(os.Getenv("VOLCENGINE_ACCESS_KEY"))
	}
	if config.SecretKey.IsNull() || config.SecretKey.IsUnknown() {
		config.SecretKey = types.StringValue(os.Getenv("VOLCENGINE_SECRET_KEY"))
	}
	if config.Region.IsNull() || config.Region.IsUnknown() {
		config.Region = types.StringValue(os.Getenv("VOLCENGINE_REGION"))
	}
	if config.AccessKey.ValueString() == "" {
		response.Diagnostics.AddError("Missing AccessKey", "AccessKey must be set")
	}
	if config.SecretKey.ValueString() == "" {
		response.Diagnostics.AddError("Missing SecretKey", "SecretKey must be set")
	}
	if config.Region.ValueString() == "" {
		response.Diagnostics.AddError("Missing Region", "Region must be set")
	}
	if config.DisableSSL.IsNull() || config.DisableSSL.IsUnknown() {
		if disableSSLString := os.Getenv("VOLCENGINE_DISABLE_SSL"); disableSSLString != "" {
			disableSSLBool, _ := strconv.ParseBool(disableSSLString)
			config.DisableSSL = types.BoolValue(disableSSLBool)
		}
	}
	if config.ProxyURL.IsNull() || config.ProxyURL.IsUnknown() {
		if proxyURL := os.Getenv("VOLCENGINE_PROXY_URL"); proxyURL != "" {
			config.ProxyURL = types.StringValue(proxyURL)
		}
	}
	if config.CustomerHeaders.IsNull() || config.CustomerHeaders.IsUnknown() {
		if customerHeader := os.Getenv("VOLCENGINE_CUSTOMER_HEADERS"); customerHeader != "" {
			config.CustomerHeaders = types.StringValue(customerHeader)
		}
	}
	if config.Endpoints == nil {
		config.Endpoints = &endpointData{}
	}
	if config.Endpoints.CloudControlAPI.IsNull() || config.Endpoints.CloudControlAPI.IsUnknown() {
		if ccEndpoint := os.Getenv("VOLCENGINE_CC_ENDPOINT"); ccEndpoint != "" {
			config.Endpoints.CloudControlAPI = types.StringValue(ccEndpoint)
		}
	}
	if config.Endpoints.STS.IsNull() || config.Endpoints.STS.IsUnknown() {
		if stsEndpoint := os.Getenv("VOLCENGINE_STS_ENDPOINT"); stsEndpoint != "" {
			config.Endpoints.STS = types.StringValue(stsEndpoint)
		}
	}
	if config.AssumeRole == nil {
		config.AssumeRole = &AssumeRoleData{
			AssumeRoleTRN: types.StringNull(),
			Duration:      types.Int32Null(),
			Policy:        types.StringNull(),
		}
	}
	if config.AssumeRole.AssumeRoleTRN.IsNull() || config.AssumeRole.AssumeRoleTRN.IsUnknown() {
		if trn := os.Getenv("VOLCENGINE_ASSUME_ROLE_TRN"); trn != "" {
			config.AssumeRole.AssumeRoleTRN = types.StringValue(trn)
		}
	}
	if config.AssumeRole.Duration.IsNull() || config.AssumeRole.Duration.IsUnknown() {
		if duration := os.Getenv("VOLCENGINE_ASSUME_ROLE_DURATION_SECONDS"); duration != "" {
			durationInt, _ := strconv.Atoi(duration)
			config.AssumeRole.Duration = types.Int32Value(int32(durationInt))
		}
	}
	if config.AssumeRole.Policy.IsNull() || config.AssumeRole.Policy.IsUnknown() {
		if policy := os.Getenv("VOLCENGINE_ASSUME_ROLE_POLICY"); policy != "" {
			config.AssumeRole.Policy = types.StringValue("VOLCENGINE_ASSUME_ROLE_Policy")
		}
	}

	providerData, diags := newProviderData(ctx, &config)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	p.providerData = providerData
	response.DataSourceData = providerData
	response.ResourceData = providerData
}

func (p *VolcengineCCProvider) Resources(ctx context.Context) []func() resource.Resource {
	var diags diag.Diagnostics
	var resources = make([]func() resource.Resource, 0)

	for name, factory := range registry.ResourceFactories() {
		v, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Resource",
				fmt.Sprintf("Error getting the %s Resource, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		resources = append(resources, func() resource.Resource {
			return v
		})
	}

	return resources
}

func (p *VolcengineCCProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	var diags diag.Diagnostics
	dataSources := make([]func() datasource.DataSource, 0)

	for name, factory := range registry.DataSourceFactories() {
		v, err := factory(ctx)

		if err != nil {
			diags.AddError(
				"Error getting Data Source",
				fmt.Sprintf("Error getting the %s Data Source, this is an error in the provider.\n%s\n", name, err),
			)

			continue
		}

		dataSources = append(dataSources, func() datasource.DataSource {
			return v
		})
	}

	return dataSources
}

func newProviderData(ctx context.Context, c *configModel) (*providerData, diag.Diagnostics) {
	var diags diag.Diagnostics
	version := fmt.Sprintf("%s/%s (terraform/%s)", common.TerraformProviderName, common.TerraformProviderVersion, c.terraformVersion)
	ctx, logger := baselogging.NewTfLogger(ctx)

	if diags.HasError() {
		return nil, diags
	}

	config := volcengine.NewConfig().
		WithRegion(c.Region.ValueString()).
		WithCredentials(credentials.NewStaticCredentials(c.AccessKey.ValueString(), c.SecretKey.ValueString(), "")).
		WithDisableSSL(c.DisableSSL.ValueBool()).
		WithExtendHttpRequest(func(ctx context.Context, request *http.Request) {
			request.Header.Set("user-agent", version)
		})

	if !(c.CustomerHeaders.IsNull() || c.CustomerHeaders.IsUnknown()) {
		customHeaderMap := make(map[string]string)
		headers := c.CustomerHeaders.ValueString()
		if headers != "" {
			hs1 := strings.Split(headers, ",")
			for _, hh := range hs1 {
				hs2 := strings.Split(hh, ":")
				if len(hs2) == 2 {
					customHeaderMap[hs2[0]] = hs2[1]
				}
			}
		}
		config.WithExtendHttpRequest(func(ctx context.Context, request *http.Request) {
			if len(customHeaderMap) > 0 {
				for k, v := range customHeaderMap {
					request.Header.Add(k, v)
				}
			}
		})
	}

	if c.Endpoints != nil && !c.Endpoints.CloudControlAPI.IsNull() {
		config.WithEndpoint(c.Endpoints.CloudControlAPI.ValueString())
	} else {
		config.WithEndpoint(fmt.Sprintf("cloudcontrol.%s.volcengineapi.com}", c.Region.ValueString()))
	}

	if !(c.ProxyURL.IsNull() || c.ProxyURL.IsUnknown()) {
		u, _ := url.Parse(c.ProxyURL.ValueString())
		t := &http.Transport{
			Proxy: http.ProxyURL(u),
		}
		httpClient := http.DefaultClient
		httpClient.Transport = t
	}

	if c.AssumeRole != nil && !c.AssumeRole.AssumeRoleTRN.IsNull() {
		accountId, roleName, err := ParseTrn(c.AssumeRole.AssumeRoleTRN.ValueString())
		if err != nil {
			diags.AddError(err.Error(), err.Error())
			return nil, diags

		}

		if c.AssumeRole.Duration.IsNull() || c.AssumeRole.Duration.IsUnknown() {
			c.AssumeRole.Duration = types.Int32Value(3600)
		}

		stsValue := credentials.StsValue{
			AccessKey:       c.AccessKey.ValueString(),
			SecurityKey:     c.SecretKey.ValueString(),
			RoleName:        roleName, // 扮演角色名称
			AccountId:       accountId,
			Region:          c.Region.ValueString(),
			DurationSeconds: int(c.AssumeRole.Duration.ValueInt32()),
		}
		if c.Endpoints != nil && !c.Endpoints.STS.IsNull() {
			stsValue.Host = c.Endpoints.STS.ValueString()
		}
		if c.DisableSSL.ValueBool() {
			stsValue.Schema = "http"
		}
		config.WithCredentials(credentials.NewStsCredentials(stsValue))
	}

	sess, err := session.NewSession(config)
	if err != nil {
		diags.AddError(err.Error(), err.Error())
		return nil, diags
	}

	cloudcontrolClient := cloudcontrol.New(sess)
	if err != nil {
		diags.AddError(err.Error(), err.Error())
		return nil, diags
	}

	providerData := &providerData{
		ccAPIClient: cloudcontrolClient,
		logger:      logger,
		region:      c.Region.String(),
	}

	return providerData, diags
}

func ParseTrn(trn string) (string, string, error) {
	re := regexp.MustCompile(`^trn:iam::([^:]+):role/(.+)$`)
	matches := re.FindStringSubmatch(trn)
	if len(matches) == 3 {
		accountId := matches[1]
		roleName := matches[2]
		return accountId, roleName, nil
	} else {
		return "", "", errors.New("invalid trn")
	}
}
