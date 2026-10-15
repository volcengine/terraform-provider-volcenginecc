// Copyright (c) 2026 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Observe the selector installed on the final SDK client, after session setup.
func TestCloudControlEnvironmentProxyRouting(t *testing.T) {
	const httpProxy = "http://http-proxy.example:8080"
	const httpsProxy = "http://https-proxy.example:8080"
	tests := []struct{ name, target, noProxy, want string }{
		{"HTTPS selection", "https://api.example.com/", "", httpsProxy},
		{"HTTP environment ignored", "http://api.example.com/", "", ""},
		{"domain bypass", "https://api.example.com/", ".example.com", ""},
		{"unmatched domain", "https://api.other.com/", ".example.com", httpsProxy},
		{"CIDR bypass", "https://10.12.0.9/", "10.0.0.0/8", ""},
		{"IPv6 bypass", "https://[2001:db8::1]/", "2001:db8::/32", ""},
		{"port bypass", "https://api.example.com:8443/", "api.example.com:8443", ""},
		{"wildcard bypass", "https://api.example.com/", "*", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearProxyEnvironment(t)
			clearAssumeRoleProxyEnvironment(t)
			t.Setenv("HTTP_PROXY", httpProxy)
			t.Setenv("HTTPS_PROXY", httpsProxy)
			t.Setenv("NO_PROXY", tt.noProxy)
			c := testConfigModel()
			c.AccessKey = types.StringValue("test-ak")
			c.SecretKey = types.StringValue("test-sk")
			setProxyDefaultsFromEnvironment(c)
			data, diags := newProviderData(context.Background(), c)
			if diags.HasError() {
				t.Fatal(diags)
			}
			transport := data.ccAPIClient.Config.HTTPClient.Transport.(*http.Transport)
			defer transport.CloseIdleConnections()
			target, _ := url.Parse(tt.target)
			got, err := transport.Proxy(&http.Request{URL: target})
			if err != nil {
				t.Fatal(err)
			}
			actual := ""
			if got != nil {
				actual = got.String()
			}
			if actual != tt.want {
				t.Fatalf("proxy=%q, want %q", actual, tt.want)
			}
		})
	}
}

func TestCloudControlProxyPrecedenceAndSnapshot(t *testing.T) {
	const httpProxy = "http://http-proxy.example:8080"
	const httpsProxy = "http://https-proxy.example:8080"
	tests := []struct {
		name              string
		env               map[string]string
		explicit, noProxy types.String
		want              string
	}{
		{name: "HTTP_PROXY ignored", env: map[string]string{"HTTP_PROXY": httpProxy}},
		{name: "lowercase HTTPS", env: map[string]string{"https_proxy": httpsProxy, "http_proxy": httpProxy}, want: httpsProxy},
		{name: "dedicated proxy honors NO_PROXY", env: map[string]string{"VOLCENGINE_PROXY_URL": httpProxy, "NO_PROXY": "*"}},
		{name: "explicit proxy wins", env: map[string]string{"VOLCENGINE_PROXY_URL": httpProxy, "HTTPS_PROXY": httpProxy}, explicit: types.StringValue(httpsProxy), want: httpsProxy},
		{name: "explicit no_proxy wins", env: map[string]string{"HTTPS_PROXY": httpsProxy, "NO_PROXY": "*"}, noProxy: types.StringValue("unrelated.example"), want: httpsProxy},
		{name: "empty no_proxy clears inherited rule", env: map[string]string{"HTTPS_PROXY": httpsProxy, "NO_PROXY": "*"}, noProxy: types.StringValue(""), want: httpsProxy},
		{name: "dedicated bypass wins", env: map[string]string{"HTTPS_PROXY": httpsProxy, "VOLCENGINE_NO_PROXY": "*", "NO_PROXY": "unrelated.example"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearProxyEnvironment(t)
			clearAssumeRoleProxyEnvironment(t)
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			c := testConfigModel()
			c.ProxyURL = tt.explicit
			c.NoProxy = tt.noProxy
			setProxyDefaultsFromEnvironment(c)
			data, diags := newProviderData(context.Background(), c)
			if diags.HasError() {
				t.Fatal(diags)
			}
			transport := data.ccAPIClient.Config.HTTPClient.Transport.(*http.Transport)
			defer transport.CloseIdleConnections()
			t.Setenv("HTTPS_PROXY", "http://changed.example:8080")
			t.Setenv("NO_PROXY", "*")
			target, _ := url.Parse("https://api.example.com/")
			selected, err := transport.Proxy(&http.Request{URL: target})
			if err != nil {
				t.Fatal(err)
			}
			actual := ""
			if selected != nil {
				actual = selected.String()
			}
			if actual != tt.want {
				t.Fatalf("proxy=%q, want %q", actual, tt.want)
			}
		})
	}
}

func TestProxyEnvironmentValidationIsRedacted(t *testing.T) {
	clearProxyEnvironment(t)
	clearAssumeRoleProxyEnvironment(t)
	t.Setenv("HTTPS_PROXY", "http://user:secret@%zz")
	c := testConfigModel()
	_, diags := newProviderData(context.Background(), c)
	if !diags.HasError() {
		t.Fatal("invalid environment proxy must fail configuration")
	}
	for _, d := range diags {
		if strings.Contains(d.Detail(), "secret") {
			t.Fatal("environment proxy diagnostic exposed credentials")
		}
	}
	c.AssumeRole = &AssumeRoleData{AssumeRoleTRN: types.StringValue("trn:iam::222222222222:role/test")}
	_, diags = buildCredentials(c)
	if !diags.HasError() {
		t.Fatal("invalid STS environment proxy must not silently connect directly")
	}
	for _, d := range diags {
		if strings.Contains(d.Detail(), "secret") {
			t.Fatal("STS diagnostic exposed proxy credentials")
		}
	}
	// HTTP_PROXY is no longer an input, including malformed values.
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("HTTP_PROXY", "http://user:secret@%zz")
	if _, diags := newProviderData(context.Background(), c); diags.HasError() {
		t.Fatal(diags)
	}
}

func TestCloudControlHTTPSProxyBypassRequest(t *testing.T) {
	clearProxyEnvironment(t)
	clearAssumeRoleProxyEnvironment(t)
	var requests atomic.Int32
	target := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		cloudControlHandler(make(chan string, 1))(w, r)
	}))
	defer target.Close()
	if err := target.Certificate().VerifyHostname("example.com"); err != nil {
		t.Fatal(err)
	}
	endpoint, _ := url.Parse(target.URL)
	alias := "example.com:" + endpoint.Port()
	base := target.Client().Transport.(*http.Transport).Clone()
	dialer := &net.Dialer{}
	base.DialContext = func(ctx context.Context, network, address string) (net.Conn, error) {
		if address == alias {
			address = endpoint.Host
		}
		return dialer.DialContext(ctx, network, address)
	}
	previous := http.DefaultTransport
	http.DefaultTransport = base
	t.Cleanup(func() { http.DefaultTransport = previous; base.CloseIdleConnections() })
	t.Setenv("HTTPS_PROXY", "http://127.0.0.1:1")
	t.Setenv("NO_PROXY", "example.com")
	c := testConfigModel()
	c.AccessKey = types.StringValue("source-ak")
	c.SecretKey = types.StringValue("source-sk")
	c.Endpoints.CloudControlAPI = types.StringValue(alias)
	setProxyDefaultsFromEnvironment(c)
	data, diags := newProviderData(context.Background(), c)
	if diags.HasError() {
		t.Fatal(diags)
	}
	defer data.ccAPIClient.Config.HTTPClient.CloseIdleConnections()
	callCloudControl(t, data.ccAPIClient)
	if requests.Load() != 1 {
		t.Fatalf("direct HTTPS requests=%d, want 1", requests.Load())
	}
}
