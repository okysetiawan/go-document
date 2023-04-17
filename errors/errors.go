package errors

import (
	"errors"
	"fmt"
)

type Code int

const (
	CodeWriteHeader Code = iota + 1000
	CodeWriteRows
	CodeWriteToDevice
	CodeWriteTypesNil
	CodeWriteTypesInvalid
)

const (
	CodeReadReader Code = iota + 2000
	CodeReadEmpty
	CodeReadFileFailed
	CodeReadFileNotExist
	CodeReadUnmarshal
)

type ErrorWrite struct {
	cause error
	code  Code
}

func (e *ErrorWrite) Error() string {
	return fmt.Sprintf("failed with status %d, cause: %v", e.code, e.cause)
}

func (e *ErrorWrite) Cause() error { return e.cause }

func (e *ErrorWrite) WithCode(code Code) error {
	return WithCode(e.cause, code)
}

func WithCode(cause error, code Code) *ErrorWrite {
	switch cause.(type) {
	case *ErrorWrite:
		cause = cause.(*ErrorWrite).cause
	}
	return &ErrorWrite{cause: cause, code: code}
}

func New(msg string, args ...interface{}) *ErrorWrite { return &ErrorWrite{cause: errors.New(msg)} }

func Is(err, target error) bool { return errors.Is(err, target) }

func IsCode(err error, code Code) bool {
	if errw, ok := err.(*ErrorWrite); ok {
		return errw.code == code
	}

	return false
}
