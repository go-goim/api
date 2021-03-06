// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: message/v1/message.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SendMessageReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendMessageReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendMessageReqMultiError,
// or nil if none found.
func (m *SendMessageReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetFrom()) < 20 {
		err := SendMessageReqValidationError{
			field:  "From",
			reason: "value length must be at least 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTo()) < 20 {
		err := SendMessageReqValidationError{
			field:  "To",
			reason: "value length must be at least 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SessionType

	if _, ok := MessageContentType_name[int32(m.GetContentType())]; !ok {
		err := SendMessageReqValidationError{
			field:  "ContentType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetContent()); l < 1 || l > 4096 {
		err := SendMessageReqValidationError{
			field:  "Content",
			reason: "value length must be between 1 and 4096 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.SessionId != nil {
		// no validation rules for SessionId
	}

	if len(errors) > 0 {
		return SendMessageReqMultiError(errors)
	}

	return nil
}

// SendMessageReqMultiError is an error wrapping multiple validation errors
// returned by SendMessageReq.ValidateAll() if the designated constraints
// aren't met.
type SendMessageReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageReqMultiError) AllErrors() []error { return m }

// SendMessageReqValidationError is the validation error returned by
// SendMessageReq.Validate if the designated constraints aren't met.
type SendMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageReqValidationError) ErrorName() string { return "SendMessageReqValidationError" }

// Error satisfies the builtin error interface
func (e SendMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageReqValidationError{}

// Validate checks the field values on SendMessageResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SendMessageResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageRespMultiError, or nil if none found.
func (m *SendMessageResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionId

	// no validation rules for MsgId

	if len(errors) > 0 {
		return SendMessageRespMultiError(errors)
	}

	return nil
}

// SendMessageRespMultiError is an error wrapping multiple validation errors
// returned by SendMessageResp.ValidateAll() if the designated constraints
// aren't met.
type SendMessageRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageRespMultiError) AllErrors() []error { return m }

// SendMessageRespValidationError is the validation error returned by
// SendMessageResp.Validate if the designated constraints aren't met.
type SendMessageRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageRespValidationError) ErrorName() string { return "SendMessageRespValidationError" }

// Error satisfies the builtin error interface
func (e SendMessageRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageRespValidationError{}

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Message) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MessageMultiError, or nil if none found.
func (m *Message) ValidateAll() error {
	return m.validate(true)
}

func (m *Message) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MsgId

	// no validation rules for From

	// no validation rules for To

	// no validation rules for SessionType

	// no validation rules for SessionId

	// no validation rules for ContentType

	// no validation rules for Content

	// no validation rules for CreateTime

	if len(errors) > 0 {
		return MessageMultiError(errors)
	}

	return nil
}

// MessageMultiError is an error wrapping multiple validation errors returned
// by Message.ValidateAll() if the designated constraints aren't met.
type MessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageMultiError) AllErrors() []error { return m }

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on PushMessageReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushMessageReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushMessageReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushMessageReqMultiError,
// or nil if none found.
func (m *PushMessageReq) ValidateAll() error {
	return m.validate(true)
}

func (m *PushMessageReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PushMessageReqValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PushMessageReqValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PushMessageReqValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PushMessageReqMultiError(errors)
	}

	return nil
}

// PushMessageReqMultiError is an error wrapping multiple validation errors
// returned by PushMessageReq.ValidateAll() if the designated constraints
// aren't met.
type PushMessageReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushMessageReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushMessageReqMultiError) AllErrors() []error { return m }

// PushMessageReqValidationError is the validation error returned by
// PushMessageReq.Validate if the designated constraints aren't met.
type PushMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushMessageReqValidationError) ErrorName() string { return "PushMessageReqValidationError" }

// Error satisfies the builtin error interface
func (e PushMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushMessageReqValidationError{}

// Validate checks the field values on PushMessageResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PushMessageResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushMessageResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PushMessageRespMultiError, or nil if none found.
func (m *PushMessageResp) ValidateAll() error {
	return m.validate(true)
}

func (m *PushMessageResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PushMessageRespValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PushMessageRespValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PushMessageRespValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PushMessageRespMultiError(errors)
	}

	return nil
}

// PushMessageRespMultiError is an error wrapping multiple validation errors
// returned by PushMessageResp.ValidateAll() if the designated constraints
// aren't met.
type PushMessageRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushMessageRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushMessageRespMultiError) AllErrors() []error { return m }

// PushMessageRespValidationError is the validation error returned by
// PushMessageResp.Validate if the designated constraints aren't met.
type PushMessageRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushMessageRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushMessageRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushMessageRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushMessageRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushMessageRespValidationError) ErrorName() string { return "PushMessageRespValidationError" }

// Error satisfies the builtin error interface
func (e PushMessageRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushMessageResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushMessageRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushMessageRespValidationError{}

// Validate checks the field values on QueryOfflineMessageReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryOfflineMessageReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryOfflineMessageReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryOfflineMessageReqMultiError, or nil if none found.
func (m *QueryOfflineMessageReq) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryOfflineMessageReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUserId()) < 20 {
		err := QueryOfflineMessageReqValidationError{
			field:  "UserId",
			reason: "value length must be at least 20 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLastMsgId() <= 0 {
		err := QueryOfflineMessageReqValidationError{
			field:  "LastMsgId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for OnlyCount

	// no validation rules for Page

	if val := m.GetPageSize(); val < 1 || val > 100 {
		err := QueryOfflineMessageReqValidationError{
			field:  "PageSize",
			reason: "value must be inside range [1, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QueryOfflineMessageReqMultiError(errors)
	}

	return nil
}

// QueryOfflineMessageReqMultiError is an error wrapping multiple validation
// errors returned by QueryOfflineMessageReq.ValidateAll() if the designated
// constraints aren't met.
type QueryOfflineMessageReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryOfflineMessageReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryOfflineMessageReqMultiError) AllErrors() []error { return m }

// QueryOfflineMessageReqValidationError is the validation error returned by
// QueryOfflineMessageReq.Validate if the designated constraints aren't met.
type QueryOfflineMessageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryOfflineMessageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryOfflineMessageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryOfflineMessageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryOfflineMessageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryOfflineMessageReqValidationError) ErrorName() string {
	return "QueryOfflineMessageReqValidationError"
}

// Error satisfies the builtin error interface
func (e QueryOfflineMessageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryOfflineMessageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryOfflineMessageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryOfflineMessageReqValidationError{}

// Validate checks the field values on QueryOfflineMessageResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryOfflineMessageResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryOfflineMessageResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryOfflineMessageRespMultiError, or nil if none found.
func (m *QueryOfflineMessageResp) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryOfflineMessageResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QueryOfflineMessageRespValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QueryOfflineMessageRespValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QueryOfflineMessageRespValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Total

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryOfflineMessageRespValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryOfflineMessageRespValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryOfflineMessageRespValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QueryOfflineMessageRespMultiError(errors)
	}

	return nil
}

// QueryOfflineMessageRespMultiError is an error wrapping multiple validation
// errors returned by QueryOfflineMessageResp.ValidateAll() if the designated
// constraints aren't met.
type QueryOfflineMessageRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryOfflineMessageRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryOfflineMessageRespMultiError) AllErrors() []error { return m }

// QueryOfflineMessageRespValidationError is the validation error returned by
// QueryOfflineMessageResp.Validate if the designated constraints aren't met.
type QueryOfflineMessageRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryOfflineMessageRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryOfflineMessageRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryOfflineMessageRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryOfflineMessageRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryOfflineMessageRespValidationError) ErrorName() string {
	return "QueryOfflineMessageRespValidationError"
}

// Error satisfies the builtin error interface
func (e QueryOfflineMessageRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryOfflineMessageResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryOfflineMessageRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryOfflineMessageRespValidationError{}
