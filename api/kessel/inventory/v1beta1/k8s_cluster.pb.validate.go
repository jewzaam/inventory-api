// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kessel/inventory/v1beta1/k8s_cluster.proto

package v1beta1

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

// Validate checks the field values on K8SCluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *K8SCluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on K8SCluster with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in K8SClusterMultiError, or
// nil if none found.
func (m *K8SCluster) ValidateAll() error {
	return m.validate(true)
}

func (m *K8SCluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return K8SClusterValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReporterData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "ReporterData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "ReporterData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReporterData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return K8SClusterValidationError{
				field:  "ReporterData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetReporters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, K8SClusterValidationError{
						field:  fmt.Sprintf("Reporters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, K8SClusterValidationError{
						field:  fmt.Sprintf("Reporters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return K8SClusterValidationError{
					field:  fmt.Sprintf("Reporters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, K8SClusterValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return K8SClusterValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return K8SClusterMultiError(errors)
	}

	return nil
}

// K8SClusterMultiError is an error wrapping multiple validation errors
// returned by K8SCluster.ValidateAll() if the designated constraints aren't met.
type K8SClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m K8SClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m K8SClusterMultiError) AllErrors() []error { return m }

// K8SClusterValidationError is the validation error returned by
// K8SCluster.Validate if the designated constraints aren't met.
type K8SClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e K8SClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e K8SClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e K8SClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e K8SClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e K8SClusterValidationError) ErrorName() string { return "K8SClusterValidationError" }

// Error satisfies the builtin error interface
func (e K8SClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sK8SCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = K8SClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = K8SClusterValidationError{}
