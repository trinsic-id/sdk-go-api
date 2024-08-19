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

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session struct for Session
type Session struct {
	Id string `json:"id"`
	// The state of the session
	State SessionState `json:"state"`
	// If the session is in state `IdvFailed`, this field contains the reason for failure.
	FailCode *SessionFailCode `json:"failCode,omitempty"`
	// The underlying verification for this Session
	Verification Verification `json:"verification"`
	// The fields that were requested to be disclosed when the Session was created
	DisclosedFields DisclosedFields `json:"disclosedFields"`
	// The unix timestamp, in seconds, when this session was created
	Created int64 `json:"created"`
	// The unix timestamp, in seconds, when this session's state last changed
	Updated int64 `json:"updated"`
}

type _Session Session

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(id string, state SessionState, verification Verification, disclosedFields DisclosedFields, created int64, updated int64) *Session {
	this := Session{}
	this.Id = id
	this.State = state
	this.Verification = verification
	this.DisclosedFields = disclosedFields
	this.Created = created
	this.Updated = updated
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *Session) GetState() SessionState {
	if o == nil {
		var ret SessionState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Session) GetStateOk() (*SessionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Session) SetState(v SessionState) {
	o.State = v
}

// GetFailCode returns the FailCode field value if set, zero value otherwise.
func (o *Session) GetFailCode() SessionFailCode {
	if o == nil || IsNil(o.FailCode) {
		var ret SessionFailCode
		return ret
	}
	return *o.FailCode
}

// GetFailCodeOk returns a tuple with the FailCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetFailCodeOk() (*SessionFailCode, bool) {
	if o == nil || IsNil(o.FailCode) {
		return nil, false
	}
	return o.FailCode, true
}

// HasFailCode returns a boolean if a field has been set.
func (o *Session) HasFailCode() bool {
	if o != nil && !IsNil(o.FailCode) {
		return true
	}

	return false
}

// SetFailCode gets a reference to the given SessionFailCode and assigns it to the FailCode field.
func (o *Session) SetFailCode(v SessionFailCode) {
	o.FailCode = &v
}

// GetVerification returns the Verification field value
func (o *Session) GetVerification() Verification {
	if o == nil {
		var ret Verification
		return ret
	}

	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
func (o *Session) GetVerificationOk() (*Verification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verification, true
}

// SetVerification sets field value
func (o *Session) SetVerification(v Verification) {
	o.Verification = v
}

// GetDisclosedFields returns the DisclosedFields field value
func (o *Session) GetDisclosedFields() DisclosedFields {
	if o == nil {
		var ret DisclosedFields
		return ret
	}

	return o.DisclosedFields
}

// GetDisclosedFieldsOk returns a tuple with the DisclosedFields field value
// and a boolean to check if the value has been set.
func (o *Session) GetDisclosedFieldsOk() (*DisclosedFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisclosedFields, true
}

// SetDisclosedFields sets field value
func (o *Session) SetDisclosedFields(v DisclosedFields) {
	o.DisclosedFields = v
}

// GetCreated returns the Created field value
func (o *Session) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Session) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Session) SetCreated(v int64) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Session) GetUpdated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Session) GetUpdatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Session) SetUpdated(v int64) {
	o.Updated = v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["state"] = o.State
	if !IsNil(o.FailCode) {
		toSerialize["failCode"] = o.FailCode
	}
	toSerialize["verification"] = o.Verification
	toSerialize["disclosedFields"] = o.DisclosedFields
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

func (o *Session) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"state",
		"verification",
		"disclosedFields",
		"created",
		"updated",
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

	varSession := _Session{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSession)

	if err != nil {
		return err
	}

	*o = Session(varSession)

	return err
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


