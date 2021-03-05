package logger

type AuditLogger struct {
	sl        StandardLogger
	operation OperationValue
}

func newAuditLogger(l *StandardLogger, operation string) *AuditLogger {
	var standardLogger = &AuditLogger{sl: *l, operation: Operation(operation)}
	return standardLogger
}

func (l *AuditLogger) Audit(message string) {
	l.sl.AuditWithOperation(message, l.operation)
}

func (l *AuditLogger) Start(message string, vs ...AuditValue) {
	vs = append(vs, Start())
	l.AuditWithFields(message, vs...)
}

func (l *AuditLogger) End(message string, vs ...AuditValue) {
	vs = append(vs, End())
	l.AuditWithFields(message, vs...)
}

func (l *AuditLogger) AuditWithFields(message string, vs ...AuditValue) {
	l.sl.AuditWithOperationAndFields(message, l.operation, vs...)
}
