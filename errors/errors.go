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
	code  int
}

func (e *ErrorWrite) Error() string {
	return fmt.Sprintf("failed with status %d, cause: %v", e.code, e.cause)
}

func (e *ErrorWrite) Cause() error { return e.cause }

func WithCode(cause error, code Code) *ErrorWrite {
	return &ErrorWrite{}
}
