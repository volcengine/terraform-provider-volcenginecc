// Copyright (c) 2026 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	frameworkprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func clearAssumeRoleProxyEnvironment(t *testing.T) {
	t.Helper()
	for _, name := range []string{"VOLCENGINE_ASSUME_ROLE_PROXY_URL", "VOLCENGINE_ASSUME_ROLE_NO_PROXY", "HTTP_PROXY", "http_proxy", "HTTPS_PROXY", "https_proxy", "NO_PROXY", "no_proxy"} {
		t.Setenv(name, "")
	}
}

func TestAssumeRoleProxyPrecedence(t *testing.T) {
	tests := []struct {
		name  string
		value types.String
		env   map[string]string
		want  string
	}{
		{"explicit", types.StringValue("http://explicit:8080"), map[string]string{"VOLCENGINE_ASSUME_ROLE_PROXY_URL": "http://dedicated:8080", "HTTP_PROXY": "http://standard:8080"}, "http://explicit:8080"},
		{"dedicated environment", types.StringNull(), map[string]string{"VOLCENGINE_ASSUME_ROLE_PROXY_URL": "http://dedicated:8080", "HTTP_PROXY": "http://standard:8080"}, "http://dedicated:8080"},
		{"HTTP_PROXY ignored", types.StringNull(), map[string]string{"HTTP_PROXY": "http://standard:8080", "http_proxy": "http://lower:8080"}, ""},
		{"lowercase", types.StringNull(), map[string]string{"https_proxy": "http://lower:8080"}, "http://lower:8080"},
		{"explicit empty", types.StringValue(""), map[string]string{"HTTP_PROXY": "http://standard:8080", "VOLCENGINE_ASSUME_ROLE_PROXY_URL": "http://dedicated:8080"}, ""},
		{"unset", types.StringNull(), nil, ""},
		{"HTTPS_PROXY", types.StringNull(), map[string]string{"HTTPS_PROXY": "http://https:8080", "HTTP_PROXY": "http://http:8080", "VOLCENGINE_PROXY_URL": "http://cloudcontrol:8080"}, "http://https:8080"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearAssumeRoleProxyEnvironment(t)
			for name, value := range tt.env {
				t.Setenv(name, value)
			}
			c := testConfigModel()
			c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: tt.value}
			called := false
			_, diags := buildCredentialsWithFactory(c, func(_ *credentials.Credentials, value credentials.StsValue) *credentials.Credentials {
				called = true
				if value.ProxyURL != tt.want {
					t.Errorf("proxy = %q, want %q", value.ProxyURL, tt.want)
				}
				return credentials.NewStaticCredentials("target-ak", "target-sk", "target-token")
			})
			if diags.HasError() || !called {
				t.Fatalf("credentials not constructed: %v", diags)
			}
		})
	}
}

func TestAssumeRoleProxyValidation(t *testing.T) {
	for _, value := range []string{"http://user:secret@%zz", "http://", "proxy:8080", "ftp://user:secret@proxy:8080"} {
		t.Run(value, func(t *testing.T) {
			c := testConfigModel()
			c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: types.StringValue(value)}
			_, diags := buildCredentialsWithFactory(c, func(_ *credentials.Credentials, _ credentials.StsValue) *credentials.Credentials {
				t.Fatal("invalid proxy reached credential factory")
				return nil
			})
			if !diags.HasError() {
				t.Fatal("expected invalid proxy diagnostic")
			}
			if strings.Contains(fmt.Sprint(diags), "secret") {
				t.Fatal("diagnostic exposed proxy password")
			}
		})
	}
	for _, value := range []string{"http://user:password@proxy:8080", "https://proxy:8443", "socks5://proxy:1080", "socks5h://proxy:1080"} {
		if _, err := parseAssumeRoleProxyURL(value); err != nil {
			t.Errorf("valid proxy rejected: %v", err)
		}
	}
	// A malformed ambient proxy must not prevent configurations without AssumeRole.
	t.Setenv("HTTP_PROXY", "invalid")
	if _, diags := buildCredentials(testConfigModel()); diags.HasError() {
		t.Fatal(diags)
	}
}

func writeSTSResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, `{"ResponseMetadata":{"RequestId":"local-test"},"Result":{"Credentials":{"AccessKeyId":"target-ak","SecretAccessKey":"target-sk","SessionToken":"target-token"}}}`)
}

// Exercises framework decoding, Configure, the source wrapper and the real SDK
// against a local forward proxy, with no cloud credentials or cloud requests.
func TestConfigureAssumeRoleThroughProxy(t *testing.T) {
	for _, mode := range []string{"terraform", "dedicated environment", "environment role"} {
		t.Run(mode, func(t *testing.T) {
			clearAssumeRoleProxyEnvironment(t)
			var requests atomic.Int32
			proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				if r.URL.Host != "sts.invalid" || r.URL.Query().Get("Action") != "AssumeRole" {
					t.Errorf("unexpected STS request: %s", r.URL)
				}
				if !strings.Contains(r.Header.Get("Authorization"), "source-ak") {
					t.Error("missing source signature")
				}
				if r.Header.Get("X-Security-Token") != "source-token" {
					t.Error("missing source session token")
				}
				if r.Header.Get("Proxy-Authorization") != "Basic dXNlcjpwYXNz" {
					t.Error("missing proxy Basic authentication")
				}
				writeSTSResponse(w)
			}))
			defer proxy.Close()
			proxyURL, _ := url.Parse(proxy.URL)
			proxyURL.User = url.UserPassword("user", "pass")
			role := map[string]interface{}{"assume_role_trn": "trn:iam::222222222222:role/test"}
			switch mode {
			case "terraform":
				role["proxy_url"] = proxyURL.String()
				t.Setenv("HTTP_PROXY", "http://127.0.0.1:1")
			case "dedicated environment", "environment role":
				t.Setenv("VOLCENGINE_ASSUME_ROLE_PROXY_URL", proxyURL.String())
				t.Setenv("HTTP_PROXY", "http://127.0.0.1:1")
			}
			// Cloud Control bypass settings do not affect the separate STS route.
			t.Setenv("VOLCENGINE_NO_PROXY", "*")
			input := map[string]interface{}{
				"access_key": "source-ak", "secret_key": "source-sk", "session_token": "source-token", "region": "cn-beijing", "disable_ssl": true,
				"assume_role": role, "endpoints": map[string]interface{}{"sts": "sts.invalid"},
				"proxy_url": "http://127.0.0.1:1",
			}
			if mode == "environment role" {
				delete(input, "assume_role")
				t.Setenv("VOLCENGINE_ASSUME_ROLE_TRN", "trn:iam::222222222222:role/test")
			}
			p := configureAssumeRoleForTest(t, input)
			creds := p.providerData.ccAPIClient.Config.Credentials
			for i := 0; i < 2; i++ {
				got, err := creds.Get()
				if err != nil {
					t.Fatal(err)
				}
				if got.AccessKeyID != "target-ak" || got.SessionToken != "target-token" {
					t.Fatalf("unexpected credentials: %#v", got)
				}
			}
			if requests.Load() != 1 {
				t.Fatalf("cached credentials made %d requests", requests.Load())
			}
			t.Setenv("HTTP_PROXY", "http://127.0.0.1:1")
			t.Setenv("VOLCENGINE_ASSUME_ROLE_PROXY_URL", "http://127.0.0.1:1")
			creds.Expire()
			if _, err := creds.Get(); err != nil {
				t.Fatalf("refresh lost proxy snapshot: %v", err)
			}
			if requests.Load() != 2 {
				t.Fatalf("refresh made %d requests", requests.Load())
			}
		})
	}
}

func configureAssumeRoleForTest(t *testing.T, input map[string]interface{}) *VolcengineCCProvider {
	t.Helper()
	ctx := context.Background()
	p := &VolcengineCCProvider{}
	var schemaResponse frameworkprovider.SchemaResponse
	p.Schema(ctx, frameworkprovider.SchemaRequest{}, &schemaResponse)
	role := schemaResponse.Schema.Attributes["assume_role"].(providerschema.SingleNestedAttribute)
	proxy := role.Attributes["proxy_url"].(providerschema.StringAttribute)
	if !proxy.Optional || !proxy.Sensitive {
		t.Fatal("proxy must be optional and sensitive")
	}
	rawJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := tftypes.ValueFromJSON(rawJSON, schemaResponse.Schema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatal(err)
	}
	var response frameworkprovider.ConfigureResponse
	p.Configure(ctx, frameworkprovider.ConfigureRequest{Config: tfsdk.Config{Raw: raw, Schema: schemaResponse.Schema}}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	return p
}

func TestAssumeRoleHTTPSProxyForHTTPS(t *testing.T) {
	clearAssumeRoleProxyEnvironment(t)
	var connects atomic.Int32
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connects.Add(1)
		if r.Method != http.MethodConnect || r.Host != "sts.volcengineapi.com:443" {
			t.Errorf("unexpected HTTPS proxy request: %s %s", r.Method, r.Host)
		}
		// Stop before TLS: proves the default HTTPS STS endpoint uses HTTPS_PROXY
		// without trusting a test certificate or contacting any external service.
		w.WriteHeader(http.StatusForbidden)
	}))
	defer proxy.Close()
	t.Setenv("HTTPS_PROXY", proxy.URL)
	c := testConfigModel()
	c.AccessKey, c.SecretKey = types.StringValue("source-ak"), types.StringValue("source-sk")
	c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test")}
	creds, diags := buildCredentialsWithFactory(c, func(source *credentials.Credentials, value credentials.StsValue) *credentials.Credentials {
		value.MaxRetries, value.RetryInterval = 1, time.Millisecond
		return newAssumeRoleCredentials(source, value)
	})
	if diags.HasError() {
		t.Fatal(diags)
	}
	if _, err := creds.Get(); err == nil {
		t.Fatal("proxy denial must fail credential retrieval")
	}
	if connects.Load() == 0 {
		t.Fatal("HTTPS STS bypassed HTTPS_PROXY")
	}
}

func TestAssumeRoleProxyIsolationAndDirect(t *testing.T) {
	clearAssumeRoleProxyEnvironment(t)
	var directCalls, proxyCalls atomic.Int32
	direct := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { directCalls.Add(1); writeSTSResponse(w) }))
	defer direct.Close()
	proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { proxyCalls.Add(1); writeSTSResponse(w) }))
	defer proxy.Close()
	t.Setenv("HTTP_PROXY", proxy.URL)
	makeCreds := func(value types.String) *credentials.Credentials {
		c := testConfigModel()
		c.AccessKey, c.SecretKey, c.DisableSSL = types.StringValue("source-ak"), types.StringValue("source-sk"), types.BoolValue(true)
		c.Endpoints.STS = types.StringValue(strings.TrimPrefix(direct.URL, "http://"))
		c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: value}
		creds, diags := buildCredentials(c)
		if diags.HasError() {
			t.Fatal(diags)
		}
		return creds
	}
	proxied, unproxied := makeCreds(types.StringValue(proxy.URL)), makeCreds(types.StringValue(""))
	for _, creds := range []*credentials.Credentials{proxied, unproxied, proxied, unproxied} {
		creds.Expire()
		if _, err := creds.Get(); err != nil {
			t.Fatal(err)
		}
	}
	if directCalls.Load() != 2 || proxyCalls.Load() != 2 {
		t.Fatalf("client isolation failed: direct=%d proxy=%d", directCalls.Load(), proxyCalls.Load())
	}
}

// Reproduces real STS endpoints that negotiate HTTP/2 through CONNECT.
func TestAssumeRoleProxyHTTPSNegotiation(t *testing.T) {
	clearAssumeRoleProxyEnvironment(t)
	target := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			t.Errorf("STS protocol = %s, want HTTP/2", r.Proto)
		}
		writeSTSResponse(w)
	}))
	target.EnableHTTP2 = true
	target.StartTLS()
	defer target.Close()
	caFile := filepath.Join(t.TempDir(), "sts-ca.pem")
	if err := os.WriteFile(caFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: target.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	proxy, observations := newAuthenticatedConnectProxy(t, "Basic dXNlcjpwYXNz")
	defer proxy.Close()
	proxyURL, _ := url.Parse(proxy.URL)
	proxyURL.User = url.UserPassword("user", "pass")
	c := testConfigModel()
	c.AccessKey, c.SecretKey = types.StringValue("source-ak"), types.StringValue("source-sk")
	c.Endpoints.STS = types.StringValue(strings.TrimPrefix(target.URL, "https://"))
	c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: types.StringValue(proxyURL.String())}
	creds, diags := buildCredentialsWithFactory(c, func(source *credentials.Credentials, value credentials.StsValue) *credentials.Credentials {
		value.CAFile = caFile
		value.MaxRetries, value.RetryInterval = 1, time.Millisecond
		return newAssumeRoleCredentials(source, value)
	})
	if diags.HasError() {
		t.Fatal(diags)
	}
	got, err := creds.Get()
	if err != nil {
		t.Fatal(err)
	}
	if got.AccessKeyID != "target-ak" {
		t.Fatal("unexpected target identity")
	}
	observation := receiveProxyObservation(t, observations)
	if observation.method != http.MethodConnect {
		t.Fatal("expected CONNECT")
	}
}

func TestConfigureAssumeRoleNoProxyAndRefresh(t *testing.T) {
	for _, tt := range []struct {
		name, noProxy string
		direct        bool
	}{
		{"explicit bypass", "*", true},
		{"empty overrides environment", "", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			clearAssumeRoleProxyEnvironment(t)
			var directCalls, proxyCalls atomic.Int32
			target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { directCalls.Add(1); writeSTSResponse(w) }))
			defer target.Close()
			proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { proxyCalls.Add(1); writeSTSResponse(w) }))
			defer proxy.Close()
			t.Setenv("NO_PROXY", "*")
			p := configureAssumeRoleForTest(t, map[string]interface{}{
				"access_key": "source-ak", "secret_key": "source-sk", "region": "cn-beijing", "disable_ssl": true,
				"endpoints":   map[string]interface{}{"sts": strings.TrimPrefix(target.URL, "http://")},
				"assume_role": map[string]interface{}{"assume_role_trn": "trn:iam::222222222222:role/test", "proxy_url": proxy.URL, "no_proxy": tt.noProxy},
			})
			creds := p.providerData.ccAPIClient.Config.Credentials
			for i := 0; i < 2; i++ {
				t.Setenv("NO_PROXY", "") // Already-constructed credentials retain their route.
				creds.Expire()
				if _, err := creds.Get(); err != nil {
					t.Fatal(err)
				}
			}
			if tt.direct {
				if directCalls.Load() != 2 || proxyCalls.Load() != 0 {
					t.Fatalf("direct=%d proxy=%d", directCalls.Load(), proxyCalls.Load())
				}
			} else if directCalls.Load() != 0 || proxyCalls.Load() != 2 {
				t.Fatalf("direct=%d proxy=%d", directCalls.Load(), proxyCalls.Load())
			}
		})
	}
}

func TestAssumeRoleProxyRouting(t *testing.T) {
	const explicit = "http://explicit-proxy.example:8080"
	const standard = "http://https-proxy.example:8080"
	tests := []struct {
		name, host     string
		disableSSL     bool
		proxy, noProxy types.String
		env            map[string]string
		want           string
	}{
		{name: "HTTPS prefers HTTPS_PROXY", host: "sts.example.com", env: map[string]string{"HTTPS_PROXY": standard, "HTTP_PROXY": "http://http-proxy.example:8080"}, want: standard},
		{name: "HTTP environment ignored", host: "sts.example.com", disableSSL: true, env: map[string]string{"HTTPS_PROXY": standard, "HTTP_PROXY": explicit}},
		{name: "HTTP_PROXY ignored", host: "sts.example.com", env: map[string]string{"HTTP_PROXY": explicit}},
		{name: "lowercase HTTPS and bypass", host: "sts.example.com", env: map[string]string{"https_proxy": standard, "no_proxy": ".example.com"}},
		{name: "environment wildcard", host: "sts.example.com", env: map[string]string{"HTTPS_PROXY": standard, "NO_PROXY": "*"}},
		{name: "explicit URL honors environment bypass", host: "sts.example.com", proxy: types.StringValue(explicit), env: map[string]string{"NO_PROXY": ".example.com"}},
		{name: "explicit no_proxy overrides environment", host: "sts.example.com", proxy: types.StringValue(explicit), noProxy: types.StringValue("unrelated.example"), env: map[string]string{"NO_PROXY": "*"}, want: explicit},
		{name: "explicit empty bypass", host: "sts.example.com", proxy: types.StringValue(explicit), noProxy: types.StringValue(""), env: map[string]string{"NO_PROXY": "*"}, want: explicit},
		{name: "dedicated bypass", host: "sts.example.com", env: map[string]string{"HTTPS_PROXY": standard, "VOLCENGINE_ASSUME_ROLE_NO_PROXY": "sts.example.com", "NO_PROXY": "unrelated.example"}},
		{name: "dedicated URL", host: "sts.example.com", env: map[string]string{"VOLCENGINE_ASSUME_ROLE_PROXY_URL": explicit, "HTTPS_PROXY": standard}, want: explicit},
		{name: "CIDR", host: "10.2.3.4", env: map[string]string{"HTTPS_PROXY": standard}, noProxy: types.StringValue("10.0.0.0/8")},
		{name: "IPv6 CIDR", host: "[2001:db8::1]", env: map[string]string{"HTTPS_PROXY": standard}, noProxy: types.StringValue("2001:db8::/32")},
		{name: "port", host: "sts.example.com:8443", env: map[string]string{"HTTPS_PROXY": standard}, noProxy: types.StringValue("sts.example.com:8443")},
		{name: "other port", host: "sts.example.com:443", env: map[string]string{"HTTPS_PROXY": standard}, noProxy: types.StringValue("sts.example.com:8443"), want: standard},
		{name: "CC bypass is independent", host: "sts.example.com", env: map[string]string{"HTTPS_PROXY": standard, "VOLCENGINE_NO_PROXY": "*"}, want: standard},
		{name: "explicit direct", host: "sts.example.com", proxy: types.StringValue(""), env: map[string]string{"HTTPS_PROXY": standard}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearProxyEnvironment(t)
			clearAssumeRoleProxyEnvironment(t)
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			c := testConfigModel()
			c.DisableSSL = types.BoolValue(tt.disableSSL)
			c.Endpoints.STS = types.StringValue(tt.host)
			c.NoProxy = types.StringValue("*") // Top-level bypass must never leak into STS.
			c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: tt.proxy, NoProxy: tt.noProxy}
			var got string
			_, diags := buildCredentialsWithFactory(c, func(_ *credentials.Credentials, value credentials.StsValue) *credentials.Credentials {
				got = value.ProxyURL
				return credentials.NewStaticCredentials("ak", "sk", "")
			})
			if diags.HasError() {
				t.Fatal(diags)
			}
			if got != tt.want {
				t.Fatalf("proxy=%q, want %q", got, tt.want)
			}
		})
	}
}

func TestAssumeRoleHTTPSNoProxyRequest(t *testing.T) {
	clearAssumeRoleProxyEnvironment(t)
	var requests atomic.Int32
	target := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { requests.Add(1); writeSTSResponse(w) }))
	defer target.Close()
	caFile := filepath.Join(t.TempDir(), "sts-ca.pem")
	if err := os.WriteFile(caFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: target.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("HTTPS_PROXY", "http://127.0.0.1:1")
	c := testConfigModel()
	c.AccessKey = types.StringValue("source-ak")
	c.SecretKey = types.StringValue("source-sk")
	c.Endpoints.STS = types.StringValue(strings.TrimPrefix(target.URL, "https://"))
	c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test"), ProxyURL: types.StringValue("http://127.0.0.1:1"), NoProxy: types.StringValue("*")}
	creds, diags := buildCredentialsWithFactory(c, func(source *credentials.Credentials, value credentials.StsValue) *credentials.Credentials {
		value.CAFile = caFile
		return newAssumeRoleCredentials(source, value)
	})
	if diags.HasError() {
		t.Fatal(diags)
	}
	if _, err := creds.Get(); err != nil {
		t.Fatal(err)
	}
	if requests.Load() != 1 {
		t.Fatalf("direct STS requests=%d, want 1", requests.Load())
	}
}
