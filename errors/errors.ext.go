// This file Written Manually.

package errors

import (
	"fmt"
)

/*
 * Define BaseResponse
 */

var _ error = &Error{}

func NewErrorWithCode(code ErrorCode) *Error {
	return &Error{
		ErrorCode: code,
		Reason:    code.String(),
	}
}

func ErrorOK() *Error {
	return NewErrorWithCode(ErrorCode_OK)
}

func NewErrorWithError(err error) *Error {
	// check err is BaseResponse
	if br, ok := err.(*Error); ok {
		return br
	}

	return &Error{
		ErrorCode: ErrorCode_InternalError,
		Reason:    ErrorCode_InternalError.String(),
		Message:   err.Error(),
	}
}

// Error is implement of error interface.
func (x *Error) Error() string {
	return fmt.Sprintf("ErrorCode: %d, Reason: %s, Message: %s", x.ErrorCode, x.Reason, x.Message)
}

// Err returns as error.
func (x *Error) Err() error {
	if x.ErrorCode == ErrorCode_OK {
		return nil
	}

	return x
}

func (x *Error) Success() bool {
	return x.ErrorCode == ErrorCode_OK
}

func (x *Error) WithMessage(msg string) *Error {
	x.Message = msg
	return x
}

func (x *Error) WithError(err error) *Error {
	x.Message = err.Error()
	return x
}

/*
 * ErrorCode As Error.
 */

func (x ErrorCode) Error() string {
	return NewErrorWithCode(x).Error()
}

func (x ErrorCode) Err() error {
	if x == ErrorCode_OK {
		return nil
	}

	return NewErrorWithCode(x)
}

// Err2 returns as Error we defined.
func (x ErrorCode) Err2() *Error {
	return NewErrorWithCode(x)
}

func (x ErrorCode) WithMessage(msg string) *Error {
	return NewErrorWithCode(x).WithMessage(msg)
}

func (x ErrorCode) WithError(err error) *Error {
	return NewErrorWithCode(x).WithError(err)
}
