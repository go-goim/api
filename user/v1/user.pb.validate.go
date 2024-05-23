// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/user.proto

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

// Validate checks the field values on UserResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserResponseMultiError, or
// nil if none found.
func (m *UserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserResponseMultiError(errors)
	}

	return nil
}

// UserResponseMultiError is an error wrapping multiple validation errors
// returned by UserResponse.ValidateAll() if the designated constraints aren't met.
type UserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserResponseMultiError) AllErrors() []error { return m }

// UserResponseValidationError is the validation error returned by
// UserResponse.Validate if the designated constraints aren't met.
type UserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserResponseValidationError) ErrorName() string { return "UserResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserResponseValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Name

	// no validation rules for Avatar

	// no validation rules for Password

	// no validation rules for LoginStatus

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if m.Email != nil {
		// no validation rules for Email
	}

	if m.Phone != nil {
		// no validation rules for Phone
	}

	if m.PushServerIp != nil {
		// no validation rules for PushServerIp
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on UserList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserListMultiError, or nil
// if none found.
func (m *UserList) ValidateAll() error {
	return m.validate(true)
}

func (m *UserList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserListMultiError(errors)
	}

	return nil
}

// UserListMultiError is an error wrapping multiple validation errors returned
// by UserList.ValidateAll() if the designated constraints aren't met.
type UserListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListMultiError) AllErrors() []error { return m }

// UserListValidationError is the validation error returned by
// UserList.Validate if the designated constraints aren't met.
type UserListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListValidationError) ErrorName() string { return "UserListValidationError" }

// Error satisfies the builtin error interface
func (e UserListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListValidationError{}

// Validate checks the field values on GetUserInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserInfoRequestMultiError, or nil if none found.
func (m *GetUserInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUid() <= 0 {
		err := GetUserInfoRequestValidationError{
			field:  "Uid",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserInfoRequestMultiError(errors)
	}

	return nil
}

// GetUserInfoRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserInfoRequestMultiError) AllErrors() []error { return m }

// GetUserInfoRequestValidationError is the validation error returned by
// GetUserInfoRequest.Validate if the designated constraints aren't met.
type GetUserInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserInfoRequestValidationError) ErrorName() string {
	return "GetUserInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserInfoRequestValidationError{}

// Validate checks the field values on QueryUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryUserRequestMultiError, or nil if none found.
func (m *QueryUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofFieldPresent := false
	switch v := m.Field.(type) {
	case *QueryUserRequest_Email:
		if v == nil {
			err := QueryUserRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = QueryUserRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *QueryUserRequest_Phone:
		if v == nil {
			err := QueryUserRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if !_QueryUserRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := QueryUserRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^1[3-9]\\\\d{9}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofFieldPresent {
		err := QueryUserRequestValidationError{
			field:  "Field",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QueryUserRequestMultiError(errors)
	}

	return nil
}

func (m *QueryUserRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *QueryUserRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// QueryUserRequestMultiError is an error wrapping multiple validation errors
// returned by QueryUserRequest.ValidateAll() if the designated constraints
// aren't met.
type QueryUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryUserRequestMultiError) AllErrors() []error { return m }

// QueryUserRequestValidationError is the validation error returned by
// QueryUserRequest.Validate if the designated constraints aren't met.
type QueryUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryUserRequestValidationError) ErrorName() string { return "QueryUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e QueryUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryUserRequestValidationError{}

var _QueryUserRequest_Phone_Pattern = regexp.MustCompile("^1[3-9]\\d{9}$")

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 20 {
		err := CreateUserRequestValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 20 {
		err := CreateUserRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofFieldPresent := false
	switch v := m.Field.(type) {
	case *CreateUserRequest_Email:
		if v == nil {
			err := CreateUserRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = CreateUserRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *CreateUserRequest_Phone:
		if v == nil {
			err := CreateUserRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if !_CreateUserRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := CreateUserRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^1[3-9]\\\\d{9}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofFieldPresent {
		err := CreateUserRequestValidationError{
			field:  "Field",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

func (m *CreateUserRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateUserRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

var _CreateUserRequest_Phone_Pattern = regexp.MustCompile("^1[3-9]\\d{9}$")

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUid() <= 0 {
		err := UpdateUserRequestValidationError{
			field:  "Uid",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetName() != "" {

		if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 20 {
			err := UpdateUserRequestValidationError{
				field:  "Name",
				reason: "value length must be between 2 and 20 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetEmail() != "" {

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = UpdateUserRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPhone() != "" {

		if !_UpdateUserRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := UpdateUserRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^1[3-9]\\\\d{9}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAvatar() != "" {

		if utf8.RuneCountInString(m.GetAvatar()) > 128 {
			err := UpdateUserRequestValidationError{
				field:  "Avatar",
				reason: "value length must be at most 128 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPassword() != "" {

		if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 20 {
			err := UpdateUserRequestValidationError{
				field:  "Password",
				reason: "value length must be between 6 and 20 runes, inclusive",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateUserRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateUserRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

var _UpdateUserRequest_Phone_Pattern = regexp.MustCompile("^1[3-9]\\d{9}$")

// Validate checks the field values on UserLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserLoginRequestMultiError, or nil if none found.
func (m *UserLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Password

	// no validation rules for LoginType

	oneofFieldPresent := false
	switch v := m.Field.(type) {
	case *UserLoginRequest_Email:
		if v == nil {
			err := UserLoginRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if err := m._validateEmail(m.GetEmail()); err != nil {
			err = UserLoginRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *UserLoginRequest_Phone:
		if v == nil {
			err := UserLoginRequestValidationError{
				field:  "Field",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofFieldPresent = true

		if !_UserLoginRequest_Phone_Pattern.MatchString(m.GetPhone()) {
			err := UserLoginRequestValidationError{
				field:  "Phone",
				reason: "value does not match regex pattern \"^1[3-9]\\\\d{9}$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofFieldPresent {
		err := UserLoginRequestValidationError{
			field:  "Field",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserLoginRequestMultiError(errors)
	}

	return nil
}

func (m *UserLoginRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UserLoginRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// UserLoginRequestMultiError is an error wrapping multiple validation errors
// returned by UserLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type UserLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserLoginRequestMultiError) AllErrors() []error { return m }

// UserLoginRequestValidationError is the validation error returned by
// UserLoginRequest.Validate if the designated constraints aren't met.
type UserLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserLoginRequestValidationError) ErrorName() string { return "UserLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserLoginRequestValidationError{}

var _UserLoginRequest_Phone_Pattern = regexp.MustCompile("^1[3-9]\\d{9}$")
