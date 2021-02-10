package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAppNameHook_Fire(t *testing.T) {
	e := logrus.Entry{}
	ch := AppNameHook{
		appName: "dummy",
	}
	_ = ch.Fire(&e)

	if e.Data == nil {
		t.Errorf("Data not present.")
	}

	n := e.Data["appname"].(string)
	if e.Data["appname"] == nil || n != "dummy" {
		t.Errorf("Appname %s is not correct.", n)
	}
}
