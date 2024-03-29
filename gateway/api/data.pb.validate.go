// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gateway/api/data.proto

package api

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

// Validate checks the field values on RegisterNewUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterNewUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetEmployeeId()); l < 1 || l > 50 {
		return RegisterNewUserRequestValidationError{
			field:  "EmployeeId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 255 {
		return RegisterNewUserRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 8 || l > 255 {
		return RegisterNewUserRequestValidationError{
			field:  "Password",
			reason: "value length must be between 8 and 255 runes, inclusive",
		}
	}

	// no validation rules for Email

	// no validation rules for Role

	// no validation rules for Position

	// no validation rules for Department

	// no validation rules for Gender

	// no validation rules for Address

	// no validation rules for Phone

	return nil
}

// RegisterNewUserRequestValidationError is the validation error returned by
// RegisterNewUserRequest.Validate if the designated constraints aren't met.
type RegisterNewUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterNewUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterNewUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterNewUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterNewUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterNewUserRequestValidationError) ErrorName() string {
	return "RegisterNewUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterNewUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterNewUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterNewUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterNewUserRequestValidationError{}

// Validate checks the field values on UpdateCitizenIdentityCardRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *UpdateCitizenIdentityCardRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateCitizenIdentityCardRequestValidationError is the validation error
// returned by UpdateCitizenIdentityCardRequest.Validate if the designated
// constraints aren't met.
type UpdateCitizenIdentityCardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCitizenIdentityCardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCitizenIdentityCardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCitizenIdentityCardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCitizenIdentityCardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCitizenIdentityCardRequestValidationError) ErrorName() string {
	return "UpdateCitizenIdentityCardRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCitizenIdentityCardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCitizenIdentityCardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCitizenIdentityCardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCitizenIdentityCardRequestValidationError{}

// Validate checks the field values on UpdateCitizenIdentityCardResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *UpdateCitizenIdentityCardResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCitizenIdentityCardResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateCitizenIdentityCardResponseValidationError is the validation error
// returned by UpdateCitizenIdentityCardResponse.Validate if the designated
// constraints aren't met.
type UpdateCitizenIdentityCardResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCitizenIdentityCardResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCitizenIdentityCardResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCitizenIdentityCardResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCitizenIdentityCardResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCitizenIdentityCardResponseValidationError) ErrorName() string {
	return "UpdateCitizenIdentityCardResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCitizenIdentityCardResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCitizenIdentityCardResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCitizenIdentityCardResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCitizenIdentityCardResponseValidationError{}

// Validate checks the field values on CitizenIdentityCard with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CitizenIdentityCard) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if m.GetRegisterDate() <= 0 {
		return CitizenIdentityCardValidationError{
			field:  "RegisterDate",
			reason: "value must be greater than 0",
		}
	}

	if m.GetExpireDate() <= 0 {
		return CitizenIdentityCardValidationError{
			field:  "ExpireDate",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Provider

	if m.GetBorn() <= 0 {
		return CitizenIdentityCardValidationError{
			field:  "Born",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CitizenIdentityCardValidationError is the validation error returned by
// CitizenIdentityCard.Validate if the designated constraints aren't met.
type CitizenIdentityCardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CitizenIdentityCardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CitizenIdentityCardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CitizenIdentityCardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CitizenIdentityCardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CitizenIdentityCardValidationError) ErrorName() string {
	return "CitizenIdentityCardValidationError"
}

// Error satisfies the builtin error interface
func (e CitizenIdentityCardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCitizenIdentityCard.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CitizenIdentityCardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CitizenIdentityCardValidationError{}

// Validate checks the field values on RegisterNewUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterNewUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Token

	return nil
}

// RegisterNewUserResponseValidationError is the validation error returned by
// RegisterNewUserResponse.Validate if the designated constraints aren't met.
type RegisterNewUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterNewUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterNewUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterNewUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterNewUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterNewUserResponseValidationError) ErrorName() string {
	return "RegisterNewUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterNewUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterNewUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterNewUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterNewUserResponseValidationError{}

// Validate checks the field values on AuthorizeUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthorizeUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Email

	if l := utf8.RuneCountInString(m.GetEmployeeId()); l < 1 || l > 10 {
		return AuthorizeUserRequestValidationError{
			field:  "EmployeeId",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 8 || l > 255 {
		return AuthorizeUserRequestValidationError{
			field:  "Password",
			reason: "value length must be between 8 and 255 runes, inclusive",
		}
	}

	return nil
}

// AuthorizeUserRequestValidationError is the validation error returned by
// AuthorizeUserRequest.Validate if the designated constraints aren't met.
type AuthorizeUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeUserRequestValidationError) ErrorName() string {
	return "AuthorizeUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizeUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeUserRequestValidationError{}

// Validate checks the field values on AuthorizeUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthorizeUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Token

	return nil
}

// AuthorizeUserResponseValidationError is the validation error returned by
// AuthorizeUserResponse.Validate if the designated constraints aren't met.
type AuthorizeUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeUserResponseValidationError) ErrorName() string {
	return "AuthorizeUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizeUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeUserResponseValidationError{}

// Validate checks the field values on RegisterCICForUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterCICForUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return RegisterCICForUserRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Front

	// no validation rules for Back

	return nil
}

// RegisterCICForUserRequestValidationError is the validation error returned by
// RegisterCICForUserRequest.Validate if the designated constraints aren't met.
type RegisterCICForUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterCICForUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterCICForUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterCICForUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterCICForUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterCICForUserRequestValidationError) ErrorName() string {
	return "RegisterCICForUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterCICForUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterCICForUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterCICForUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterCICForUserRequestValidationError{}

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Image) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Data

	return nil
}

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on RegisterCICForUserResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RegisterCICForUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterCICForUserResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RegisterCICForUserResponseValidationError is the validation error returned
// by RegisterCICForUserResponse.Validate if the designated constraints aren't met.
type RegisterCICForUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterCICForUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterCICForUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterCICForUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterCICForUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterCICForUserResponseValidationError) ErrorName() string {
	return "RegisterCICForUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterCICForUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterCICForUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterCICForUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterCICForUserResponseValidationError{}
