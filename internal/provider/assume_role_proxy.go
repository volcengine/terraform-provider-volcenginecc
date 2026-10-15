// Copyright (c) 2026 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// STS has a fixed endpoint for the lifetime of these credentials. Resolve both
// proxy and bypass rules once for that endpoint; refreshes reuse the snapshot.
// An empty SDK ProxyURL keeps its default direct STS transport.
func assumeRoleProxyURL(value, noProxy types.String, scheme, host string) (string, error) {
	bypass := noProxy.ValueString()
	if noProxy.IsNull() || noProxy.IsUnknown() {
		bypass = firstNonEmptyEnvironmentValue("VOLCENGINE_ASSUME_ROLE_NO_PROXY", "NO_PROXY", "no_proxy")
	}
	bypass = strings.TrimSpace(bypass)
	raw := value.ValueString()
	explicit := !value.IsNull() && !value.IsUnknown()
	if !explicit {
		raw = firstNonEmptyEnvironmentValue("VOLCENGINE_ASSUME_ROLE_PROXY_URL")
		explicit = raw != ""
	}
	selectProxy := newEnvironmentProxySelector(bypass)
	if explicit {
		if raw == "" {
			return "", nil
		}
		parsed, err := parseAssumeRoleProxyURL(raw)
		if err != nil {
			return "", err
		}
		selectProxy = newProxySelector(parsed, bypass, nil)
	}
	selected, err := selectProxy(&http.Request{URL: &url.URL{Scheme: scheme, Host: host}})
	if err != nil {
		// Do not include an environment proxy URL or its credentials in diagnostics.
		return "", errors.New("unable to resolve the AssumeRole proxy for the STS endpoint")
	}
	if selected == nil {
		return "", nil
	}
	return selected.String(), nil
}

func parseAssumeRoleProxyURL(raw string) (*url.URL, error) {
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Hostname() == "" {
		return nil, errors.New("assume_role.proxy_url must be a valid absolute proxy URL with a host")
	}
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	switch parsed.Scheme {
	case "http", "https", "socks5", "socks5h":
		return parsed, nil
	default:
		return nil, errors.New("assume_role.proxy_url must use the http, https, socks5, or socks5h scheme")
	}
}
