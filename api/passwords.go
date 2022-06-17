package pmapi

import (
	"encoding/base64"

	"github.com/ProtonMail/go-srp"
	"github.com/pkg/errors"
)

// HashMailboxPassword expectects 128bit long salt encoded by standard base64.
func HashMailboxPassword(password []byte, salt string) ([]byte, error) {
	if salt == "" {
		return password, nil
	}

	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode salt")
	}

	hash, err := srp.MailboxPassword(password, decodedSalt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
	}

	return hash[len(hash)-31:], nil
}
