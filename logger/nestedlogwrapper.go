// +build nestedformatter

package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
)

func (l *StandardLogger) setFormatter() {
	l.Logger.Formatter = &nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"appname", "package", "function", "line"},
	}
}
