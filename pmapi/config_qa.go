// +build build_qa

package pmapi

import (
	"crypto/tls"
	"net/http"
	"os"
	"strings"
)

func getRootURL() string {
	// This config allows to dynamically change ROOT URL.
	url := os.Getenv("PMAPI_ROOT_URL")
	if strings.HasPrefix(url, "http") {
		return url
	}
	if url != "" {
		return "https://" + url
	}
	return "https://api.protonmail.ch"
}

func newProxyDialerAndTransport(cfg Config) (*ProxyTLSDialer, http.RoundTripper) {
	transport := CreateTransportWithDialer(NewBasicTLSDialer(cfg))

	// TLS certificate of testing environment might be self-signed.
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return nil, transport
}
