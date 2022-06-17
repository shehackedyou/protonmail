package api

import (
	"context"
)

// Report sends request as json or multipart (if has attachment).
func (m *manager) ReportBug(ctx context.Context, rep ReportBugReq) error {
	if rep.ClientType == 0 {
		rep.ClientType = EmailClientType
	}

	if rep.Client == "" {
		rep.Client = m.cfg.GetUserAgent()
	}

	if rep.ClientVersion == "" {
		rep.ClientVersion = m.cfg.AppVersion
	}

	r := m.r(ctx).SetMultipartFormData(rep.GetMultipartFormData())

	for _, att := range rep.Attachments {
		r = r.SetMultipartField(att.name, att.name, att.mime, att.body)
	}

	if _, err := wrapNoConnection(r.Post("/reports/bug")); err != nil {
		return err
	}

	return nil
}
