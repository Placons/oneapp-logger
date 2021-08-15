package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var LoggingLevels = logrus.AllLevels

type StandardLogger struct {
	*logrus.Logger
	appName string
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

	var standardLogger = &StandardLogger{baseLogger, appName}

	standardLogger.setFormatter()

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

func (l *StandardLogger) Audit(op string) *AuditLogger {
	return newAuditLogger(l, op)
}

func (l *StandardLogger) AuditWithOperation(message string, operation OperationValue) {
	// log level will actually be exchanged in the audit hook
	l.AuditWithOperationAndFields(message, operation, nil)
}

func (l *StandardLogger) AuditWithOperationAndFields(message string, operation OperationValue, vs ...AuditValue) {
	l.AuditWithOperationAndUUIDAndFields(message, operation, nil, vs...)
}

func (l *StandardLogger) AuditWithOperationAndUUIDAndFields(message string, operation OperationValue, u AuditUUIDValue, vs ...AuditValue) {
	fields := map[string]interface{}{}
	fields["audit"] = true
	opKey, opValue := operation.Get()
	fields[opKey] = opValue
	if u != nil {
		uKey, uValue := u.Get()
		fields[uKey] = uValue
	}

	// log level will actually be exchanged in the audit hook
	for _, a := range vs {
		if a != nil {
			k, v := a.Get()
			fields[k] = v
		}
	}

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
