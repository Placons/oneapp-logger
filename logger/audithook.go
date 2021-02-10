package logger

import (
	"github.com/sirupsen/logrus"
)

type AuditHook struct {
	appName string
}

func (AuditHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (h AuditHook) Fire(e *logrus.Entry) error {
	if e.Data != nil && e.Data["audit"] == true {
		e.Level = logrus.InfoLevel
	}
	return nil
}
