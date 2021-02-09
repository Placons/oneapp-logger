package logger

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCallerHook_Fire(t *testing.T) {
	e := logrus.Entry{}
	ch := CallerHook{}
	_ = ch.Fire(&e)

	if e.Data == nil {
		t.Errorf("Data not present.")
	}

	f := fmt.Sprintf("%s", e.Data["function"].(string))
	if e.Data["function"] == nil || !strings.HasSuffix(f, "TestCallerHook_Fire") {
		t.Errorf("Function %s is not correct.", f)
	}
	l := e.Data["line"].(int)
	if e.Data["line"] == nil || l != 14 {
		t.Errorf("Line %d is not correct.", l)
	}
}
