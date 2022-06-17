package api

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("protonmail", "api") //nolint[gochecknoglobals]
