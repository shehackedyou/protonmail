package api

import (
	"github.com/go-resty/resty"
	"github.com/sirupsen/logrus"
)

// restyLogger decreases debug level to trace level so resty logs
// are not logged as debug but trace instead. Resty logging is too
// verbose which we don't want to have in debug level.
type restyLogger struct {
	logrus *logrus.Entry
}

func (l *restyLogger) Errorf(format string, v ...interface{}) {
	l.logrus.Errorf(format, v...)
}

func (l *restyLogger) Warnf(format string, v ...interface{}) {
	l.logrus.Warnf(format, v...)
}

func (l *restyLogger) Debugf(format string, v ...interface{}) {
	l.logrus.Tracef(format, v...)
}

func (m *manager) SetLogging(logger *logrus.Entry, verbose bool) {
	if verbose {
		m.rc.SetLogger(&restyLogger{logrus: logger})
		m.rc.SetDebug(true)
		return
	}

	m.rc.OnBeforeRequest(func(_ *resty.Client, req *resty.Request) error {
		logger.Infof("Requesting %s %s", req.Method, req.URL)
		return nil
	})
	m.rc.OnAfterResponse(func(_ *resty.Client, res *resty.Response) error {
		log := logger.WithFields(logrus.Fields{
			"error":    res.Error(),
			"status":   res.StatusCode(),
			"duration": res.Time(),
		})
		if res.Request == nil {
			log.Warn("Requested unknown request")
			return nil
		}
		log.Debugf("Requested %s %s", res.Request.Method, res.Request.URL)
		return nil
	})
	m.rc.OnError(func(req *resty.Request, err error) {
		logger.WithError(err).Warnf("Failed request %s %s", req.Method, req.URL)
	})
}
