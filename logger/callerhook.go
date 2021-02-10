package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

const LogCallerAllLevels = "LOG_CALLER_ALL_LEVELS"

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
	if e.Data == nil {
		e.Data = make(logrus.Fields)
	}
	if e.Data["audit"] == true {
		return nil // do not log caller information on audit logs
	}
	targetFrameIndex := 0 + 2
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)
	frame := runtime.Frame{Function: "unknown"}
	frames := runtime.CallersFrames(programCounters[:n])
	for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
		var frameCandidate runtime.Frame
		frameCandidate, more = frames.Next()
		if frameIndex == targetFrameIndex {
			frame = frameCandidate
		}
	}

	if e.Data == nil {
		e.Data = make(logrus.Fields)
	}
	e.Data["file"] = frame.File
	e.Data["function"] = frame.Function
	e.Data["line"] = frame.Line
	return nil
}
