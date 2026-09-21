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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/common"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/customresources"
	baselogging "github.com/volcengine/terraform-provider-volcenginecc/internal/logging"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/registry"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials/clicreds"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"golang.org/x/net/http/httpproxy"
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
				Description: "The Access Key for Volcengine Provider. It can also be sourced from the `VOLCENGINE_ACCESS_KEY` environment variable",
				Optional:    true,
			},
			"secret_key": schema.StringAttribute{
				Description: "The Secret Key for Volcengine Provider. It can also be sourced from the `VOLCENGINE_SECRET_KEY` environment variable",
				Optional:    true,
			},
			"session_token": schema.StringAttribute{
				Description: "The Session Token for Volcengine Provider. It can also be sourced from the `VOLCENGINE_SESSION_TOKEN` environment variable",
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
				Description: "HTTP, HTTPS, SOCKS5, or SOCKS5H proxy URL for Cloud Control API requests. It can also be sourced from the `VOLCENGINE_PROXY_URL` environment variable.",
			},
			"proxy_authorization": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Value of the Proxy-Authorization header for Cloud Control API proxy requests, for example `Basic <token>`. It can also be sourced from the `VOLCENGINE_PROXY_AUTHORIZATION` environment variable.",
			},
			"no_proxy": schema.StringAttribute{
				Optional:    true,
				Description: "Comma-separated hosts, domain suffixes, IP addresses, or CIDR ranges that bypass proxy_url. It follows standard NO_PROXY matching and can be sourced from VOLCENGINE_NO_PROXY, NO_PROXY, or no_proxy.",
			},
			"proxy_include_domains": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Hosts, domain suffixes, IP addresses, or CIDR ranges that use proxy_url while all other destinations connect directly. It can be sourced as a comma-separated list from VOLCENGINE_PROXY_INCLUDE_DOMAINS and cannot be combined with no_proxy.",
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
					"proxy_url": schema.StringAttribute{
						Description: "HTTP, HTTPS, SOCKS5, or SOCKS5H proxy URL for AssumeRole STS requests, including HTTPS requests. Defaults to VOLCENGINE_ASSUME_ROLE_PROXY_URL, HTTP_PROXY, then http_proxy. An explicit empty string skips these environment defaults. This proxy is independent of the top-level Cloud Control proxy settings and does not use NO_PROXY.",
						Optional:    true,
						Sensitive:   true,
					},
				},
				Optional:    true,
				Description: "An `assume_role` block that uses the selected source credentials to obtain target-role credentials. Only one `assume_role` block may be in the configuration.",
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
			"profile": schema.StringAttribute{
				Description: "The Profile for Volcengine Provider. It can be sourced from the `VOLCENGINE_PROFILE` environment variable. Complete AccessKey and SecretKey credentials take precedence when both sources are configured",
				Optional:    true,
			},
			"file_path": schema.StringAttribute{
				Description: "The Profile configuration file path for Volcengine Provider. It defaults to `~/.volcengine/config.json` and can be sourced from the `VOLCENGINE_FILE_PATH` environment variable",
				Optional:    true,
			},
		},
	}
}

type configModel struct {
	AccessKey           types.String    `tfsdk:"access_key"`
	SecretKey           types.String    `tfsdk:"secret_key"`
	SessionToken        types.String    `tfsdk:"session_token"`
	Region              types.String    `tfsdk:"region"`
	DisableSSL          types.Bool      `tfsdk:"disable_ssl"`
	CustomerHeaders     types.String    `tfsdk:"customer_headers"`
	ProxyURL            types.String    `tfsdk:"proxy_url"`
	ProxyAuthorization  types.String    `tfsdk:"proxy_authorization"`
	NoProxy             types.String    `tfsdk:"no_proxy"`
	ProxyIncludeDomains types.Set       `tfsdk:"proxy_include_domains"`
	AssumeRole          *AssumeRoleData `tfsdk:"assume_role"`
	Endpoints           *endpointData   `tfsdk:"endpoints"`
	Profile             types.String    `tfsdk:"profile"`
	FilePath            types.String    `tfsdk:"file_path"`
	terraformVersion    string
}
type AssumeRoleData struct {
	AssumeRoleTRN types.String `tfsdk:"assume_role_trn"`
	Duration      types.Int32  `tfsdk:"duration_seconds"`
	Policy        types.String `tfsdk:"policy"`
	ProxyURL      types.String `tfsdk:"proxy_url"`
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
	if config.SessionToken.IsNull() || config.SessionToken.IsUnknown() {
		config.SessionToken = types.StringValue(os.Getenv("VOLCENGINE_SESSION_TOKEN"))
	}
	if config.Profile.IsNull() || config.Profile.IsUnknown() {
		config.Profile = types.StringValue(os.Getenv("VOLCENGINE_PROFILE"))
	}
	if config.FilePath.IsNull() || config.FilePath.IsUnknown() {
		config.FilePath = types.StringValue(os.Getenv("VOLCENGINE_FILE_PATH"))
	}
	if config.Region.IsNull() || config.Region.IsUnknown() {
		config.Region = types.StringValue(os.Getenv("VOLCENGINE_REGION"))
	}
	// 认证分两阶段解析：先用环境变量填充 Terraform 未配置的字段，再按
	// AK/SK、Profile、默认凭证链的顺序选择源凭证。源凭证确定后，
	// buildCredentials 再按需使用 AssumeRole 包装它。
	if config.Region.ValueString() == "" {
		response.Diagnostics.AddError("Missing Region", "Region must be set")
	}
	if config.DisableSSL.IsNull() || config.DisableSSL.IsUnknown() {
		if disableSSLString := os.Getenv("VOLCENGINE_DISABLE_SSL"); disableSSLString != "" {
			disableSSLBool, _ := strconv.ParseBool(disableSSLString)
			config.DisableSSL = types.BoolValue(disableSSLBool)
		}
	}
	setProxyDefaultsFromEnvironment(&config)
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
			ProxyURL:      types.StringNull(),
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
			config.AssumeRole.Policy = types.StringValue(policy)
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

		// Allow hand-written code to wrap or replace the auto-generated
		// resource without modifying the generated `*_resource_gen.go`.
		// See internal/customresources for details.
		if wrap, ok := customresources.Lookup(name); ok {
			v, err = wrap(ctx, v)
			if err != nil {
				diags.AddError(
					"Error wrapping Resource",
					fmt.Sprintf("Error applying custom override for the %s Resource.\n%s\n", name, err),
				)
				continue
			}
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

type assumeRoleCredentialsFactory func(*credentials.Credentials, credentials.StsValue) *credentials.Credentials

// buildSourceCredentials 根据环境变量填充后的最终字段选择调用云服务或 STS 的源凭证。
// 完整 AK/SK 沿用原有优先级，高于 Profile/file_path；两种来源同时存在时继续使用
// AK/SK 并返回警告，避免用户在不知情的情况下使用了非预期身份。
func buildSourceCredentials(c *configModel) (*credentials.Credentials, diag.Diagnostics) {
	var diags diag.Diagnostics
	switch {
	case c.AccessKey.ValueString() != "" && c.SecretKey.ValueString() != "":
		if c.Profile.ValueString() != "" || c.FilePath.ValueString() != "" {
			diags.AddWarning(
				"Multiple Credential Sources Configured",
				"Both complete AccessKey/SecretKey credentials and Profile/file_path credentials are configured. The provider selected AccessKey/SecretKey according to the credential precedence and ignored Profile/file_path. Remove the unused credential configuration or environment variables to avoid using an unintended identity.",
			)
		}
		return credentials.NewStaticCredentials(c.AccessKey.ValueString(), c.SecretKey.ValueString(), c.SessionToken.ValueString()), diags
	case c.Profile.ValueString() != "" || c.FilePath.ValueString() != "":
		return clicreds.NewCliCredentials(c.FilePath.ValueString(), c.Profile.ValueString()), diags
	default:
		return defaults.NewDefaultCredentialProvider(), diags
	}
}

// buildCredentials 先选择可刷新的源凭证，再按需使用 AssumeRole 凭证提供方包装源凭证。
// AssumeRole 不再与 Profile、静态凭证或默认凭证链竞争优先级。
func buildCredentials(c *configModel) (*credentials.Credentials, diag.Diagnostics) {
	return buildCredentialsWithFactory(c, newAssumeRoleCredentials)
}

// buildCredentialsWithFactory 使用可注入的工厂构造最终凭证，使测试能够验证源凭证
// 选择与 AssumeRole 参数，而无需发起真实 STS 请求。源凭证的产生过程对 Provider
// 保持透明；静态凭证、Profile 和默认凭证链都可以作为一次 Provider AssumeRole 的输入。
func buildCredentialsWithFactory(c *configModel, factory assumeRoleCredentialsFactory) (*credentials.Credentials, diag.Diagnostics) {
	sourceCredentials, diags := buildSourceCredentials(c)
	if diags.HasError() {
		return nil, diags
	}
	if !hasAssumeRole(c) {
		return sourceCredentials, diags
	}

	accountId, roleName, err := ParseTrn(c.AssumeRole.AssumeRoleTRN.ValueString())
	if err != nil {
		diags.AddError("Invalid AssumeRole TRN", err.Error())
		return nil, diags
	}
	if c.AssumeRole.Duration.IsNull() || c.AssumeRole.Duration.IsUnknown() {
		c.AssumeRole.Duration = types.Int32Value(int32(defaultAssumeRoleDuration.Seconds()))
	}

	proxyURL, err := assumeRoleProxyURL(c.AssumeRole.ProxyURL)
	if err != nil {
		diags.AddError("Invalid AssumeRole Proxy Configuration", err.Error())
		return nil, diags
	}

	stsValue := credentials.StsValue{
		RoleName:        roleName,
		AccountId:       accountId,
		Schema:          "https",
		Region:          c.Region.ValueString(),
		DurationSeconds: int(c.AssumeRole.Duration.ValueInt32()),
		Policy:          c.AssumeRole.Policy.ValueString(),
		ProxyURL:        proxyURL,
	}
	if c.Endpoints != nil && !c.Endpoints.STS.IsNull() && c.Endpoints.STS.ValueString() != "" {
		stsValue.Host = c.Endpoints.STS.ValueString()
	} else {
		stsValue.Host = "sts.volcengineapi.com"
	}
	if c.DisableSSL.ValueBool() {
		stsValue.Schema = "http"
	}

	return factory(sourceCredentials, stsValue), diags
}

// hasAssumeRole reports whether the resolved configuration contains a usable target role TRN.
func hasAssumeRole(c *configModel) bool {
	return c.AssumeRole != nil && isKnownNonEmptyString(c.AssumeRole.AssumeRoleTRN)
}

// isKnownNonEmptyString reports whether a Terraform string is configured with a known,
// non-empty value.
func isKnownNonEmptyString(value types.String) bool {
	return !value.IsNull() && !value.IsUnknown() && value.ValueString() != ""
}

func setProxyDefaultsFromEnvironment(c *configModel) {
	if c.ProxyURL.IsNull() || c.ProxyURL.IsUnknown() {
		if proxyURL := os.Getenv("VOLCENGINE_PROXY_URL"); proxyURL != "" {
			c.ProxyURL = types.StringValue(proxyURL)
		}
	}
	if c.ProxyAuthorization.IsNull() || c.ProxyAuthorization.IsUnknown() {
		if proxyAuthorization := os.Getenv("VOLCENGINE_PROXY_AUTHORIZATION"); proxyAuthorization != "" {
			c.ProxyAuthorization = types.StringValue(proxyAuthorization)
		}
	}
	if c.ProxyIncludeDomains.IsNull() || c.ProxyIncludeDomains.IsUnknown() {
		if domains := splitProxyPatterns(os.Getenv("VOLCENGINE_PROXY_INCLUDE_DOMAINS")); len(domains) > 0 {
			values := make([]attr.Value, 0, len(domains))
			for _, domain := range domains {
				values = append(values, types.StringValue(domain))
			}
			c.ProxyIncludeDomains = types.SetValueMust(types.StringType, values)
		}
	}
	if c.NoProxy.IsNull() || c.NoProxy.IsUnknown() {
		if noProxy := os.Getenv("VOLCENGINE_NO_PROXY"); noProxy != "" {
			c.NoProxy = types.StringValue(noProxy)
			return
		}
		if !hasProxyIncludeDomains(c.ProxyIncludeDomains) {
			if noProxy := firstNonEmptyEnvironmentValue("NO_PROXY", "no_proxy"); noProxy != "" {
				c.NoProxy = types.StringValue(noProxy)
			}
		}
	}
}

func firstNonEmptyEnvironmentValue(names ...string) string {
	for _, name := range names {
		if value := os.Getenv(name); value != "" {
			return value
		}
	}
	return ""
}

func splitProxyPatterns(value string) []string {
	seen := make(map[string]struct{})
	var patterns []string
	for _, pattern := range strings.Split(value, ",") {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			continue
		}
		if _, exists := seen[pattern]; exists {
			continue
		}
		seen[pattern] = struct{}{}
		patterns = append(patterns, pattern)
	}
	return patterns
}

func hasProxyIncludeDomains(value types.Set) bool {
	return !value.IsNull() && !value.IsUnknown() && len(value.Elements()) > 0
}

type proxyAuthorizationRoundTripper struct {
	transport     http.RoundTripper
	authorization string
	proxy         func(*http.Request) (*url.URL, error)
}

func (t *proxyAuthorizationRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if request.URL.Scheme != "http" {
		return t.transport.RoundTrip(request)
	}
	proxyURL, err := t.proxy(request)
	if err != nil {
		return nil, err
	}
	if proxyURL == nil {
		return t.transport.RoundTrip(request)
	}

	requestCopy := request.Clone(request.Context())
	requestCopy.Header = request.Header.Clone()
	requestCopy.Header.Set("Proxy-Authorization", t.authorization)
	return t.transport.RoundTrip(requestCopy)
}

func (t *proxyAuthorizationRoundTripper) CloseIdleConnections() {
	if transport, ok := t.transport.(interface{ CloseIdleConnections() }); ok {
		transport.CloseIdleConnections()
	}
}

func newProviderHTTPClient() (*http.Client, error) {
	defaultTransport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return nil, errors.New("the default HTTP transport cannot be cloned")
	}

	return &http.Client{Transport: defaultTransport.Clone()}, nil
}

func newProxyHTTPClient(rawProxyURL, proxyAuthorization, noProxy string, proxyIncludeDomains []string) (*http.Client, string, error) {
	proxyURL, err := url.Parse(rawProxyURL)
	if err != nil {
		return nil, "", errors.New("proxy_url must be a valid absolute URL")
	}

	proxyURL.Scheme = strings.ToLower(proxyURL.Scheme)
	switch proxyURL.Scheme {
	case "http", "https", "socks5", "socks5h":
	default:
		return nil, "", errors.New("proxy_url must use the http, https, socks5, or socks5h scheme")
	}
	if proxyURL.Host == "" {
		return nil, "", errors.New("proxy_url must include a host")
	}
	if proxyAuthorization != "" && proxyURL.Scheme != "http" && proxyURL.Scheme != "https" {
		return nil, "", errors.New("proxy_authorization requires proxy_url to use the http or https scheme")
	}
	if proxyAuthorization != "" && proxyURL.User != nil {
		return nil, "", errors.New("proxy_url user information and proxy_authorization cannot be configured together")
	}
	if proxyAuthorization != "" {
		if strings.TrimSpace(proxyAuthorization) == "" || !validHTTPHeaderValue(proxyAuthorization) {
			return nil, "", errors.New("proxy_authorization must be a non-empty valid HTTP header value")
		}
	}
	noProxy = strings.TrimSpace(noProxy)
	proxyIncludeDomains, err = normalizeProxyPatterns(proxyIncludeDomains)
	if err != nil {
		return nil, "", err
	}
	if noProxy != "" && len(proxyIncludeDomains) > 0 {
		return nil, "", errors.New("no_proxy and proxy_include_domains cannot be configured together")
	}

	httpClient, err := newProviderHTTPClient()
	if err != nil {
		return nil, "", err
	}
	transport := httpClient.Transport.(*http.Transport)
	transport.Proxy = newProxySelector(proxyURL, noProxy, proxyIncludeDomains)
	if proxyAuthorization != "" {
		transport.GetProxyConnectHeader = nil
		transport.ProxyConnectHeader = make(http.Header)
		transport.ProxyConnectHeader.Set("Proxy-Authorization", proxyAuthorization)
	}

	return httpClient, proxyURL.String(), nil
}

func normalizeProxyPatterns(patterns []string) ([]string, error) {
	normalized := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		if pattern == "" {
			return nil, errors.New("proxy_include_domains must not contain an empty value")
		}
		normalized = append(normalized, pattern)
	}
	return normalized, nil
}

func newProxySelector(proxyURL *url.URL, noProxy string, proxyIncludeDomains []string) func(*http.Request) (*url.URL, error) {
	if noProxy == "" && len(proxyIncludeDomains) == 0 {
		return http.ProxyURL(proxyURL)
	}

	proxyConfig := func(noProxyValue string) func(*url.URL) (*url.URL, error) {
		return (&httpproxy.Config{
			HTTPProxy:  proxyURL.String(),
			HTTPSProxy: proxyURL.String(),
			NoProxy:    noProxyValue,
		}).ProxyFunc()
	}
	if noProxy != "" {
		proxyForURL := proxyConfig(noProxy)
		return func(request *http.Request) (*url.URL, error) {
			return proxyForURL(request.URL)
		}
	}

	proxyForEveryURL := proxyConfig("")
	proxyUnlessIncluded := proxyConfig(strings.Join(proxyIncludeDomains, ","))
	return func(request *http.Request) (*url.URL, error) {
		configuredProxy, err := proxyForEveryURL(request.URL)
		if err != nil || configuredProxy == nil {
			return nil, err
		}
		unmatchedProxy, err := proxyUnlessIncluded(request.URL)
		if err != nil {
			return nil, err
		}
		if unmatchedProxy != nil {
			return nil, nil
		}
		return configuredProxy, nil
	}
}

func validHTTPHeaderValue(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] == '\t' {
			continue
		}
		if value[i] < 0x20 || value[i] == 0x7f {
			return false
		}
	}
	return true
}

func withProxyAuthorization(client *http.Client, proxyAuthorization string, proxy func(*http.Request) (*url.URL, error)) *http.Client {
	if proxyAuthorization == "" {
		return client
	}

	clientCopy := *client
	clientCopy.Transport = &proxyAuthorizationRoundTripper{
		transport:     client.Transport,
		authorization: proxyAuthorization,
		proxy:         proxy,
	}
	return &clientCopy
}

func newProviderData(ctx context.Context, c *configModel) (*providerData, diag.Diagnostics) {
	var diags diag.Diagnostics
	version := fmt.Sprintf("%s/%s (terraform/%s)", common.TerraformProviderName, common.TerraformProviderVersion, c.terraformVersion)
	ctx, logger := baselogging.NewTfLogger(ctx)

	creds, credDiags := buildCredentials(c)
	diags.Append(credDiags...)
	if diags.HasError() {
		return nil, diags
	}

	config := volcengine.NewConfig().
		WithRegion(c.Region.ValueString()).
		WithCredentials(creds).
		WithDisableSSL(c.DisableSSL.ValueBool()).
		WithExtendHttpRequest(func(ctx context.Context, request *http.Request) {
			request.Header.Set("user-agent", version)
		})

	httpClient, err := newProviderHTTPClient()
	if err != nil {
		diags.AddError("Error Creating HTTP Client", err.Error())
		return nil, diags
	}
	config.WithHTTPClient(httpClient)

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
		config.WithEndpoint(fmt.Sprintf("cloudcontrol.%s.volcengineapi.com", c.Region.ValueString()))
	}

	proxyURL := c.ProxyURL.ValueString()
	proxyAuthorization := c.ProxyAuthorization.ValueString()
	noProxy := strings.TrimSpace(c.NoProxy.ValueString())
	proxyIncludeDomains, err := proxyIncludeDomainsFromConfig(c.ProxyIncludeDomains)
	if err != nil {
		diags.AddError("Invalid Proxy Configuration", err.Error())
		return nil, diags
	}
	if proxyURL == "" && proxyAuthorization != "" {
		diags.AddError("Invalid Proxy Configuration", "proxy_url must be configured when proxy_authorization is set")
		return nil, diags
	}
	if proxyURL == "" && len(proxyIncludeDomains) > 0 {
		diags.AddError("Invalid Proxy Configuration", "proxy_url must be configured when proxy_include_domains is set")
		return nil, diags
	}
	var configuredProxy func(*http.Request) (*url.URL, error)
	if proxyURL != "" {
		httpClient, normalizedProxyURL, err := newProxyHTTPClient(proxyURL, proxyAuthorization, noProxy, proxyIncludeDomains)
		if err != nil {
			diags.AddError("Invalid Proxy Configuration", err.Error())
			return nil, diags
		}
		config.WithHTTPProxy(normalizedProxyURL).
			WithHTTPSProxy(normalizedProxyURL).
			WithHTTPClient(httpClient)
		configuredProxy = httpClient.Transport.(*http.Transport).Proxy
	}

	sess, err := session.NewSession(config)
	if err != nil {
		diags.AddError(err.Error(), err.Error())
		return nil, diags
	}

	cloudcontrolClient := cloudcontrol.New(sess)
	if configuredProxy != nil {
		transport, ok := cloudcontrolClient.Config.HTTPClient.Transport.(*http.Transport)
		if !ok {
			diags.AddError("Error Configuring Proxy", fmt.Sprintf("Cloud Control HTTP transport has type %T, want *http.Transport", cloudcontrolClient.Config.HTTPClient.Transport))
			return nil, diags
		}
		transport.Proxy = configuredProxy
		cloudcontrolClient.Config.HTTPClient = withProxyAuthorization(cloudcontrolClient.Config.HTTPClient, proxyAuthorization, configuredProxy)
	}

	providerData := &providerData{
		ccAPIClient: cloudcontrolClient,
		logger:      logger,
		region:      c.Region.String(),
	}

	return providerData, diags
}

func proxyIncludeDomainsFromConfig(value types.Set) ([]string, error) {
	if value.IsNull() || value.IsUnknown() {
		return nil, nil
	}

	domains := make([]string, 0, len(value.Elements()))
	for _, element := range value.Elements() {
		domain, ok := element.(types.String)
		if !ok || domain.IsNull() || domain.IsUnknown() {
			return nil, errors.New("proxy_include_domains must contain only known string values")
		}
		domains = append(domains, domain.ValueString())
	}
	return normalizeProxyPatterns(domains)
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
