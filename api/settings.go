package pmapi

import (
	"context"

	"github.com/go-resty/resty"
)

type MailSettings struct {
	DisplayName        string
	Signature          string `json:",omitempty"`
	Theme              string `json:",omitempty"`
	AutoSaveContacts   int
	AutoWildcardSearch int
	ComposerMode       int
	MessageButtons     int
	ShowImages         int
	ShowMoved          int
	ViewMode           int
	ViewLayout         int
	SwipeLeft          int
	SwipeRight         int
	AlsoArchive        int
	Hotkeys            int
	PMSignature        int
	ImageProxy         int
	TLS                int
	RightToLeft        int
	AttachPublicKey    int
	Sign               int
	PGPScheme          PackageFlag
	PromptPin          int
	Autocrypt          int
	NumMessagePerPage  int
	DraftMIMEType      string
	ReceiveMIMEType    string
	ShowMIMEType       string

	// Undocumented -- there's only `null` in example:
	// AutoResponder string
}

// GetMailSettings gets contact details specified by contact ID.
func (c *client) GetMailSettings(ctx context.Context) (settings MailSettings, err error) {
	var res struct {
		MailSettings MailSettings
	}

	if _, err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get("/mail/v4/settings")
	}); err != nil {
		return MailSettings{}, err
	}

	return res.MailSettings, nil
}
