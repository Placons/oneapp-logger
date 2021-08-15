package logger

import "github.com/google/uuid"

type AuditLogger struct {
	sl        StandardLogger
	operation OperationValue
	uuid      AuditUUIDValue
}

func newAuditLogger(l *StandardLogger, operation string) *AuditLogger {
	u, err := uuid.NewUUID()
	if err == nil {
		var standardLogger = &AuditLogger{sl: *l, operation: Operation(operation), uuid: AuditUUID(u.String())}
		return standardLogger
	} else {
		var standardLogger = &AuditLogger{sl: *l, operation: Operation(operation), uuid: AuditUUID("error")}
		return standardLogger
	}
}

func (l *AuditLogger) Audit(message string) {
	l.sl.AuditWithOperationAndUUIDAndFields(message, l.operation, l.uuid)
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
	l.sl.AuditWithOperationAndUUIDAndFields(message, l.operation, l.uuid, vs...)
}
