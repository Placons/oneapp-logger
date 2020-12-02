package logger

import (
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
