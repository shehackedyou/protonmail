package protonmail

import (
	"io/ioutil"

	"golang.org/x/net/context"

	"github.com/ProtonMail/gopenpgp/crypto"
)

// DownloadAndVerify downloads a file and its signature from the given locations `file` and `sig`.
// The file and its signature are verified using the given keyring `kr`.
// If the file is verified successfully, it can be read from the returned reader.
// TLS fingerprinting is used to verify that connections are only made to known servers.
func (m *manager) DownloadAndVerify(kr *crypto.KeyRing, url, sig string) ([]byte, error) {
	fb, err := m.fetchFile(url)
	if err != nil {
		return nil, err
	}

	sb, err := m.fetchFile(sig)
	if err != nil {
		return nil, err
	}

	if err := kr.VerifyDetached(
		crypto.NewPlainMessage(fb),
		crypto.NewPGPSignature(sb),
		crypto.GetUnixTime(),
	); err != nil {
		return nil, err
	}

	return fb, nil
}

func (m *manager) fetchFile(url string) ([]byte, error) {
	res, err := m.r(ContextWithoutRetry(context.Background())).SetDoNotParseResponse(true).Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.RawBody())
	if err != nil {
		return nil, err
	}

	if err := res.RawBody().Close(); err != nil {
		return nil, err
	}

	return b, nil
}
