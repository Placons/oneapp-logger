package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type ErrorType string

const (
	DownstreamServiceError ErrorType = "DownstreamServiceError"
	BadRequestError        ErrorType = "BadRequestError"
	ResponseParsingError   ErrorType = "ResponseParsingError"
)

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

func (l *StandardLogger) Info(message string) {
	l.withStandardFields().Info(message)
}

func (l *StandardLogger) Debug(message string) {
	l.withStandardFields().Debug(message)
}

func (l *StandardLogger) Warn(message string) {
	l.withStandardFields().Warn(message)
}
func (l *StandardLogger) withStandardFields() *logrus.Entry {
	standardFields := logrus.Fields{
		"appname": l.appname,
	}

	return l.WithFields(standardFields)
}
