package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestAuditLogger_Audit(t *testing.T) {
	l, buf := prepareAuditLogger()
	l.AuditWithFields("Test", Generic("abc", 1))

	if !strings.Contains(buf.String(), "abc\":1") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestAuditLogger_Start(t *testing.T) {
	l, buf := prepareAuditLogger()
	l.Start("Test", UserId("42"))

	if !strings.Contains(buf.String(), "userId\":\"42") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
	if !strings.Contains(buf.String(), "start\":") {
		t.Errorf("Log message is missing start time: %s", buf.String())
	}
}

func TestAuditLogger_TinkUser(t *testing.T) {
	l, buf := prepareAuditLogger()
	l.Start("Test", TinkUserId("42"))

	if !strings.Contains(buf.String(), "tinkUserId\":\"42") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestAuditLogger_YoliUser(t *testing.T) {
	l, buf := prepareAuditLogger()
	l.Start("Test", YoliUserId("42"))

	if !strings.Contains(buf.String(), "yoliUserId\":\"42") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func TestAuditLogger_Generic(t *testing.T) {
	l, buf := prepareAuditLogger()
	l.Start("Test", Generic("testKey", "testValue"))

	if !strings.Contains(buf.String(), "testKey\":\"testValue") {
		t.Errorf("Log message not correct: %s", buf.String())
	}
}

func prepareAuditLogger() (*AuditLogger, *bytes.Buffer) {
	sl := NewStandardLogger("my-test-app")
	al := sl.Audit("operation")
	var buf bytes.Buffer
	sl.SetOutput(&buf)
	return al, &buf
}
