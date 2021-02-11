package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAuditHook_Fire(t *testing.T) {
	e := logrus.Entry{
		Data: map[string]interface{}{
			"audit": true,
		},
		Level: logrus.ErrorLevel,
	}
	ch := AuditHook{}
	_ = ch.Fire(&e)

	if e.Level != logrus.InfoLevel {
		t.Errorf("Level %s is not correct.", e.Level)
	}
}
