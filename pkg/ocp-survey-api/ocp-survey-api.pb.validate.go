// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-survey-api/ocp-survey-api.proto

package ocp_survey_api

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
)

// Validate checks the field values on Survey with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Survey) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Link

	return nil
}

// SurveyValidationError is the validation error returned by Survey.Validate if
// the designated constraints aren't met.
type SurveyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SurveyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SurveyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SurveyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SurveyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SurveyValidationError) ErrorName() string { return "SurveyValidationError" }

// Error satisfies the builtin error interface
func (e SurveyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSurvey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SurveyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SurveyValidationError{}

// Validate checks the field values on CreateSurveyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSurveyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateSurveyV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Link

	return nil
}

// CreateSurveyV1RequestValidationError is the validation error returned by
// CreateSurveyV1Request.Validate if the designated constraints aren't met.
type CreateSurveyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSurveyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSurveyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSurveyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSurveyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSurveyV1RequestValidationError) ErrorName() string {
	return "CreateSurveyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSurveyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSurveyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSurveyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSurveyV1RequestValidationError{}

// Validate checks the field values on CreateSurveyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSurveyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SurveyId

	return nil
}

// CreateSurveyV1ResponseValidationError is the validation error returned by
// CreateSurveyV1Response.Validate if the designated constraints aren't met.
type CreateSurveyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSurveyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSurveyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSurveyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSurveyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSurveyV1ResponseValidationError) ErrorName() string {
	return "CreateSurveyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSurveyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSurveyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSurveyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSurveyV1ResponseValidationError{}

// Validate checks the field values on MultiCreateSurveyV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateSurveyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSurveys()) < 1 {
		return MultiCreateSurveyV1RequestValidationError{
			field:  "Surveys",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetSurveys() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateSurveyV1RequestValidationError{
					field:  fmt.Sprintf("Surveys[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateSurveyV1RequestValidationError is the validation error returned
// by MultiCreateSurveyV1Request.Validate if the designated constraints aren't met.
type MultiCreateSurveyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateSurveyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateSurveyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateSurveyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateSurveyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateSurveyV1RequestValidationError) ErrorName() string {
	return "MultiCreateSurveyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateSurveyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateSurveyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateSurveyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateSurveyV1RequestValidationError{}

// Validate checks the field values on MultiCreateSurveyV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateSurveyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiCreateSurveyV1ResponseValidationError is the validation error returned
// by MultiCreateSurveyV1Response.Validate if the designated constraints
// aren't met.
type MultiCreateSurveyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateSurveyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateSurveyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateSurveyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateSurveyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateSurveyV1ResponseValidationError) ErrorName() string {
	return "MultiCreateSurveyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateSurveyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateSurveyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateSurveyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateSurveyV1ResponseValidationError{}

// Validate checks the field values on DescribeSurveyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSurveyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSurveyId() <= 0 {
		return DescribeSurveyV1RequestValidationError{
			field:  "SurveyId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeSurveyV1RequestValidationError is the validation error returned by
// DescribeSurveyV1Request.Validate if the designated constraints aren't met.
type DescribeSurveyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSurveyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSurveyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSurveyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSurveyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSurveyV1RequestValidationError) ErrorName() string {
	return "DescribeSurveyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSurveyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSurveyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSurveyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSurveyV1RequestValidationError{}

// Validate checks the field values on DescribeSurveyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSurveyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSurvey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeSurveyV1ResponseValidationError{
				field:  "Survey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeSurveyV1ResponseValidationError is the validation error returned by
// DescribeSurveyV1Response.Validate if the designated constraints aren't met.
type DescribeSurveyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSurveyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSurveyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSurveyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSurveyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSurveyV1ResponseValidationError) ErrorName() string {
	return "DescribeSurveyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSurveyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSurveyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSurveyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSurveyV1ResponseValidationError{}

// Validate checks the field values on ListSurveysV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSurveysV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListSurveysV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	if m.GetOffset() < 0 {
		return ListSurveysV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListSurveysV1RequestValidationError is the validation error returned by
// ListSurveysV1Request.Validate if the designated constraints aren't met.
type ListSurveysV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSurveysV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSurveysV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSurveysV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSurveysV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSurveysV1RequestValidationError) ErrorName() string {
	return "ListSurveysV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSurveysV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSurveysV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSurveysV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSurveysV1RequestValidationError{}

// Validate checks the field values on ListSurveysV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSurveysV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSurveys() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSurveysV1ResponseValidationError{
					field:  fmt.Sprintf("Surveys[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListSurveysV1ResponseValidationError is the validation error returned by
// ListSurveysV1Response.Validate if the designated constraints aren't met.
type ListSurveysV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSurveysV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSurveysV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSurveysV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSurveysV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSurveysV1ResponseValidationError) ErrorName() string {
	return "ListSurveysV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSurveysV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSurveysV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSurveysV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSurveysV1ResponseValidationError{}

// Validate checks the field values on UpdateSurveyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateSurveyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSurvey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSurveyV1RequestValidationError{
				field:  "Survey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateSurveyV1RequestValidationError is the validation error returned by
// UpdateSurveyV1Request.Validate if the designated constraints aren't met.
type UpdateSurveyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSurveyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSurveyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSurveyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSurveyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSurveyV1RequestValidationError) ErrorName() string {
	return "UpdateSurveyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSurveyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSurveyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSurveyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSurveyV1RequestValidationError{}

// Validate checks the field values on UpdateSurveyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateSurveyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateSurveyV1ResponseValidationError is the validation error returned by
// UpdateSurveyV1Response.Validate if the designated constraints aren't met.
type UpdateSurveyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSurveyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSurveyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSurveyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSurveyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSurveyV1ResponseValidationError) ErrorName() string {
	return "UpdateSurveyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSurveyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSurveyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSurveyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSurveyV1ResponseValidationError{}

// Validate checks the field values on RemoveSurveyV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSurveyV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSurveyId() <= 0 {
		return RemoveSurveyV1RequestValidationError{
			field:  "SurveyId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveSurveyV1RequestValidationError is the validation error returned by
// RemoveSurveyV1Request.Validate if the designated constraints aren't met.
type RemoveSurveyV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSurveyV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSurveyV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSurveyV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSurveyV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSurveyV1RequestValidationError) ErrorName() string {
	return "RemoveSurveyV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSurveyV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSurveyV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSurveyV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSurveyV1RequestValidationError{}

// Validate checks the field values on RemoveSurveyV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSurveyV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveSurveyV1ResponseValidationError is the validation error returned by
// RemoveSurveyV1Response.Validate if the designated constraints aren't met.
type RemoveSurveyV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSurveyV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSurveyV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSurveyV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSurveyV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSurveyV1ResponseValidationError) ErrorName() string {
	return "RemoveSurveyV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSurveyV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSurveyV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSurveyV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSurveyV1ResponseValidationError{}
