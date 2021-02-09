package logger

import (
	"github.com/sirupsen/logrus"
)

type AppNameHook struct {
	appName string
}

func (AppNameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h AppNameHook) Fire(e *logrus.Entry) error {
	if e.Data == nil {
		e.Data = make(logrus.Fields)
	}
	e.Data["appname"] = h.appName
	return nil
}
