package logger

import "time"

type AuditElement struct {
	key   string
	value interface{}
}

func (a AuditElement) Get() (string, interface{}) {
	return a.key, a.value
}

type AuditValue interface {
	Get() (string, interface{})
}

type OperationValue AuditValue

func Operation(value interface{}) OperationValue {
	return AuditElement{
		key:   "operation",
		value: value,
	}
}

func Generic(key string, value interface{}) AuditValue {
	return AuditElement{
		key:   key,
		value: value,
	}
}

type UserValue AuditValue

func UserId(value string) UserValue {
	return AuditElement{
		key:   "userId",
		value: value,
	}
}
func TinkUserId(value string) UserValue {
	return AuditElement{
		key:   "tinkUserId",
		value: value,
	}
}
func YoliUserId(value string) UserValue {
	return AuditElement{
		key:   "yoliUserId",
		value: value,
	}
}

type TimeValue AuditValue

func timeNow(name string) TimeValue {
	return AuditElement{
		key:   name,
		value: time.Now(),
	}
}

func Start() TimeValue {
	return timeNow("start")
}

func End() TimeValue {
	return timeNow("end")
}
