# oneapp-logger

Logging library to standardize logging across go services

## Usage

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