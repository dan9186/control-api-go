// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: control/module/v1/module.proto

package module_pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _module_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SubscribeResponseMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubscribeResponseMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// SubscribeResponseMessageValidationError is the validation error returned by
// SubscribeResponseMessage.Validate if the designated constraints aren't met.
type SubscribeResponseMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeResponseMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeResponseMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeResponseMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeResponseMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeResponseMessageValidationError) ErrorName() string {
	return "SubscribeResponseMessageValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeResponseMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeResponseMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeResponseMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeResponseMessageValidationError{}

// Validate checks the field values on ShutdownRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ShutdownRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ShutdownRequestValidationError is the validation error returned by
// ShutdownRequest.Validate if the designated constraints aren't met.
type ShutdownRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShutdownRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShutdownRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShutdownRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShutdownRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShutdownRequestValidationError) ErrorName() string { return "ShutdownRequestValidationError" }

// Error satisfies the builtin error interface
func (e ShutdownRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShutdownRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShutdownRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShutdownRequestValidationError{}

// Validate checks the field values on ShutdownResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ShutdownResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ShutdownResponseValidationError is the validation error returned by
// ShutdownResponse.Validate if the designated constraints aren't met.
type ShutdownResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShutdownResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShutdownResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShutdownResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShutdownResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShutdownResponseValidationError) ErrorName() string { return "ShutdownResponseValidationError" }

// Error satisfies the builtin error interface
func (e ShutdownResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShutdownResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShutdownResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShutdownResponseValidationError{}