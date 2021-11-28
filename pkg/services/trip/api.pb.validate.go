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

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Country

	// no validation rules for Rating

	// no validation rules for Description

	// no validation rules for Day

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

	// no validation rules for Id

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

	for idx, item := range m.GetSights() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TripValidationError{
						field:  fmt.Sprintf("Sights[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TripValidationError{
						field:  fmt.Sprintf("Sights[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TripValidationError{
					field:  fmt.Sprintf("Sights[%v]", idx),
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

// Validate checks the field values on AlbumRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AlbumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AlbumRequestMultiError, or
// nil if none found.
func (m *AlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAlbumId() <= 0 {
		err := AlbumRequestValidationError{
			field:  "AlbumId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := AlbumRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AlbumRequestMultiError(errors)
	}
	return nil
}

// AlbumRequestMultiError is an error wrapping multiple validation errors
// returned by AlbumRequest.ValidateAll() if the designated constraints aren't met.
type AlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AlbumRequestMultiError) AllErrors() []error { return m }

// AlbumRequestValidationError is the validation error returned by
// AlbumRequest.Validate if the designated constraints aren't met.
type AlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlbumRequestValidationError) ErrorName() string { return "AlbumRequestValidationError" }

// Error satisfies the builtin error interface
func (e AlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlbumRequestValidationError{}

// Validate checks the field values on Album with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Album) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Album with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AlbumMultiError, or nil if none found.
func (m *Album) ValidateAll() error {
	return m.validate(true)
}

func (m *Album) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TripId

	// no validation rules for Author

	// no validation rules for Title

	// no validation rules for Description

	if len(errors) > 0 {
		return AlbumMultiError(errors)
	}
	return nil
}

// AlbumMultiError is an error wrapping multiple validation errors returned by
// Album.ValidateAll() if the designated constraints aren't met.
type AlbumMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AlbumMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AlbumMultiError) AllErrors() []error { return m }

// AlbumValidationError is the validation error returned by Album.Validate if
// the designated constraints aren't met.
type AlbumValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlbumValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlbumValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlbumValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlbumValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlbumValidationError) ErrorName() string { return "AlbumValidationError" }

// Error satisfies the builtin error interface
func (e AlbumValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlbum.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlbumValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlbumValidationError{}

// Validate checks the field values on ModifyAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ModifyAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ModifyAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ModifyAlbumRequestMultiError, or nil if none found.
func (m *ModifyAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ModifyAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAlbum()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ModifyAlbumRequestValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ModifyAlbumRequestValidationError{
					field:  "Album",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlbum()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ModifyAlbumRequestValidationError{
				field:  "Album",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	if len(errors) > 0 {
		return ModifyAlbumRequestMultiError(errors)
	}
	return nil
}

// ModifyAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by ModifyAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type ModifyAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ModifyAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ModifyAlbumRequestMultiError) AllErrors() []error { return m }

// ModifyAlbumRequestValidationError is the validation error returned by
// ModifyAlbumRequest.Validate if the designated constraints aren't met.
type ModifyAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModifyAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModifyAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModifyAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModifyAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModifyAlbumRequestValidationError) ErrorName() string {
	return "ModifyAlbumRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ModifyAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModifyAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModifyAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModifyAlbumRequestValidationError{}

// Validate checks the field values on SightsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SightsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SightsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SightsRequestMultiError, or
// nil if none found.
func (m *SightsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SightsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TripId

	if len(errors) > 0 {
		return SightsRequestMultiError(errors)
	}
	return nil
}

// SightsRequestMultiError is an error wrapping multiple validation errors
// returned by SightsRequest.ValidateAll() if the designated constraints
// aren't met.
type SightsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SightsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SightsRequestMultiError) AllErrors() []error { return m }

// SightsRequestValidationError is the validation error returned by
// SightsRequest.Validate if the designated constraints aren't met.
type SightsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SightsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SightsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SightsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SightsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SightsRequestValidationError) ErrorName() string { return "SightsRequestValidationError" }

// Error satisfies the builtin error interface
func (e SightsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSightsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SightsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SightsRequestValidationError{}

// Validate checks the field values on Sights with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Sights) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sights with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SightsMultiError, or nil if none found.
func (m *Sights) ValidateAll() error {
	return m.validate(true)
}

func (m *Sights) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SightsMultiError(errors)
	}
	return nil
}

// SightsMultiError is an error wrapping multiple validation errors returned by
// Sights.ValidateAll() if the designated constraints aren't met.
type SightsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SightsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SightsMultiError) AllErrors() []error { return m }

// SightsValidationError is the validation error returned by Sights.Validate if
// the designated constraints aren't met.
type SightsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SightsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SightsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SightsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SightsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SightsValidationError) ErrorName() string { return "SightsValidationError" }

// Error satisfies the builtin error interface
func (e SightsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSights.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SightsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SightsValidationError{}

// Validate checks the field values on ByUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ByUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ByUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ByUserRequestMultiError, or
// nil if none found.
func (m *ByUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ByUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return ByUserRequestMultiError(errors)
	}
	return nil
}

// ByUserRequestMultiError is an error wrapping multiple validation errors
// returned by ByUserRequest.ValidateAll() if the designated constraints
// aren't met.
type ByUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ByUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ByUserRequestMultiError) AllErrors() []error { return m }

// ByUserRequestValidationError is the validation error returned by
// ByUserRequest.Validate if the designated constraints aren't met.
type ByUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ByUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ByUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ByUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ByUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ByUserRequestValidationError) ErrorName() string { return "ByUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e ByUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sByUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ByUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ByUserRequestValidationError{}

// Validate checks the field values on Trips with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Trips) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Trips with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TripsMultiError, or nil if none found.
func (m *Trips) ValidateAll() error {
	return m.validate(true)
}

func (m *Trips) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTrips() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TripsValidationError{
						field:  fmt.Sprintf("Trips[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TripsValidationError{
						field:  fmt.Sprintf("Trips[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TripsValidationError{
					field:  fmt.Sprintf("Trips[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TripsMultiError(errors)
	}
	return nil
}

// TripsMultiError is an error wrapping multiple validation errors returned by
// Trips.ValidateAll() if the designated constraints aren't met.
type TripsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TripsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TripsMultiError) AllErrors() []error { return m }

// TripsValidationError is the validation error returned by Trips.Validate if
// the designated constraints aren't met.
type TripsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TripsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TripsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TripsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TripsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TripsValidationError) ErrorName() string { return "TripsValidationError" }

// Error satisfies the builtin error interface
func (e TripsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrips.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TripsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TripsValidationError{}

// Validate checks the field values on Albums with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Albums) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Albums with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AlbumsMultiError, or nil if none found.
func (m *Albums) ValidateAll() error {
	return m.validate(true)
}

func (m *Albums) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAlbums() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AlbumsValidationError{
						field:  fmt.Sprintf("Albums[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AlbumsValidationError{
						field:  fmt.Sprintf("Albums[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AlbumsValidationError{
					field:  fmt.Sprintf("Albums[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AlbumsMultiError(errors)
	}
	return nil
}

// AlbumsMultiError is an error wrapping multiple validation errors returned by
// Albums.ValidateAll() if the designated constraints aren't met.
type AlbumsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AlbumsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AlbumsMultiError) AllErrors() []error { return m }

// AlbumsValidationError is the validation error returned by Albums.Validate if
// the designated constraints aren't met.
type AlbumsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlbumsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlbumsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlbumsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlbumsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlbumsValidationError) ErrorName() string { return "AlbumsValidationError" }

// Error satisfies the builtin error interface
func (e AlbumsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlbums.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlbumsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlbumsValidationError{}
