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

// ResultCollectionMethod the model 'ResultCollectionMethod'
type ResultCollectionMethod string

// List of ResultCollectionMethod
const (
	RESULTCOLLECTIONMETHOD_POLL_RESULT ResultCollectionMethod = "PollResult"
	RESULTCOLLECTIONMETHOD_CAPTURE_REDIRECT ResultCollectionMethod = "CaptureRedirect"
)

// All allowed values of ResultCollectionMethod enum
var AllowedResultCollectionMethodEnumValues = []ResultCollectionMethod{
	"PollResult",
	"CaptureRedirect",
}

func (v *ResultCollectionMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResultCollectionMethod(value)
	for _, existing := range AllowedResultCollectionMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResultCollectionMethod", value)
}

// NewResultCollectionMethodFromValue returns a pointer to a valid ResultCollectionMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResultCollectionMethodFromValue(v string) (*ResultCollectionMethod, error) {
	ev := ResultCollectionMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResultCollectionMethod: valid values are %v", v, AllowedResultCollectionMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResultCollectionMethod) IsValid() bool {
	for _, existing := range AllowedResultCollectionMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResultCollectionMethod value
func (v ResultCollectionMethod) Ptr() *ResultCollectionMethod {
	return &v
}

type NullableResultCollectionMethod struct {
	value *ResultCollectionMethod
	isSet bool
}

func (v NullableResultCollectionMethod) Get() *ResultCollectionMethod {
	return v.value
}

func (v *NullableResultCollectionMethod) Set(val *ResultCollectionMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCollectionMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCollectionMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCollectionMethod(val *ResultCollectionMethod) *NullableResultCollectionMethod {
	return &NullableResultCollectionMethod{value: val, isSet: true}
}

func (v NullableResultCollectionMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCollectionMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

