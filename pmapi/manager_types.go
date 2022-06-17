package pmapi

import (
	"context"
	"net/http"
	"time"

	"github.com/ProtonMail/gopenpgp/crypto"
	"github.com/sirupsen/logrus"
)

type Manager interface {
	NewClient(string, string, string, time.Time) Client
	NewClientWithRefresh(context.Context, string, string) (Client, *AuthRefresh, error)
	NewClientWithLogin(context.Context, string, []byte) (Client, *Auth, error)

	DownloadAndVerify(kr *crypto.KeyRing, url, sig string) ([]byte, error)
	ReportBug(context.Context, ReportBugReq) error
	SendSimpleMetric(context.Context, string, string, string) error

	SetLogging(logger *logrus.Entry, verbose bool)
	SetTransport(http.RoundTripper)
	SetCookieJar(http.CookieJar)
	SetRetryCount(int)
	AddConnectionObserver(ConnectionObserver)

	AllowProxy()
	DisallowProxy()
}
