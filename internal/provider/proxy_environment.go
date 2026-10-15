// Copyright (c) 2026 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/http/httpproxy"
)

// Snapshot the HTTPS proxy per configured provider instead of using the
// process-wide cache in http.ProxyFromEnvironment. HTTP_PROXY is deliberately
// ignored; plain HTTP requests require an explicit provider proxy.
func newEnvironmentProxySelector(noProxy string) func(*http.Request) (*url.URL, error) {
	config := &httpproxy.Config{
		HTTPSProxy: firstNonEmptyEnvironmentValue("HTTPS_PROXY", "https_proxy"),
		NoProxy:    noProxy,
	}
	var proxyErr error
	if raw := config.HTTPSProxy; raw != "" {
		if !strings.Contains(raw, "://") {
			raw = "http://" + raw
		}
		parsed, err := url.Parse(raw)
		if err != nil || parsed.Hostname() == "" {
			proxyErr = errors.New("HTTPS_PROXY must be a valid proxy URL with a host")
		} else {
			parsed.Scheme = strings.ToLower(parsed.Scheme)
			switch parsed.Scheme {
			case "http", "https", "socks5", "socks5h":
				config.HTTPSProxy = parsed.String()
			default:
				proxyErr = errors.New("HTTPS_PROXY must use the http, https, socks5, or socks5h scheme")
			}
		}
	}
	selectProxy := config.ProxyFunc()
	return func(request *http.Request) (*url.URL, error) {
		if request.URL.Scheme == "https" && proxyErr != nil {
			return nil, proxyErr
		}
		return selectProxy(request.URL)
	}
}
