// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shared/v1/pagination.proto

package sV1

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

// Validate checks the field values on PaginationRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PaginationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaginationRequestMultiError, or nil if none found.
func (m *PaginationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return PaginationRequestMultiError(errors)
	}

	return nil
}

// PaginationRequestMultiError is an error wrapping multiple validation errors
// returned by PaginationRequest.ValidateAll() if the designated constraints
// aren't met.
type PaginationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationRequestMultiError) AllErrors() []error { return m }

// PaginationRequestValidationError is the validation error returned by
// PaginationRequest.Validate if the designated constraints aren't met.
type PaginationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationRequestValidationError) ErrorName() string {
	return "PaginationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PaginationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationRequestValidationError{}

// Validate checks the field values on PaginationReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PaginationReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PaginationReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PaginationReplyMultiError, or nil if none found.
func (m *PaginationReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PaginationReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalNumber

	// no validation rules for TotalPage

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return PaginationReplyMultiError(errors)
	}

	return nil
}

// PaginationReplyMultiError is an error wrapping multiple validation errors
// returned by PaginationReply.ValidateAll() if the designated constraints
// aren't met.
type PaginationReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationReplyMultiError) AllErrors() []error { return m }

// PaginationReplyValidationError is the validation error returned by
// PaginationReply.Validate if the designated constraints aren't met.
type PaginationReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationReplyValidationError) ErrorName() string { return "PaginationReplyValidationError" }

// Error satisfies the builtin error interface
func (e PaginationReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaginationReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationReplyValidationError{}
