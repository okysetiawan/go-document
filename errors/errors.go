package errors

import "fmt"

type Code int

const (
	FailedWriteHeader Code = iota + 1000
	FailedWriteRows
	FailedWriteToDevice
)

type ErrorWrite struct {
	cause error
	code  Code
}

func (e *ErrorWrite) Error() string {
	return fmt.Sprintf("failed with status %d, cause: %v", e.code, e.cause)
}

func (e *ErrorWrite) Cause() error { return e.cause }

func WithCode(cause error, code Code) *ErrorWrite {
	switch cause.(type) {
	case *ErrorWrite:
		cause = cause.(*ErrorWrite).cause
	}
	return &ErrorWrite{cause: cause, code: code}
}
