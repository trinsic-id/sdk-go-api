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

// checks if the GetSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionResponse{}

// GetSessionResponse struct for GetSessionResponse
type GetSessionResponse struct {
	Session Session `json:"session"`
}

type _GetSessionResponse GetSessionResponse

// NewGetSessionResponse instantiates a new GetSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionResponse(session Session) *GetSessionResponse {
	this := GetSessionResponse{}
	this.Session = session
	return &this
}

// NewGetSessionResponseWithDefaults instantiates a new GetSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionResponseWithDefaults() *GetSessionResponse {
	this := GetSessionResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *GetSessionResponse) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *GetSessionResponse) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *GetSessionResponse) SetSession(v Session) {
	o.Session = v
}

func (o GetSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	return toSerialize, nil
}

func (o *GetSessionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSessionResponse := _GetSessionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSessionResponse)

	if err != nil {
		return err
	}

	*o = GetSessionResponse(varGetSessionResponse)

	return err
}

type NullableGetSessionResponse struct {
	value *GetSessionResponse
	isSet bool
}

func (v NullableGetSessionResponse) Get() *GetSessionResponse {
	return v.value
}

func (v *NullableGetSessionResponse) Set(val *GetSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionResponse(val *GetSessionResponse) *NullableGetSessionResponse {
	return &NullableGetSessionResponse{value: val, isSet: true}
}

func (v NullableGetSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

