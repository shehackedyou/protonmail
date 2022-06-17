package pmapi

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
)

// ErrTLSMismatch indicates that no TLS fingerprint match could be found.
var ErrTLSMismatch = errors.New("no TLS fingerprint match found")

type pinChecker struct {
	trustedPins []string
}

func newPinChecker(trustedPins []string) *pinChecker {
	return &pinChecker{
		trustedPins: trustedPins,
	}
}

// checkCertificate returns whether the connection presents a known TLS certificate.
func (p *pinChecker) checkCertificate(conn net.Conn) error {
	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		return errors.New("connection is not a TLS connection")
	}

	connState := tlsConn.ConnectionState()

	for _, peerCert := range connState.PeerCertificates {
		fingerprint := certFingerprint(peerCert)

		for _, pin := range p.trustedPins {
			if pin == fingerprint {
				return nil
			}
		}
	}

	return ErrTLSMismatch
}

func certFingerprint(cert *x509.Certificate) string {
	hash := sha256.Sum256(cert.RawSubjectPublicKeyInfo)
	return fmt.Sprintf(`pin-sha256=%q`, base64.StdEncoding.EncodeToString(hash[:]))
}
