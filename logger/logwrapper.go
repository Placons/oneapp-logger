package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var LoggingLevels = logrus.AllLevels

type StandardLogger struct {
	*logrus.Logger
}

func NewStandardLogger(appName string) *StandardLogger {
	var baseLogger = logrus.New()
	baseLogger.AddHook(CallerHook{})
	baseLogger.AddHook(AppNameHook{
		appName: appName,
	})
	baseLogger.AddHook(AuditHook{
		appName: appName,
	})
	baseLogger.SetOutput(os.Stdout)

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

func (l *StandardLogger) ErrorWithErr(message string, err error) {
	l.WithField("errorTrace", err).Error(message)
}

func (l *StandardLogger) ErrorWithErrAndFields(message string, err error, fields map[string]interface{}) {
	l.withFields(fields).WithField("errorTrace", err).Error(message)
}

func (l *StandardLogger) DebugWithFields(message string, fields map[string]interface{}) {
	l.withFields(fields).Debug(message)
}

func (l *StandardLogger) InfoWithFields(message string, fields map[string]interface{}) {
	l.withFields(fields).Info(message)
}

func (l *StandardLogger) WarnWithFields(message string, fields map[string]interface{}) {
	l.withFields(fields).Warn(message)
}

func (l *StandardLogger) Audit(message string) {
	// log level will actually be exchanged in the audit hook
	l.AuditWithFields(message, nil)
}

func (l *StandardLogger) AuditWithFields(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = map[string]interface{}{}
	}
	fields["type"] = "audit"
	// log level will actually be exchanged in the audit hook
	l.withFields(fields).Error(message)
}

func (l *StandardLogger) SetLoggingLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		log.Printf("Unknown logging level %s found. Falling back to INFO", lvl)
		l.Level = logrus.InfoLevel
		return
	}
	l.Level = level
}

func (l *StandardLogger) withFields(fields map[string]interface{}) *logrus.Entry {
	return l.WithFields(fields)
}
