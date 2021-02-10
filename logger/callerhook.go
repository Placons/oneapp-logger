package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

const LogCallerAllLevels = "LOG_CALLER_ALL_LEVELS"

var skipPackages = [...]string{"github.com/sirupsen/logrus", "github.com/Placons/oneapp-logger/logger"}

const maximumCallerDepth = 25
const minimumCallerDepth = 2

type CallerHook struct {
	levels []logrus.Level
}

func (h CallerHook) Levels() []logrus.Level {
	if os.Getenv(LogCallerAllLevels) != "" {
		return LoggingLevels
	}
	if h.levels == nil {
		return []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
	}
	return h.levels
}

func (CallerHook) Fire(e *logrus.Entry) error {
	defer func() {
		if r := recover(); r != nil {
			// Recovering without call to log to prevent infinite loop
			// we ignore any errors that might have occurred here!
			fmt.Println("PANIC in caller hook ", r)
		}
	}()
	frame := getCaller()
	if e.Data == nil {
		e.Data = make(logrus.Fields)
	}
	e.Data["package"] = getPackageName(frame.Func.Name())
	e.Data["function"] = getFunctionName(frame.Func.Name())
	e.Data["line"] = frame.Line
	return nil
}

// getCaller retrieves the name of the first non-logrus calling function
func getCaller() *runtime.Frame {
	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if !isSkipPackage(pkg) {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// getPackageName reduces a fully qualified function name to the package name
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func getFunctionName(f string) string {
	lastPeriod := strings.LastIndex(f, ".")
	f = f[lastPeriod+1:]
	return f
}

func isSkipPackage(p string) bool {
	for _, s := range skipPackages {
		if s == p {
			return true
		}
	}
	return false
}
