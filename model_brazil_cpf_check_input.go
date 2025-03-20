/*
Trinsic API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BrazilCpfCheckInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrazilCpfCheckInput{}

// BrazilCpfCheckInput struct for BrazilCpfCheckInput
type BrazilCpfCheckInput struct {
	// The user's 11-digit, numeric CPF Number
	CpfNumber string `json:"cpfNumber" validate:"regexp=^\\\\d*$"`
}

type _BrazilCpfCheckInput BrazilCpfCheckInput

// NewBrazilCpfCheckInput instantiates a new BrazilCpfCheckInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrazilCpfCheckInput(cpfNumber string) *BrazilCpfCheckInput {
	this := BrazilCpfCheckInput{}
	this.CpfNumber = cpfNumber
	return &this
}

// NewBrazilCpfCheckInputWithDefaults instantiates a new BrazilCpfCheckInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrazilCpfCheckInputWithDefaults() *BrazilCpfCheckInput {
	this := BrazilCpfCheckInput{}
	return &this
}

// GetCpfNumber returns the CpfNumber field value
func (o *BrazilCpfCheckInput) GetCpfNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CpfNumber
}

// GetCpfNumberOk returns a tuple with the CpfNumber field value
// and a boolean to check if the value has been set.
func (o *BrazilCpfCheckInput) GetCpfNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpfNumber, true
}

// SetCpfNumber sets field value
func (o *BrazilCpfCheckInput) SetCpfNumber(v string) {
	o.CpfNumber = v
}

func (o BrazilCpfCheckInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrazilCpfCheckInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpfNumber"] = o.CpfNumber
	return toSerialize, nil
}

func (o *BrazilCpfCheckInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cpfNumber",
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

	varBrazilCpfCheckInput := _BrazilCpfCheckInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrazilCpfCheckInput)

	if err != nil {
		return err
	}

	*o = BrazilCpfCheckInput(varBrazilCpfCheckInput)

	return err
}

type NullableBrazilCpfCheckInput struct {
	value *BrazilCpfCheckInput
	isSet bool
}

func (v NullableBrazilCpfCheckInput) Get() *BrazilCpfCheckInput {
	return v.value
}

func (v *NullableBrazilCpfCheckInput) Set(val *BrazilCpfCheckInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBrazilCpfCheckInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBrazilCpfCheckInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrazilCpfCheckInput(val *BrazilCpfCheckInput) *NullableBrazilCpfCheckInput {
	return &NullableBrazilCpfCheckInput{value: val, isSet: true}
}

func (v NullableBrazilCpfCheckInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrazilCpfCheckInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


