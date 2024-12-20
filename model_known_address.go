/*
Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package trinsic_api

import (
	"encoding/json"
)

// checks if the KnownAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KnownAddress{}

// KnownAddress Address information for an individual
type KnownAddress struct {
	Line1 *string `json:"line1,omitempty"`
	Line2 *string `json:"line2,omitempty"`
	Line3 *string `json:"line3,omitempty"`
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewKnownAddress instantiates a new KnownAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKnownAddress() *KnownAddress {
	this := KnownAddress{}
	return &this
}

// NewKnownAddressWithDefaults instantiates a new KnownAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKnownAddressWithDefaults() *KnownAddress {
	this := KnownAddress{}
	return &this
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *KnownAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine1() bool {
	if o != nil && !IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *KnownAddress) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *KnownAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine2() bool {
	if o != nil && !IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *KnownAddress) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *KnownAddress) GetLine3() string {
	if o == nil || IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine3() bool {
	if o != nil && !IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *KnownAddress) SetLine3(v string) {
	o.Line3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *KnownAddress) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *KnownAddress) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *KnownAddress) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *KnownAddress) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *KnownAddress) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *KnownAddress) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *KnownAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *KnownAddress) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *KnownAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *KnownAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KnownAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *KnownAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *KnownAddress) SetCountry(v string) {
	o.Country = &v
}

func (o KnownAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KnownAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Line1) {
		toSerialize["line1"] = o.Line1
	}
	if !IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !IsNil(o.Line3) {
		toSerialize["line3"] = o.Line3
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableKnownAddress struct {
	value *KnownAddress
	isSet bool
}

func (v NullableKnownAddress) Get() *KnownAddress {
	return v.value
}

func (v *NullableKnownAddress) Set(val *KnownAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableKnownAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableKnownAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKnownAddress(val *KnownAddress) *NullableKnownAddress {
	return &NullableKnownAddress{value: val, isSet: true}
}

func (v NullableKnownAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKnownAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


