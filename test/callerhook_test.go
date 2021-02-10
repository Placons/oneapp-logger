package loggertest

import (
	"bytes"
	"github.com/Placons/oneapp-logger/logger"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"testing"
)

// we need a new package for this test as the logger package is skipped for caller calculation
func TestStandardLogger_LogWithCaller(t *testing.T) {
	_ = os.Setenv(logger.LogCallerAllLevels, "true")
	l, buf := prepareLogger()
	defer func() {
		l.SetOutput(os.Stderr)
		_ = os.Unsetenv(logger.LogCallerAllLevels)
	}()

	l.Info("test")
	// log message should contain goals
	if !strings.Contains(buf.String(), "\"function\":\"TestStandardLogger_LogWithCaller\"") {
		t.Errorf("Log message does goals field: %s", buf.String())
	}

}

func prepareLogger() (*logger.StandardLogger, *bytes.Buffer) {
	l := logger.NewStandardLogger("my-test-app")
	l.SetLevel(logrus.DebugLevel)
	var buf bytes.Buffer
	l.SetOutput(&buf)
	return l, &buf
}
