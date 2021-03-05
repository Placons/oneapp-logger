# oneapp-logger

Logging library to standardize logging across go services. It provides a logwrapper for basic logging mechanism, supporting all deifferent
levels of logging (error, info, warn, debug etc). Under the hood it is using [logrus](https://github.com/sirupsen/logrus)

## Usage

### Import oneapp-logger

```go
import (
	"github.com/Placons/oneapp-logger/logger"
)
```

or

```console
go get github.com/Placons/oneapp-logger
```

In case you are stuck with a specific verison and can't get the latest one:

```console
export GOSUMDB=off
```

From your service initialize a standard logger passing the application's name. By default this is producing logs in json format

```go
standardLogger := logger.NewStandardLogger("oneapp-logs-playground")
```

### Log info:

```go
standardLogger.Info("Some custom info message")
```

Log Warning, debug levels as above

#### Example logging

```json
{
  "appname": "oneapp-logs-playground",
  "level": "error",
  "msg": "Downstream service failed due to http status 500",
  "time": "2020-11-27T08:59:12+02:00"
}
```

# Logging Levels

You can set the logging level on a Logger, then it will only log entries with that severity or anything above it. The library is exposing
ERROR, WARN, INFO, DEBUG, TRACE. The default logging level is info and you may override this and set it for example to debug:

```go
standardLogger.SetLoggingLevel("DEBUG")
```

# Logging caller information

By default, logs with information about the caller are printed only on Panic, Fatal and Error level. To enable this information on all log
levels add LOG_CALLER_ALL_LEVELS as environment variable with any value.

# Audit logging

Audit logs will ALWAYS be printed and always as Info level regardless of the currently selected log level.

```go
standardLogger := logger.NewStandardLogger("my-test-app")
standardLogger.SetLoggingLevel("DEBUG")
standardLogger.AuditWithOperation("TestMsg", Operation("test-operation"), Start())
```

This will result in (see LogLevel)

```json
{
  "appname": "my-test-app",
  "level": "info",
  "msg": "TestMsg",
  "time": "2021-02-09T20:30:03+01:00",
  "audit": "true",
  "operation": "Test",
  "start": "2021-02-09T20:30:03+01:00"
}
```
For useful audit logs you should always provide some values. The easiest way to use the audit logging is
to get a special AuditLogger instance that has the operation preset and contains some
additional utility functions for printing start and end.

This will produce the same output as above.
```go
standardLogger := logger.NewStandardLogger("my-test-app")
auditLogger := standardLogger.Audit("test-operation")
auditLogger.Start("TestMsg", Operation("Test"))
```
