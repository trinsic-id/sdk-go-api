/*
Trinsic API

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
	Line1 NullableString `json:"line1,omitempty"`
	Line2 NullableString `json:"line2,omitempty"`
	Line3 NullableString `json:"line3,omitempty"`
	City NullableString `json:"city,omitempty"`
	Subdivision NullableString `json:"subdivision,omitempty"`
	// Deprecated. Use `Subdivision` instead.
	// Deprecated
	State NullableString `json:"state,omitempty"`
	PostalCode NullableString `json:"postalCode,omitempty"`
	Country NullableString `json:"country,omitempty"`
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

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *KnownAddress) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *KnownAddress) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *KnownAddress) UnsetLine1() {
	o.Line1.Unset()
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetLine2() string {
	if o == nil || IsNil(o.Line2.Get()) {
		var ret string
		return ret
	}
	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// HasLine2 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine2() bool {
	if o != nil && o.Line2.IsSet() {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given NullableString and assigns it to the Line2 field.
func (o *KnownAddress) SetLine2(v string) {
	o.Line2.Set(&v)
}
// SetLine2Nil sets the value for Line2 to be an explicit nil
func (o *KnownAddress) SetLine2Nil() {
	o.Line2.Set(nil)
}

// UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
func (o *KnownAddress) UnsetLine2() {
	o.Line2.Unset()
}

// GetLine3 returns the Line3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetLine3() string {
	if o == nil || IsNil(o.Line3.Get()) {
		var ret string
		return ret
	}
	return *o.Line3.Get()
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line3.Get(), o.Line3.IsSet()
}

// HasLine3 returns a boolean if a field has been set.
func (o *KnownAddress) HasLine3() bool {
	if o != nil && o.Line3.IsSet() {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given NullableString and assigns it to the Line3 field.
func (o *KnownAddress) SetLine3(v string) {
	o.Line3.Set(&v)
}
// SetLine3Nil sets the value for Line3 to be an explicit nil
func (o *KnownAddress) SetLine3Nil() {
	o.Line3.Set(nil)
}

// UnsetLine3 ensures that no value is present for Line3, not even an explicit nil
func (o *KnownAddress) UnsetLine3() {
	o.Line3.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *KnownAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *KnownAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *KnownAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *KnownAddress) UnsetCity() {
	o.City.Unset()
}

// GetSubdivision returns the Subdivision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetSubdivision() string {
	if o == nil || IsNil(o.Subdivision.Get()) {
		var ret string
		return ret
	}
	return *o.Subdivision.Get()
}

// GetSubdivisionOk returns a tuple with the Subdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetSubdivisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subdivision.Get(), o.Subdivision.IsSet()
}

// HasSubdivision returns a boolean if a field has been set.
func (o *KnownAddress) HasSubdivision() bool {
	if o != nil && o.Subdivision.IsSet() {
		return true
	}

	return false
}

// SetSubdivision gets a reference to the given NullableString and assigns it to the Subdivision field.
func (o *KnownAddress) SetSubdivision(v string) {
	o.Subdivision.Set(&v)
}
// SetSubdivisionNil sets the value for Subdivision to be an explicit nil
func (o *KnownAddress) SetSubdivisionNil() {
	o.Subdivision.Set(nil)
}

// UnsetSubdivision ensures that no value is present for Subdivision, not even an explicit nil
func (o *KnownAddress) UnsetSubdivision() {
	o.Subdivision.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *KnownAddress) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *KnownAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *KnownAddress) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
// Deprecated
func (o *KnownAddress) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *KnownAddress) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *KnownAddress) UnsetState() {
	o.State.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *KnownAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *KnownAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *KnownAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *KnownAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KnownAddress) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KnownAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *KnownAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *KnownAddress) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *KnownAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *KnownAddress) UnsetCountry() {
	o.Country.Unset()
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
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.Line2.IsSet() {
		toSerialize["line2"] = o.Line2.Get()
	}
	if o.Line3.IsSet() {
		toSerialize["line3"] = o.Line3.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Subdivision.IsSet() {
		toSerialize["subdivision"] = o.Subdivision.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postalCode"] = o.PostalCode.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
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


