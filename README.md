# oneapp-logger

Logging library to standardize logging across go services. It provides a logwrapper for basic logging mechanism, supporting all deifferent levels of logging (error, info, warn, debug etc).
Under the hood it is using [logrus](https://github.com/sirupsen/logrus)

## Usage

Import oneapp-logger
```
import (
	"github.com/Placons/oneapp-logger/logger"
)
```
or 
```
go get github.com/Placons/oneapp-logger
```
In case you are stuck with a specific verison and can't get the latest one:
```
export GOSUMDB=off
```

From your service initialize a standard logger passing the application's name. By default this is producing logs in json format
```
standardLogger := logger.NewStandardLogger("oneapp-logs-playground")
```

Log standard errors:
```
standardLogger.LogDownstreamServiceError(500)
```

Log info:
```
standardLogger.Info("Some custom info message")
```
Log Warning, debug levels as above

Example logging
```
{"appname":"oneapp-logs-playground","level":"error","msg":"Downstream service failed due to http status 500","time":"2020-11-27T08:59:12+02:00"}
```

Logging Levels
You can set the logging level on a Logger, then it will only log entries with that severity or anything above it. The library is exposing ERROR, WARN, INFO, DEBUG, TRACE.
The default logging level is info and you may override this and set it for example to debug:
```
standardLogger.SetLoggingLevel("DEBUG")
```
