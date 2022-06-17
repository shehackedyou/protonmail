package api

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/ProtonMail/gopenpgp/crypto"
	"github.com/go-resty/resty"
)

// client is a client of the protonmail API. It implements the Client interface.
type client struct {
	manager clientManager

	uid, acc, ref string
	authHandlers  []AuthRefreshHandler
	authLocker    sync.RWMutex

	user        *User
	addresses   AddressList
	userKeyRing *crypto.KeyRing
	addrKeyRing map[string]*crypto.KeyRing
	keyRingLock sync.Locker

	exp time.Time
}

func newClient(manager clientManager, uid string) *client {
	return &client{
		manager:     manager,
		uid:         uid,
		addrKeyRing: make(map[string]*crypto.KeyRing),
		keyRingLock: &sync.RWMutex{},
	}
}

func (c *client) withAuth(acc, ref string, exp time.Time) *client {
	c.acc = acc
	c.ref = ref
	c.exp = exp

	return c
}

func (c *client) r(ctx context.Context) (*resty.Request, error) {
	r := c.manager.r(ctx)

	if c.uid != "" {
		r.SetHeader("x-pm-uid", c.uid)
	}

	if time.Now().After(c.exp) {
		if err := c.authRefresh(ctx); err != nil {
			return nil, err
		}
	}

	c.authLocker.RLock()
	defer c.authLocker.RUnlock()

	if c.acc != "" {
		r.SetAuthToken(c.acc)
	}

	return r, nil
}

// do executes fn and may repeat execution in case of retry after "401 Unauthorized" error.
// Note: fn may be called more than once.
func (c *client) do(ctx context.Context, fn func(*resty.Request) (*resty.Response, error)) (*resty.Response, error) {
	r, err := c.r(ctx)
	if err != nil {
		return nil, err
	}

	res, err := wrapNoConnection(fn(r))
	if err != nil {
		if res.StatusCode() != http.StatusUnauthorized {
			// Return also response so caller has more options to decide what to do.
			return res, err
		}

		if !isAuthRefreshDisabled(ctx) {
			if err := c.authRefresh(ctx); err != nil {
				return nil, err
			}

			// We need to reconstruct request since access token is changed with authRefresh.
			r, err := c.r(ctx)
			if err != nil {
				return nil, err
			}

			return wrapNoConnection(fn(r))
		}

		return res, err
	}

	return res, nil
}
