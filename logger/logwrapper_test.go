package logger

import (
	"errors"
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
	standardLogger := NewStandardLogger("my-test-app")

	standardLogger.ErrorWithErr("Test error", errors.New("This is the error tag"))
}

func TestLogErrorWithErrAndFields(t *testing.T) {
	standardLogger := NewStandardLogger("my-test-app")

	standardLogger.ErrorWithErrAndFields("Test error", errors.New("This is the original error"), map[string]interface{}{
		"some-tag": "some-value"})
}

func TestLogInfoWithFields(t *testing.T) {
	standardLogger := NewStandardLogger("my-test-app")

	standardLogger.InfoWithFields("Test error", map[string]interface{}{
		"some-tag": "some-value"})
}
