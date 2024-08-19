/*
Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connect

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FailureMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureMessage{}

// FailureMessage struct for FailureMessage
type FailureMessage struct {
	Message string `json:"message"`
}

type _FailureMessage FailureMessage

// NewFailureMessage instantiates a new FailureMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureMessage(message string) *FailureMessage {
	this := FailureMessage{}
	this.Message = message
	return &this
}

// NewFailureMessageWithDefaults instantiates a new FailureMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureMessageWithDefaults() *FailureMessage {
	this := FailureMessage{}
	return &this
}

// GetMessage returns the Message field value
func (o *FailureMessage) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FailureMessage) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FailureMessage) SetMessage(v string) {
	o.Message = v
}

func (o FailureMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *FailureMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFailureMessage := _FailureMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFailureMessage)

	if err != nil {
		return err
	}

	*o = FailureMessage(varFailureMessage)

	return err
}

type NullableFailureMessage struct {
	value *FailureMessage
	isSet bool
}

func (v NullableFailureMessage) Get() *FailureMessage {
	return v.value
}

func (v *NullableFailureMessage) Set(val *FailureMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureMessage(val *FailureMessage) *NullableFailureMessage {
	return &NullableFailureMessage{value: val, isSet: true}
}

func (v NullableFailureMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

