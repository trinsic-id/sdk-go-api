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

// checks if the IndonesiaNikInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndonesiaNikInput{}

// IndonesiaNikInput struct for IndonesiaNikInput
type IndonesiaNikInput struct {
	// The user's full name
	FullName string `json:"fullName"`
	// The user's date of birth, in `YYYY-MM-DD` format
	DateOfBirth string `json:"dateOfBirth"`
	// The user's Indonesia NIK ID number
	NikIdNumber string `json:"nikIdNumber" validate:"regexp=^\\\\d*$"`
}

type _IndonesiaNikInput IndonesiaNikInput

// NewIndonesiaNikInput instantiates a new IndonesiaNikInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndonesiaNikInput(fullName string, dateOfBirth string, nikIdNumber string) *IndonesiaNikInput {
	this := IndonesiaNikInput{}
	this.FullName = fullName
	this.DateOfBirth = dateOfBirth
	this.NikIdNumber = nikIdNumber
	return &this
}

// NewIndonesiaNikInputWithDefaults instantiates a new IndonesiaNikInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndonesiaNikInputWithDefaults() *IndonesiaNikInput {
	this := IndonesiaNikInput{}
	return &this
}

// GetFullName returns the FullName field value
func (o *IndonesiaNikInput) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *IndonesiaNikInput) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *IndonesiaNikInput) SetFullName(v string) {
	o.FullName = v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *IndonesiaNikInput) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *IndonesiaNikInput) GetDateOfBirthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *IndonesiaNikInput) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetNikIdNumber returns the NikIdNumber field value
func (o *IndonesiaNikInput) GetNikIdNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NikIdNumber
}

// GetNikIdNumberOk returns a tuple with the NikIdNumber field value
// and a boolean to check if the value has been set.
func (o *IndonesiaNikInput) GetNikIdNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NikIdNumber, true
}

// SetNikIdNumber sets field value
func (o *IndonesiaNikInput) SetNikIdNumber(v string) {
	o.NikIdNumber = v
}

func (o IndonesiaNikInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndonesiaNikInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullName"] = o.FullName
	toSerialize["dateOfBirth"] = o.DateOfBirth
	toSerialize["nikIdNumber"] = o.NikIdNumber
	return toSerialize, nil
}

func (o *IndonesiaNikInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fullName",
		"dateOfBirth",
		"nikIdNumber",
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

	varIndonesiaNikInput := _IndonesiaNikInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIndonesiaNikInput)

	if err != nil {
		return err
	}

	*o = IndonesiaNikInput(varIndonesiaNikInput)

	return err
}

type NullableIndonesiaNikInput struct {
	value *IndonesiaNikInput
	isSet bool
}

func (v NullableIndonesiaNikInput) Get() *IndonesiaNikInput {
	return v.value
}

func (v *NullableIndonesiaNikInput) Set(val *IndonesiaNikInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIndonesiaNikInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIndonesiaNikInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndonesiaNikInput(val *IndonesiaNikInput) *NullableIndonesiaNikInput {
	return &NullableIndonesiaNikInput{value: val, isSet: true}
}

func (v NullableIndonesiaNikInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndonesiaNikInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


