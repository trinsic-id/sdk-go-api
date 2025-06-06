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

// checks if the CancelSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelSessionResponse{}

// CancelSessionResponse struct for CancelSessionResponse
type CancelSessionResponse struct {
	Session Session `json:"session"`
}

type _CancelSessionResponse CancelSessionResponse

// NewCancelSessionResponse instantiates a new CancelSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelSessionResponse(session Session) *CancelSessionResponse {
	this := CancelSessionResponse{}
	this.Session = session
	return &this
}

// NewCancelSessionResponseWithDefaults instantiates a new CancelSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelSessionResponseWithDefaults() *CancelSessionResponse {
	this := CancelSessionResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *CancelSessionResponse) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *CancelSessionResponse) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *CancelSessionResponse) SetSession(v Session) {
	o.Session = v
}

func (o CancelSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	return toSerialize, nil
}

func (o *CancelSessionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"session",
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

	varCancelSessionResponse := _CancelSessionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCancelSessionResponse)

	if err != nil {
		return err
	}

	*o = CancelSessionResponse(varCancelSessionResponse)

	return err
}

type NullableCancelSessionResponse struct {
	value *CancelSessionResponse
	isSet bool
}

func (v NullableCancelSessionResponse) Get() *CancelSessionResponse {
	return v.value
}

func (v *NullableCancelSessionResponse) Set(val *CancelSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelSessionResponse(val *CancelSessionResponse) *NullableCancelSessionResponse {
	return &NullableCancelSessionResponse{value: val, isSet: true}
}

func (v NullableCancelSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


