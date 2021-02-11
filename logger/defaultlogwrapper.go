// +build !nestedformatter

package logger

import "github.com/sirupsen/logrus"

func (l *StandardLogger) setFormatter() {
	l.Logger.Formatter = &logrus.JSONFormatter{}
}
