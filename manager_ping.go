package protonmail

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	// retryConnectionSleeps defines a smooth cool down in seconds.
	retryConnectionSleeps = []int{2, 5, 10, 30, 60} // nolint[gochecknoglobals]
)

func (m *manager) pingUntilSuccess() {
	if m.isPingOngoing() {
		logrus.Debug("Ping already ongoing")
		return
	}
	m.pingingStarted()
	defer m.pingingStopped()

	attempt := 0
	for {
		ctx := ContextWithoutRetry(context.Background())
		err := m.testPing(ctx)
		if err == nil {
			return
		}

		waitTime := getRetryConnectionSleep(attempt)
		attempt++
		logrus.WithError(err).WithField("attempt", attempt).WithField("wait", waitTime).Debug("Connection (still) not available")
		time.Sleep(waitTime)
	}
}

func (m *manager) isPingOngoing() bool {
	m.pingMutex.RLock()
	defer m.pingMutex.RUnlock()

	return m.isPinging
}

func (m *manager) pingingStarted() {
	m.pingMutex.Lock()
	defer m.pingMutex.Unlock()
	m.isPinging = true
}

func (m *manager) pingingStopped() {
	m.pingMutex.Lock()
	defer m.pingMutex.Unlock()
	m.isPinging = false
}

func getRetryConnectionSleep(idx int) time.Duration {
	if idx >= len(retryConnectionSleeps) {
		idx = len(retryConnectionSleeps) - 1
	}
	sec := retryConnectionSleeps[idx]
	return time.Duration(sec) * time.Second
}

func (m *manager) testPing(ctx context.Context) error {
	if _, err := m.r(ctx).Get("/tests/ping"); err != nil {
		return err
	}
	return nil
}
