// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/service.proto

package pb

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

// Validate checks the field values on CallReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CallReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CallReqMultiError, or nil if none found.
func (m *CallReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CallReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMsg()) < 1 {
		err := CallReqValidationError{
			field:  "Msg",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CallReqMultiError(errors)
	}

	return nil
}

// CallReqMultiError is an error wrapping multiple validation errors returned
// by CallReq.ValidateAll() if the designated constraints aren't met.
type CallReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallReqMultiError) AllErrors() []error { return m }

// CallReqValidationError is the validation error returned by CallReq.Validate
// if the designated constraints aren't met.
type CallReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallReqValidationError) ErrorName() string { return "CallReqValidationError" }

// Error satisfies the builtin error interface
func (e CallReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallReqValidationError{}

// Validate checks the field values on CallResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CallResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CallRespMultiError, or nil
// if none found.
func (m *CallResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CallResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return CallRespMultiError(errors)
	}

	return nil
}

// CallRespMultiError is an error wrapping multiple validation errors returned
// by CallResp.ValidateAll() if the designated constraints aren't met.
type CallRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallRespMultiError) AllErrors() []error { return m }

// CallRespValidationError is the validation error returned by
// CallResp.Validate if the designated constraints aren't met.
type CallRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallRespValidationError) ErrorName() string { return "CallRespValidationError" }

// Error satisfies the builtin error interface
func (e CallRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallRespValidationError{}
