package protonmail

import (
	"context"
)

// Unlock unlocks all the user and address keys using the given passphrase, creating user and address keyrings.
// If the keyrings are already present, they are not recreated.
func (c *client) Unlock(ctx context.Context, passphrase []byte) (err error) {
	c.keyRingLock.Lock()
	defer c.keyRingLock.Unlock()

	return c.unlock(ctx, passphrase)
}

// unlock unlocks the user's keys but without locking the keyring lock first.
// Should only be used internally by methods that first lock the lock.
func (c *client) unlock(ctx context.Context, passphrase []byte) error {
	if _, err := c.CurrentUser(ctx); err != nil {
		return err
	}

	if c.userKeyRing == nil {
		if err := c.unlockUser(passphrase); err != nil {
			return ErrUnlockFailed{err}
		}
	}

	for _, address := range c.addresses {
		if c.addrKeyRing[address.ID] == nil {
			if err := c.unlockAddress(passphrase, address); err != nil {
				return ErrUnlockFailed{err}
			}
		}
	}

	return nil
}

func (c *client) ReloadKeys(ctx context.Context, passphrase []byte) (err error) {
	c.keyRingLock.Lock()
	defer c.keyRingLock.Unlock()

	c.clearKeys()

	return c.unlock(ctx, passphrase)
}

func (c *client) clearKeys() {
	if c.userKeyRing != nil {
		c.userKeyRing.ClearPrivateParams()
		c.userKeyRing = nil
	}

	for id, kr := range c.addrKeyRing {
		if kr != nil {
			kr.ClearPrivateParams()
		}
		delete(c.addrKeyRing, id)
	}
}

func (c *client) IsUnlocked() bool {
	if c.userKeyRing == nil {
		return false
	}

	for _, address := range c.addresses {
		if address.HasKeys != MissingKeys && c.addrKeyRing[address.ID] == nil {
			return false
		}
	}

	return true
}
