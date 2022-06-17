package api

import (
	"context"
)

func (m *manager) SendSimpleMetric(ctx context.Context, category, action, label string) error {
	r := m.r(ctx).SetQueryParams(map[string]string{
		"Category": category,
		"Action":   action,
		"Label":    label,
	})
	if _, err := wrapNoConnection(r.Get("/metrics")); err != nil {
		return err
	}
	return nil
}
