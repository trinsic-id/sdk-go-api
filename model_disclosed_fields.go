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

// checks if the DisclosedFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisclosedFields{}

// DisclosedFields struct for DisclosedFields
type DisclosedFields struct {
	IdNumber bool `json:"idNumber"`
	GivenName bool `json:"givenName"`
	FamilyName bool `json:"familyName"`
	Address bool `json:"address"`
	DateOfBirth bool `json:"dateOfBirth"`
	Country bool `json:"country"`
	IssueDate bool `json:"issueDate"`
	ExpirationDate bool `json:"expirationDate"`
	DocumentFront bool `json:"documentFront"`
	DocumentBack bool `json:"documentBack"`
	DocumentPortrait bool `json:"documentPortrait"`
	Selfie bool `json:"selfie"`
}

type _DisclosedFields DisclosedFields

// NewDisclosedFields instantiates a new DisclosedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisclosedFields(idNumber bool, givenName bool, familyName bool, address bool, dateOfBirth bool, country bool, issueDate bool, expirationDate bool, documentFront bool, documentBack bool, documentPortrait bool, selfie bool) *DisclosedFields {
	this := DisclosedFields{}
	this.IdNumber = idNumber
	this.GivenName = givenName
	this.FamilyName = familyName
	this.Address = address
	this.DateOfBirth = dateOfBirth
	this.Country = country
	this.IssueDate = issueDate
	this.ExpirationDate = expirationDate
	this.DocumentFront = documentFront
	this.DocumentBack = documentBack
	this.DocumentPortrait = documentPortrait
	this.Selfie = selfie
	return &this
}

// NewDisclosedFieldsWithDefaults instantiates a new DisclosedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisclosedFieldsWithDefaults() *DisclosedFields {
	this := DisclosedFields{}
	return &this
}

// GetIdNumber returns the IdNumber field value
func (o *DisclosedFields) GetIdNumber() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IdNumber
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetIdNumberOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdNumber, true
}

// SetIdNumber sets field value
func (o *DisclosedFields) SetIdNumber(v bool) {
	o.IdNumber = v
}

// GetGivenName returns the GivenName field value
func (o *DisclosedFields) GetGivenName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetGivenNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *DisclosedFields) SetGivenName(v bool) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *DisclosedFields) GetFamilyName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetFamilyNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *DisclosedFields) SetFamilyName(v bool) {
	o.FamilyName = v
}

// GetAddress returns the Address field value
func (o *DisclosedFields) GetAddress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetAddressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DisclosedFields) SetAddress(v bool) {
	o.Address = v
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *DisclosedFields) GetDateOfBirth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetDateOfBirthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *DisclosedFields) SetDateOfBirth(v bool) {
	o.DateOfBirth = v
}

// GetCountry returns the Country field value
func (o *DisclosedFields) GetCountry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetCountryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DisclosedFields) SetCountry(v bool) {
	o.Country = v
}

// GetIssueDate returns the IssueDate field value
func (o *DisclosedFields) GetIssueDate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetIssueDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *DisclosedFields) SetIssueDate(v bool) {
	o.IssueDate = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *DisclosedFields) GetExpirationDate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetExpirationDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *DisclosedFields) SetExpirationDate(v bool) {
	o.ExpirationDate = v
}

// GetDocumentFront returns the DocumentFront field value
func (o *DisclosedFields) GetDocumentFront() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentFront
}

// GetDocumentFrontOk returns a tuple with the DocumentFront field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetDocumentFrontOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentFront, true
}

// SetDocumentFront sets field value
func (o *DisclosedFields) SetDocumentFront(v bool) {
	o.DocumentFront = v
}

// GetDocumentBack returns the DocumentBack field value
func (o *DisclosedFields) GetDocumentBack() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentBack
}

// GetDocumentBackOk returns a tuple with the DocumentBack field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetDocumentBackOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentBack, true
}

// SetDocumentBack sets field value
func (o *DisclosedFields) SetDocumentBack(v bool) {
	o.DocumentBack = v
}

// GetDocumentPortrait returns the DocumentPortrait field value
func (o *DisclosedFields) GetDocumentPortrait() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocumentPortrait
}

// GetDocumentPortraitOk returns a tuple with the DocumentPortrait field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetDocumentPortraitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentPortrait, true
}

// SetDocumentPortrait sets field value
func (o *DisclosedFields) SetDocumentPortrait(v bool) {
	o.DocumentPortrait = v
}

// GetSelfie returns the Selfie field value
func (o *DisclosedFields) GetSelfie() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Selfie
}

// GetSelfieOk returns a tuple with the Selfie field value
// and a boolean to check if the value has been set.
func (o *DisclosedFields) GetSelfieOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selfie, true
}

// SetSelfie sets field value
func (o *DisclosedFields) SetSelfie(v bool) {
	o.Selfie = v
}

func (o DisclosedFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisclosedFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["idNumber"] = o.IdNumber
	toSerialize["givenName"] = o.GivenName
	toSerialize["familyName"] = o.FamilyName
	toSerialize["address"] = o.Address
	toSerialize["dateOfBirth"] = o.DateOfBirth
	toSerialize["country"] = o.Country
	toSerialize["issueDate"] = o.IssueDate
	toSerialize["expirationDate"] = o.ExpirationDate
	toSerialize["documentFront"] = o.DocumentFront
	toSerialize["documentBack"] = o.DocumentBack
	toSerialize["documentPortrait"] = o.DocumentPortrait
	toSerialize["selfie"] = o.Selfie
	return toSerialize, nil
}

func (o *DisclosedFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"idNumber",
		"givenName",
		"familyName",
		"address",
		"dateOfBirth",
		"country",
		"issueDate",
		"expirationDate",
		"documentFront",
		"documentBack",
		"documentPortrait",
		"selfie",
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

	varDisclosedFields := _DisclosedFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisclosedFields)

	if err != nil {
		return err
	}

	*o = DisclosedFields(varDisclosedFields)

	return err
}

type NullableDisclosedFields struct {
	value *DisclosedFields
	isSet bool
}

func (v NullableDisclosedFields) Get() *DisclosedFields {
	return v.value
}

func (v *NullableDisclosedFields) Set(val *DisclosedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableDisclosedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableDisclosedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisclosedFields(val *DisclosedFields) *NullableDisclosedFields {
	return &NullableDisclosedFields{value: val, isSet: true}
}

func (v NullableDisclosedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisclosedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


