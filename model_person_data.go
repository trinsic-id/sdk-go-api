/*
Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connect

import (
	"encoding/json"
)

// checks if the PersonData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonData{}

// PersonData Identity information for the individual being verified
type PersonData struct {
	// Given (first) name of the individual
	GivenName *string `json:"givenName,omitempty"`
	// Family (last) name of the individual
	FamilyName *string `json:"familyName,omitempty"`
	// Middle name of the individual
	MiddleName *string `json:"middleName,omitempty"`
	// The individual's full name as a single string.                Useful for names which do not fit into a \"first middle last\" structure.
	FullName *string `json:"fullName,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	Gender *string `json:"gender,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Address information for an individual
	Address *Address `json:"address,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
}

// NewPersonData instantiates a new PersonData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonData() *PersonData {
	this := PersonData{}
	return &this
}

// NewPersonDataWithDefaults instantiates a new PersonData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonDataWithDefaults() *PersonData {
	this := PersonData{}
	return &this
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *PersonData) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *PersonData) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *PersonData) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *PersonData) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *PersonData) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *PersonData) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *PersonData) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *PersonData) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *PersonData) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *PersonData) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *PersonData) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *PersonData) SetFullName(v string) {
	o.FullName = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *PersonData) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *PersonData) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *PersonData) SetNationality(v string) {
	o.Nationality = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *PersonData) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *PersonData) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *PersonData) SetGender(v string) {
	o.Gender = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PersonData) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PersonData) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PersonData) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PersonData) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PersonData) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *PersonData) SetAddress(v Address) {
	o.Address = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *PersonData) GetDateOfBirth() string {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonData) GetDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *PersonData) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *PersonData) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

func (o PersonData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	return toSerialize, nil
}

type NullablePersonData struct {
	value *PersonData
	isSet bool
}

func (v NullablePersonData) Get() *PersonData {
	return v.value
}

func (v *NullablePersonData) Set(val *PersonData) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonData) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonData(val *PersonData) *NullablePersonData {
	return &NullablePersonData{value: val, isSet: true}
}

func (v NullablePersonData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


