/*
Trinsic API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
	"fmt"
)

// SessionErrorCode the model 'SessionErrorCode'
type SessionErrorCode string

// List of SessionErrorCode
const (
	SESSIONERRORCODE_INTERNAL SessionErrorCode = "Internal"
	SESSIONERRORCODE_OTHER SessionErrorCode = "Other"
	SESSIONERRORCODE_CANCELED SessionErrorCode = "Canceled"
	SESSIONERRORCODE_USER_ABANDONED SessionErrorCode = "UserAbandoned"
	SESSIONERRORCODE_USER_TIMED_OUT SessionErrorCode = "UserTimedOut"
	SESSIONERRORCODE_EXPIRED SessionErrorCode = "Expired"
	SESSIONERRORCODE_INVALID_IMAGE SessionErrorCode = "InvalidImage"
	SESSIONERRORCODE_INAUTHENTIC SessionErrorCode = "Inauthentic"
	SESSIONERRORCODE_UNSUPPORTED_DOCUMENT SessionErrorCode = "UnsupportedDocument"
	SESSIONERRORCODE_ASSURANCE_LEVEL_NOT_MET SessionErrorCode = "AssuranceLevelNotMet"
)

// All allowed values of SessionErrorCode enum
var AllowedSessionErrorCodeEnumValues = []SessionErrorCode{
	"Internal",
	"Other",
	"Canceled",
	"UserAbandoned",
	"UserTimedOut",
	"Expired",
	"InvalidImage",
	"Inauthentic",
	"UnsupportedDocument",
	"AssuranceLevelNotMet",
}

func (v *SessionErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionErrorCode(value)
	for _, existing := range AllowedSessionErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionErrorCode", value)
}

// NewSessionErrorCodeFromValue returns a pointer to a valid SessionErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionErrorCodeFromValue(v string) (*SessionErrorCode, error) {
	ev := SessionErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionErrorCode: valid values are %v", v, AllowedSessionErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionErrorCode) IsValid() bool {
	for _, existing := range AllowedSessionErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionErrorCode value
func (v SessionErrorCode) Ptr() *SessionErrorCode {
	return &v
}

type NullableSessionErrorCode struct {
	value *SessionErrorCode
	isSet bool
}

func (v NullableSessionErrorCode) Get() *SessionErrorCode {
	return v.value
}

func (v *NullableSessionErrorCode) Set(val *SessionErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionErrorCode(val *SessionErrorCode) *NullableSessionErrorCode {
	return &NullableSessionErrorCode{value: val, isSet: true}
}

func (v NullableSessionErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

