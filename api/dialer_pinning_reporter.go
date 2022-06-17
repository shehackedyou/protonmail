package api

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

type sentReport struct {
	r tlsReport
	t time.Time
}

type tlsReporter struct {
	cfg         Config
	trustedPins []string
	sentReports []sentReport
}

func newTLSReporter(cfg Config, trustedPins []string) *tlsReporter {
	return &tlsReporter{
		cfg:         cfg,
		trustedPins: trustedPins,
	}
}

// reportCertIssue reports a TLS key mismatch.
func (r *tlsReporter) reportCertIssue(remoteURI, host, port string, connState tls.ConnectionState) {
	var certChain []string

	if len(connState.VerifiedChains) > 0 {
		certChain = marshalCert7468(connState.VerifiedChains[len(connState.VerifiedChains)-1])
	} else {
		certChain = marshalCert7468(connState.PeerCertificates)
	}

	report := newTLSReport(host, port, connState.ServerName, certChain, r.trustedPins, r.cfg.AppVersion)

	if !r.hasRecentlySentReport(report) {
		r.recordReport(report)
		go report.sendReport(r.cfg, remoteURI)
	}
}

// hasRecentlySentReport returns whether the report was already sent within the last 24 hours.
func (r *tlsReporter) hasRecentlySentReport(report tlsReport) bool {
	var validReports []sentReport

	for _, r := range r.sentReports {
		if time.Since(r.t) < 24*time.Hour {
			validReports = append(validReports, r)
		}
	}

	r.sentReports = validReports

	for _, r := range r.sentReports {
		if cmp.Equal(report, r.r) {
			return true
		}
	}

	return false
}

// recordReport records the given report and the current time so we can check whether we recently sent this report.
func (r *tlsReporter) recordReport(report tlsReport) {
	r.sentReports = append(r.sentReports, sentReport{r: report, t: time.Now()})
}

func marshalCert7468(certs []*x509.Certificate) (pemCerts []string) {
	var buffer bytes.Buffer
	for _, cert := range certs {
		if err := pem.Encode(&buffer, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		}); err != nil {
			logrus.WithField("pkg", "pmapi/tls-pinning").WithError(err).Error("Failed to encode TLS certificate")
		}
		pemCerts = append(pemCerts, buffer.String())
		buffer.Reset()
	}

	return pemCerts
}
