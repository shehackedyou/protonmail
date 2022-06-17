// +build !build_qa

package pmapi

import (
	"net/http"
)

func getRootURL() string {
	return "https://api.protonmail.ch"
}

func newProxyDialerAndTransport(cfg Config) (*ProxyTLSDialer, http.RoundTripper) {
	basicDialer := NewBasicTLSDialer(cfg)
	pinningDialer := NewPinningTLSDialer(cfg, basicDialer)
	proxyDialer := NewProxyTLSDialer(cfg, pinningDialer)
	return proxyDialer, CreateTransportWithDialer(proxyDialer)
}
