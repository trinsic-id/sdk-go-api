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

// checks if the IdentityData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityData{}

// IdentityData struct for IdentityData
type IdentityData struct {
	OriginatingProviderId NullableString `json:"originatingProviderId,omitempty"`
	Person NullablePersonData `json:"person,omitempty"`
	Document NullableDocumentData `json:"document,omitempty"`
	AttachmentAccessKeys NullableAttachmentAccessKeys `json:"attachmentAccessKeys,omitempty"`
}

// NewIdentityData instantiates a new IdentityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityData() *IdentityData {
	this := IdentityData{}
	return &this
}

// NewIdentityDataWithDefaults instantiates a new IdentityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityDataWithDefaults() *IdentityData {
	this := IdentityData{}
	return &this
}

// GetOriginatingProviderId returns the OriginatingProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityData) GetOriginatingProviderId() string {
	if o == nil || IsNil(o.OriginatingProviderId.Get()) {
		var ret string
		return ret
	}
	return *o.OriginatingProviderId.Get()
}

// GetOriginatingProviderIdOk returns a tuple with the OriginatingProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityData) GetOriginatingProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginatingProviderId.Get(), o.OriginatingProviderId.IsSet()
}

// HasOriginatingProviderId returns a boolean if a field has been set.
func (o *IdentityData) HasOriginatingProviderId() bool {
	if o != nil && o.OriginatingProviderId.IsSet() {
		return true
	}

	return false
}

// SetOriginatingProviderId gets a reference to the given NullableString and assigns it to the OriginatingProviderId field.
func (o *IdentityData) SetOriginatingProviderId(v string) {
	o.OriginatingProviderId.Set(&v)
}
// SetOriginatingProviderIdNil sets the value for OriginatingProviderId to be an explicit nil
func (o *IdentityData) SetOriginatingProviderIdNil() {
	o.OriginatingProviderId.Set(nil)
}

// UnsetOriginatingProviderId ensures that no value is present for OriginatingProviderId, not even an explicit nil
func (o *IdentityData) UnsetOriginatingProviderId() {
	o.OriginatingProviderId.Unset()
}

// GetPerson returns the Person field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityData) GetPerson() PersonData {
	if o == nil || IsNil(o.Person.Get()) {
		var ret PersonData
		return ret
	}
	return *o.Person.Get()
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityData) GetPersonOk() (*PersonData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Person.Get(), o.Person.IsSet()
}

// HasPerson returns a boolean if a field has been set.
func (o *IdentityData) HasPerson() bool {
	if o != nil && o.Person.IsSet() {
		return true
	}

	return false
}

// SetPerson gets a reference to the given NullablePersonData and assigns it to the Person field.
func (o *IdentityData) SetPerson(v PersonData) {
	o.Person.Set(&v)
}
// SetPersonNil sets the value for Person to be an explicit nil
func (o *IdentityData) SetPersonNil() {
	o.Person.Set(nil)
}

// UnsetPerson ensures that no value is present for Person, not even an explicit nil
func (o *IdentityData) UnsetPerson() {
	o.Person.Unset()
}

// GetDocument returns the Document field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityData) GetDocument() DocumentData {
	if o == nil || IsNil(o.Document.Get()) {
		var ret DocumentData
		return ret
	}
	return *o.Document.Get()
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityData) GetDocumentOk() (*DocumentData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Document.Get(), o.Document.IsSet()
}

// HasDocument returns a boolean if a field has been set.
func (o *IdentityData) HasDocument() bool {
	if o != nil && o.Document.IsSet() {
		return true
	}

	return false
}

// SetDocument gets a reference to the given NullableDocumentData and assigns it to the Document field.
func (o *IdentityData) SetDocument(v DocumentData) {
	o.Document.Set(&v)
}
// SetDocumentNil sets the value for Document to be an explicit nil
func (o *IdentityData) SetDocumentNil() {
	o.Document.Set(nil)
}

// UnsetDocument ensures that no value is present for Document, not even an explicit nil
func (o *IdentityData) UnsetDocument() {
	o.Document.Unset()
}

// GetAttachmentAccessKeys returns the AttachmentAccessKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityData) GetAttachmentAccessKeys() AttachmentAccessKeys {
	if o == nil || IsNil(o.AttachmentAccessKeys.Get()) {
		var ret AttachmentAccessKeys
		return ret
	}
	return *o.AttachmentAccessKeys.Get()
}

// GetAttachmentAccessKeysOk returns a tuple with the AttachmentAccessKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityData) GetAttachmentAccessKeysOk() (*AttachmentAccessKeys, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentAccessKeys.Get(), o.AttachmentAccessKeys.IsSet()
}

// HasAttachmentAccessKeys returns a boolean if a field has been set.
func (o *IdentityData) HasAttachmentAccessKeys() bool {
	if o != nil && o.AttachmentAccessKeys.IsSet() {
		return true
	}

	return false
}

// SetAttachmentAccessKeys gets a reference to the given NullableAttachmentAccessKeys and assigns it to the AttachmentAccessKeys field.
func (o *IdentityData) SetAttachmentAccessKeys(v AttachmentAccessKeys) {
	o.AttachmentAccessKeys.Set(&v)
}
// SetAttachmentAccessKeysNil sets the value for AttachmentAccessKeys to be an explicit nil
func (o *IdentityData) SetAttachmentAccessKeysNil() {
	o.AttachmentAccessKeys.Set(nil)
}

// UnsetAttachmentAccessKeys ensures that no value is present for AttachmentAccessKeys, not even an explicit nil
func (o *IdentityData) UnsetAttachmentAccessKeys() {
	o.AttachmentAccessKeys.Unset()
}

func (o IdentityData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OriginatingProviderId.IsSet() {
		toSerialize["originatingProviderId"] = o.OriginatingProviderId.Get()
	}
	if o.Person.IsSet() {
		toSerialize["person"] = o.Person.Get()
	}
	if o.Document.IsSet() {
		toSerialize["document"] = o.Document.Get()
	}
	if o.AttachmentAccessKeys.IsSet() {
		toSerialize["attachmentAccessKeys"] = o.AttachmentAccessKeys.Get()
	}
	return toSerialize, nil
}

type NullableIdentityData struct {
	value *IdentityData
	isSet bool
}

func (v NullableIdentityData) Get() *IdentityData {
	return v.value
}

func (v *NullableIdentityData) Set(val *IdentityData) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityData) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityData(val *IdentityData) *NullableIdentityData {
	return &NullableIdentityData{value: val, isSet: true}
}

func (v NullableIdentityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


