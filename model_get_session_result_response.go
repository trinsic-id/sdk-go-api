/*
Connect API

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

// checks if the GetSessionResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSessionResultResponse{}

// GetSessionResultResponse struct for GetSessionResultResponse
type GetSessionResultResponse struct {
	Session Session `json:"session"`
	IdentityData *IdentityData `json:"identityData,omitempty"`
}

type _GetSessionResultResponse GetSessionResultResponse

// NewGetSessionResultResponse instantiates a new GetSessionResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSessionResultResponse(session Session) *GetSessionResultResponse {
	this := GetSessionResultResponse{}
	this.Session = session
	return &this
}

// NewGetSessionResultResponseWithDefaults instantiates a new GetSessionResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSessionResultResponseWithDefaults() *GetSessionResultResponse {
	this := GetSessionResultResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *GetSessionResultResponse) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *GetSessionResultResponse) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *GetSessionResultResponse) SetSession(v Session) {
	o.Session = v
}

// GetIdentityData returns the IdentityData field value if set, zero value otherwise.
func (o *GetSessionResultResponse) GetIdentityData() IdentityData {
	if o == nil || IsNil(o.IdentityData) {
		var ret IdentityData
		return ret
	}
	return *o.IdentityData
}

// GetIdentityDataOk returns a tuple with the IdentityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSessionResultResponse) GetIdentityDataOk() (*IdentityData, bool) {
	if o == nil || IsNil(o.IdentityData) {
		return nil, false
	}
	return o.IdentityData, true
}

// HasIdentityData returns a boolean if a field has been set.
func (o *GetSessionResultResponse) HasIdentityData() bool {
	if o != nil && !IsNil(o.IdentityData) {
		return true
	}

	return false
}

// SetIdentityData gets a reference to the given IdentityData and assigns it to the IdentityData field.
func (o *GetSessionResultResponse) SetIdentityData(v IdentityData) {
	o.IdentityData = &v
}

func (o GetSessionResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSessionResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	if !IsNil(o.IdentityData) {
		toSerialize["identityData"] = o.IdentityData
	}
	return toSerialize, nil
}

func (o *GetSessionResultResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetSessionResultResponse := _GetSessionResultResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSessionResultResponse)

	if err != nil {
		return err
	}

	*o = GetSessionResultResponse(varGetSessionResultResponse)

	return err
}

type NullableGetSessionResultResponse struct {
	value *GetSessionResultResponse
	isSet bool
}

func (v NullableGetSessionResultResponse) Get() *GetSessionResultResponse {
	return v.value
}

func (v *NullableGetSessionResultResponse) Set(val *GetSessionResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSessionResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSessionResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSessionResultResponse(val *GetSessionResultResponse) *NullableGetSessionResultResponse {
	return &NullableGetSessionResultResponse{value: val, isSet: true}
}

func (v NullableGetSessionResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSessionResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


