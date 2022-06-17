package protonmail

import (
	"context"
	"io"

	"github.com/ProtonMail/gopenpgp/crypto"
	"github.com/go-resty/resty"
)

// Client defines the interface of a PMAPI client.
type Client interface {
	Auth2FA(context.Context, string) error
	AuthSalt(ctx context.Context) (string, error)
	AuthDelete(context.Context) error
	AddAuthRefreshHandler(AuthRefreshHandler)

	GetUser(ctx context.Context) (*User, error)
	CurrentUser(ctx context.Context) (*User, error)
	UpdateUser(ctx context.Context) (*User, error)
	Unlock(ctx context.Context, passphrase []byte) (err error)
	ReloadKeys(ctx context.Context, passphrase []byte) (err error)
	IsUnlocked() bool

	Addresses() AddressList
	GetAddresses(context.Context) (addresses AddressList, err error)
	ReorderAddresses(ctx context.Context, addressIDs []string) error

	GetEvent(ctx context.Context, eventID string) (*Event, error)

	SendMessage(context.Context, string, *SendMessageReq) (sent, parent *Message, err error)
	CreateDraft(ctx context.Context, m *Message, parent string, action int) (created *Message, err error)
	Import(context.Context, ImportMsgReqs) ([]*ImportMsgRes, error)

	CountMessages(ctx context.Context, addressID string) ([]*MessagesCount, error)
	ListMessages(ctx context.Context, filter *MessagesFilter) ([]*Message, int, error)
	GetMessage(ctx context.Context, apiID string) (*Message, error)
	DeleteMessages(ctx context.Context, apiIDs []string) error
	LabelMessages(ctx context.Context, apiIDs []string, labelID string) error
	UnlabelMessages(ctx context.Context, apiIDs []string, labelID string) error
	MarkMessagesRead(ctx context.Context, apiIDs []string) error
	MarkMessagesUnread(ctx context.Context, apiIDs []string) error

	ListLabels(ctx context.Context) ([]*Label, error)
	CreateLabel(ctx context.Context, label *Label) (*Label, error)
	UpdateLabel(ctx context.Context, label *Label) (*Label, error)
	DeleteLabel(ctx context.Context, labelID string) error
	EmptyFolder(ctx context.Context, labelID string, addressID string) error

	ListFilters(ctx context.Context) ([]Filter, error)
	CreateFilter(ctx context.Context, filter Filter) (Filter, error)
	UpdateFilter(ctx context.Context, filter Filter) (Filter, error)
	DeleteFilter(ctx context.Context, filterID string) error

	GetMailSettings(ctx context.Context) (MailSettings, error)
	GetContactEmailByEmail(context.Context, string, int, int) ([]ContactEmail, error)
	GetContactByID(context.Context, string) (Contact, error)
	DecryptAndVerifyCards([]Card) ([]Card, error)

	GetAttachment(ctx context.Context, id string) (att io.ReadCloser, err error)
	CreateAttachment(ctx context.Context, att *Attachment, r io.Reader, sig io.Reader) (created *Attachment, err error)

	GetUserKeyRing() (*crypto.KeyRing, error)
	KeyRingForAddressID(string) (kr *crypto.KeyRing, err error)
	GetPublicKeysForEmail(context.Context, string) ([]PublicKey, bool, error)
}

type AuthRefreshHandler func(*AuthRefresh)

type clientManager interface {
	r(context.Context) *resty.Request
	authRefresh(context.Context, string, string) (*AuthRefresh, error)
	setSentryUserID(userID string)
}
