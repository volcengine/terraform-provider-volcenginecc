// Copyright (c) 2026 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"errors"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// assumeRoleProxyURL snapshots the proxy for the lifetime of these credentials.
// HTTP_PROXY also applies to HTTPS STS endpoints: the SDK uses CONNECT for them.
func assumeRoleProxyURL(value types.String) (string, error) {
	raw := value.ValueString()
	if value.IsNull() || value.IsUnknown() {
		raw = firstNonEmptyEnvironmentValue("VOLCENGINE_ASSUME_ROLE_PROXY_URL", "HTTP_PROXY", "http_proxy")
	}
	if raw == "" {
		return "", nil
	}
	parsed, err := url.Parse(raw)
	if err != nil || parsed.Hostname() == "" {
		// URL parser errors can contain proxy credentials. Never echo the input.
		return "", errors.New("assume_role.proxy_url must be a valid absolute proxy URL with a host")
	}
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	switch parsed.Scheme {
	case "http", "https", "socks5", "socks5h":
		return parsed.String(), nil
	default:
		return "", errors.New("assume_role.proxy_url must use the http, https, socks5, or socks5h scheme")
	}
}
