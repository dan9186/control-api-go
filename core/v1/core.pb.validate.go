// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: control/core/v1/core.proto

package core_pb

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
var _core_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SubscribeRequestMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubscribeRequestMessage) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return SubscribeRequestMessageValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return SubscribeRequestMessageValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetListener() == nil {
		return SubscribeRequestMessageValidationError{
			field:  "Listener",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetListener()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscribeRequestMessageValidationError{
				field:  "Listener",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

func (m *SubscribeRequestMessage) _validateUuid(uuid string) error {
	if matched := _core_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SubscribeRequestMessageValidationError is the validation error returned by
// SubscribeRequestMessage.Validate if the designated constraints aren't met.
type SubscribeRequestMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeRequestMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeRequestMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeRequestMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeRequestMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeRequestMessageValidationError) ErrorName() string {
	return "SubscribeRequestMessageValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeRequestMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeRequestMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeRequestMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeRequestMessageValidationError{}

// Validate checks the field values on Listener with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener) Validate() error {
	if m == nil {
		return nil
	}

	if val := m.GetPort(); val <= 2000 || val > 65353 {
		return ListenerValidationError{
			field:  "Port",
			reason: "value must be inside range (2000, 65353]",
		}
	}

	// no validation rules for MatchRegex

	return nil
}

// ListenerValidationError is the validation error returned by
// Listener.Validate if the designated constraints aren't met.
type ListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerValidationError) ErrorName() string { return "ListenerValidationError" }

// Error satisfies the builtin error interface
func (e ListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerValidationError{}

// Validate checks the field values on UnsubscribeRequestMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UnsubscribeRequestMessage) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return UnsubscribeRequestMessageValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *UnsubscribeRequestMessage) _validateUuid(uuid string) error {
	if matched := _core_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UnsubscribeRequestMessageValidationError is the validation error returned by
// UnsubscribeRequestMessage.Validate if the designated constraints aren't met.
type UnsubscribeRequestMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnsubscribeRequestMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnsubscribeRequestMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnsubscribeRequestMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnsubscribeRequestMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnsubscribeRequestMessageValidationError) ErrorName() string {
	return "UnsubscribeRequestMessageValidationError"
}

// Error satisfies the builtin error interface
func (e UnsubscribeRequestMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnsubscribeRequestMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnsubscribeRequestMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnsubscribeRequestMessageValidationError{}

// Validate checks the field values on EventMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EventMessage) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEvent() == nil {
		return EventMessageValidationError{
			field:  "Event",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventMessageValidationError{
				field:  "Event",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EventMessageValidationError is the validation error returned by
// EventMessage.Validate if the designated constraints aren't met.
type EventMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventMessageValidationError) ErrorName() string { return "EventMessageValidationError" }

// Error satisfies the builtin error interface
func (e EventMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventMessageValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Event) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return EventValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if wrapper := m.GetResponseId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return EventValidationError{
				field:  "ResponseId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	// no validation rules for Headers

	// no validation rules for Body

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

func (m *Event) _validateUuid(uuid string) error {
	if matched := _core_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}
