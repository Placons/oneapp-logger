package logger

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type ErrorType string

const (
	DownstreamServiceError ErrorType = "DownstreamServiceError"
	BadRequestError        ErrorType = "BadRequestError"
	ResponseParsingError   ErrorType = "ResponseParsingError"
)

var LoggingLevels = logrus.AllLevels

type Event struct {
	errorType ErrorType
	message   string
}

type StandardLogger struct {
	*logrus.Logger
	appname string
}

// NewLogger initializes the standard logger
func NewStandardLogger(appname string) *StandardLogger {
	var baseLogger = logrus.New()
	baseLogger.SetOutput(os.Stdout)

	var standardLogger = &StandardLogger{baseLogger, appname}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// NewLogger initializes a logger with custom formatter and output
func NewLogger(appname string, output io.Writer, formatter logrus.Formatter) *StandardLogger {
	var baseLogger = logrus.New()
	baseLogger.SetOutput(output)

	var standardLogger = &StandardLogger{baseLogger, appname}

	standardLogger.Formatter = formatter

	return standardLogger
}

// Declare variables to store log messages as new Events. Only examples error, list to be updated
var (
	downstreamServiceError = Event{DownstreamServiceError, "Downstream service returned failed due to http status %d"}
	badRequestError        = Event{BadRequestError, "Invalid request"}
	responseParsingError   = Event{ResponseParsingError, "Received a non-parseable response"}
)

func (l *StandardLogger) LogDownstreamServiceError(statusCode int) {
	l.withStandardFields().Errorf(downstreamServiceError.message, statusCode)
}

func (l *StandardLogger) LogBadRequestError() {
	l.withStandardFields().Errorf(badRequestError.message)
}

func (l *StandardLogger) LogResponseParsingError(argumentName string) {
	l.withStandardFields().Errorf(responseParsingError.message)
}

func (l *StandardLogger) Trace(message string) {
	l.withStandardFields().Trace(message)
}

func (l *StandardLogger) Info(message string) {
	l.withStandardFields().Info(message)
}

func (l *StandardLogger) Debug(message string) {
	l.withStandardFields().Debug(message)
}

func (l *StandardLogger) Warn(message string) {
	l.withStandardFields().Warn(message)
}

func (l *StandardLogger) Error(message string) {
	l.withStandardFields().Error(message)
}

func (l *StandardLogger) ErrorWithErr(message string, err error) {
	l.withStandardFields().WithField("errorTrace", err).Error(message)
}

func (l *StandardLogger) ErrorWithErrAndFields(message string, err error, fields map[string]interface{}) {
	l.withFields(fields).WithField("errorTrace", err).Error(message)
}

func (l *StandardLogger) Tracef(message string, args ...interface{}) {
	l.withStandardFields().Tracef(message, args...)
}

func (l *StandardLogger) Infof(message string, args ...interface{}) {
	l.withStandardFields().Infof(message, args...)
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

func (l *StandardLogger) Debugf(message string, args ...interface{}) {
	l.withStandardFields().Debugf(message, args...)
}

func (l *StandardLogger) Warnf(message string, args ...interface{}) {
	l.withStandardFields().Warnf(message, args...)
}

func (l *StandardLogger) Errorf(message string, args ...interface{}) {
	l.withStandardFields().Errorf(message, args...)
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

func (l *StandardLogger) withStandardFields() *logrus.Entry {
	standardFields := logrus.Fields{
		"appname": l.appname,
	}

	return l.WithFields(standardFields)
}

func (l *StandardLogger) withFields(fields map[string]interface{}) *logrus.Entry {
	return l.withStandardFields().WithFields(fields)
}
