# go-logger

Logging library to standardize logging across go services. It provides a logwrapper for basic logging mechanism, supporting all different levels of logging (error, info, warn, debug etc).
Under the hood it is using [logrus](https://github.com/sirupsen/logrus)
This library provides wrapper for all logrus methods (Info, Infof, Error, Errorf etc) plus some additional ones

## Usage

### Import
Import go-logger
```
import (
	"github.com/nsterg/go-logger/logger"
)
```
or 
```
go get github.com/nsterg/go-logger
```

### Initialize
From your service initialize a standard logger passing the application's name. By default this is producing logs in json format
```
standardLogger := logger.NewStandardLogger("my-app")
```

### Use cases
Log an error adding an additional message
```
logger.ErrorWithErr("Downstream service failed due to http status 500", err)
```

This will produce this logging
```
{"appname":"my-app","level":"error","msg":"Downstream service failed due to http status 500","errorTrace":"The actual error trace goes here","time":"2020-11-27T08:59:12+02:00"}
```

Log error with additional message and custom fields
```
logger.ErrorWithErrAndFields("Error deleting item", err, map[string]interface{}{
			"itemId": 123456,
		})
```
This will produce this logging
```
{"appname":"my-app","level":"error","msg":"Error deleting item","itemId":123456,"errorTrace":"The actual error trace goes here","time":"2020-11-27T08:59:12+02:00"}
```

The use of additional fields can also apply to all other logging levels


### Logging Levels
You can set the logging level on a Logger, then it will only log entries with that severity or anything above it. The library is exposing ERROR, WARN, INFO, DEBUG, TRACE.
The default logging level is info and you may override this and set it for example to debug:
```
standardLogger.SetLoggingLevel("DEBUG")
```
