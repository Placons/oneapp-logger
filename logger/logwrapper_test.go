package logger

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestShouldSetLoggingLevel(t *testing.T) {
	standardLogger := NewStandardLogger("my-test-app")
	standardLogger.SetLoggingLevel("DEBUG")

	actual := standardLogger.GetLevel()

	if actual != logrus.DebugLevel {
		t.Errorf("Unexpected logging level found. Expected %s, got %s", logrus.DebugLevel.String(), actual.String())
	}
}

func TestShouldSetDefaultLoggingLevelToInfoWhenGivenUnknownLevel(t *testing.T) {
	standardLogger := NewStandardLogger("my-test-app")
	standardLogger.SetLoggingLevel("some-unknown-level")

	actual := standardLogger.GetLevel()

	if actual != logrus.InfoLevel {
		t.Errorf("Unexpected logging level found. Expected %s, got %s", logrus.InfoLevel.String(), actual.String())
	}
}

func TestLogErrorWithErr(t *testing.T) {
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.ErrorWithErr("Test error", errors.New("This is the error tag"))

	if !strings.Contains(buf.String(), "Test error") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestLogErrorWithErrAndFields(t *testing.T) {
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.ErrorWithErrAndFields("Test error", errors.New("This is the original error"), map[string]interface{}{
		"some-tag": "some-value"})

	if !strings.Contains(buf.String(), "some-value") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestLogInfoWithFields(t *testing.T) {
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.InfoWithFields("Test error", map[string]interface{}{
		"some-tag": "some-value"})

	if !strings.Contains(buf.String(), "some-value") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

// test if caller information (like "function") is logged on all levels if environment LOG_CALLER_ALL_LEVELS is present
func TestStandardLogger_DebugWithCallerEnv(t *testing.T) {
	_ = os.Setenv(LogCallerAllLevels, "true")
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
		_ = os.Unsetenv(LogCallerAllLevels)
	}()

	l.Debug("test")
	if !strings.Contains(buf.String(), "function") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

// test if caller information (like "function") is *not* logged on all levels if environment LOG_CALLER_ALL_LEVELS is not present
func TestStandardLogger_DebugWithoutCallerEnv(t *testing.T) {
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.Debug("test")
	if strings.Contains(buf.String(), "function") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestStandardLogger_Audit(t *testing.T) {
	_ = os.Unsetenv(LogCallerAllLevels)
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.AuditWithOperation("test", Operation("op"))
	// log message should not be of level error, but info
	if strings.Contains(buf.String(), "\"level\":\"error\"") {
		t.Errorf("Log message does contain error level: %s", buf.String())
	}
	if !strings.Contains(buf.String(), "\"level\":\"info\"") {
		t.Errorf("Log message does not contain info level: %s", buf.String())
	}
	if strings.Contains(buf.String(), "function") {
		t.Errorf("Log message contains caller information: %s", buf.String())
	}
}

func TestStandardLogger_AuditWithOperation(t *testing.T) {
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
	}()

	l.AuditWithOperation("user updated goal count", Operation("test"))
	// log message should contain goals
	if !strings.Contains(buf.String(), "operation\":\"test") {
		t.Errorf("Log message does goals field: %s", buf.String())
	}
}

func prepareLogger() (*StandardLogger, *bytes.Buffer) {
	l := NewStandardLogger("my-test-app")
	l.SetLevel(logrus.DebugLevel)
	var buf bytes.Buffer
	l.SetOutput(&buf)
	return l, &buf
}
