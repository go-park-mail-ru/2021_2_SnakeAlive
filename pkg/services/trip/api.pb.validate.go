// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/services/trip/api.proto

package trip_service

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

// Validate checks the field values on TripRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TripRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TripRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TripRequestMultiError, or
// nil if none found.
func (m *TripRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TripRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTripId() <= 0 {
		err := TripRequestValidationError{
			field:  "TripId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := TripRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TripRequestMultiError(errors)
	}
	return nil
}

// TripRequestMultiError is an error wrapping multiple validation errors
// returned by TripRequest.ValidateAll() if the designated constraints aren't met.
type TripRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TripRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TripRequestMultiError) AllErrors() []error { return m }

// TripRequestValidationError is the validation error returned by
// TripRequest.Validate if the designated constraints aren't met.
type TripRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TripRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TripRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TripRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TripRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TripRequestValidationError) ErrorName() string { return "TripRequestValidationError" }

// Error satisfies the builtin error interface
func (e TripRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTripRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TripRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TripRequestValidationError{}

// Validate checks the field values on Sight with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Sight) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sight with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SightMultiError, or nil if none found.
func (m *Sight) ValidateAll() error {
	return m.validate(true)
}

func (m *Sight) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := SightValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := SightValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetName()) < 1 {
		err := SightValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCountry()) < 1 {
		err := SightValidationError{
			field:  "Country",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetCountry()) < 1 {
		err := SightValidationError{
			field:  "Country",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Rating

	// no validation rules for Description

	if len(errors) > 0 {
		return SightMultiError(errors)
	}
	return nil
}

// SightMultiError is an error wrapping multiple validation errors returned by
// Sight.ValidateAll() if the designated constraints aren't met.
type SightMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SightMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SightMultiError) AllErrors() []error { return m }

// SightValidationError is the validation error returned by Sight.Validate if
// the designated constraints aren't met.
type SightValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SightValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SightValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SightValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SightValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SightValidationError) ErrorName() string { return "SightValidationError" }

// Error satisfies the builtin error interface
func (e SightValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSight.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SightValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SightValidationError{}

// Validate checks the field values on Day with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Day) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Day with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DayMultiError, or nil if none found.
func (m *Day) ValidateAll() error {
	return m.validate(true)
}

func (m *Day) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSights() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DayValidationError{
						field:  fmt.Sprintf("Sights[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DayValidationError{
						field:  fmt.Sprintf("Sights[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DayValidationError{
					field:  fmt.Sprintf("Sights[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DayMultiError(errors)
	}
	return nil
}

// DayMultiError is an error wrapping multiple validation errors returned by
// Day.ValidateAll() if the designated constraints aren't met.
type DayMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DayMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DayMultiError) AllErrors() []error { return m }

// DayValidationError is the validation error returned by Day.Validate if the
// designated constraints aren't met.
type DayValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DayValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DayValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DayValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DayValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DayValidationError) ErrorName() string { return "DayValidationError" }

// Error satisfies the builtin error interface
func (e DayValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDay.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DayValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DayValidationError{}

// Validate checks the field values on Trip with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Trip) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Trip with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TripMultiError, or nil if none found.
func (m *Trip) ValidateAll() error {
	return m.validate(true)
}

func (m *Trip) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := TripValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := TripValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetTitle()) < 1 {
		err := TripValidationError{
			field:  "Title",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Description

	for idx, item := range m.GetDays() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TripValidationError{
						field:  fmt.Sprintf("Days[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TripValidationError{
						field:  fmt.Sprintf("Days[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TripValidationError{
					field:  fmt.Sprintf("Days[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TripMultiError(errors)
	}
	return nil
}

// TripMultiError is an error wrapping multiple validation errors returned by
// Trip.ValidateAll() if the designated constraints aren't met.
type TripMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TripMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TripMultiError) AllErrors() []error { return m }

// TripValidationError is the validation error returned by Trip.Validate if the
// designated constraints aren't met.
type TripValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TripValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TripValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TripValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TripValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TripValidationError) ErrorName() string { return "TripValidationError" }

// Error satisfies the builtin error interface
func (e TripValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrip.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TripValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TripValidationError{}

// Validate checks the field values on ModifyTripRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ModifyTripRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ModifyTripRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ModifyTripRequestMultiError, or nil if none found.
func (m *ModifyTripRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ModifyTripRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTrip()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ModifyTripRequestValidationError{
					field:  "Trip",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ModifyTripRequestValidationError{
					field:  "Trip",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTrip()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ModifyTripRequestValidationError{
				field:  "Trip",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetUserId() <= 0 {
		err := ModifyTripRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ModifyTripRequestMultiError(errors)
	}
	return nil
}

// ModifyTripRequestMultiError is an error wrapping multiple validation errors
// returned by ModifyTripRequest.ValidateAll() if the designated constraints
// aren't met.
type ModifyTripRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ModifyTripRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ModifyTripRequestMultiError) AllErrors() []error { return m }

// ModifyTripRequestValidationError is the validation error returned by
// ModifyTripRequest.Validate if the designated constraints aren't met.
type ModifyTripRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModifyTripRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModifyTripRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModifyTripRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModifyTripRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModifyTripRequestValidationError) ErrorName() string {
	return "ModifyTripRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ModifyTripRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModifyTripRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModifyTripRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModifyTripRequestValidationError{}
