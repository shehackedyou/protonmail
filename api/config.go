package api

import (
	"os"
	"runtime"
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

type Config struct {
	// HostURL is the base URL of API.
	HostURL string

	// AppVersion sets version to headers of each request.
	AppVersion string

	// UserAgent sets user agent to headers of each request.
	// Used only if GetUserAgent is not set.
	UserAgent string

	// GetUserAgent is dynamic version of UserAgent.
	// Overrides UserAgent.
	GetUserAgent func() string

	// UpgradeApplicationHandler is used to notify when there is a force upgrade.
	UpgradeApplicationHandler func()

	// TLSIssueHandler is used to notify when there is a TLS issue.
	TLSIssueHandler func()
}

func NewConfig(appVersionName, appVersion string) Config {
	return Config{
		HostURL:    getRootURL(),
		AppVersion: getAPIOS() + strings.Title(appVersionName) + "_" + appVersion,
	}
}

func (c *Config) getUserAgent() string {
	if c.GetUserAgent == nil {
		return c.UserAgent
	}
	return c.GetUserAgent()
}

// getAPIOS returns actual operating system.
func getAPIOS() string {
	switch os := runtime.GOOS; os {
	case "darwin": // nolint: goconst
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	}
	return "Linux"
}
