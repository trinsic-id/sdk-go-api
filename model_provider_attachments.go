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

// checks if the ProviderAttachments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderAttachments{}

// ProviderAttachments struct for ProviderAttachments
type ProviderAttachments struct {
	DocumentScan NullableDocumentScanAttachments `json:"document-scan,omitempty"`
}

// NewProviderAttachments instantiates a new ProviderAttachments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderAttachments() *ProviderAttachments {
	this := ProviderAttachments{}
	return &this
}

// NewProviderAttachmentsWithDefaults instantiates a new ProviderAttachments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderAttachmentsWithDefaults() *ProviderAttachments {
	this := ProviderAttachments{}
	return &this
}

// GetDocumentScan returns the DocumentScan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProviderAttachments) GetDocumentScan() DocumentScanAttachments {
	if o == nil || IsNil(o.DocumentScan.Get()) {
		var ret DocumentScanAttachments
		return ret
	}
	return *o.DocumentScan.Get()
}

// GetDocumentScanOk returns a tuple with the DocumentScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProviderAttachments) GetDocumentScanOk() (*DocumentScanAttachments, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentScan.Get(), o.DocumentScan.IsSet()
}

// HasDocumentScan returns a boolean if a field has been set.
func (o *ProviderAttachments) HasDocumentScan() bool {
	if o != nil && o.DocumentScan.IsSet() {
		return true
	}

	return false
}

// SetDocumentScan gets a reference to the given NullableDocumentScanAttachments and assigns it to the DocumentScan field.
func (o *ProviderAttachments) SetDocumentScan(v DocumentScanAttachments) {
	o.DocumentScan.Set(&v)
}
// SetDocumentScanNil sets the value for DocumentScan to be an explicit nil
func (o *ProviderAttachments) SetDocumentScanNil() {
	o.DocumentScan.Set(nil)
}

// UnsetDocumentScan ensures that no value is present for DocumentScan, not even an explicit nil
func (o *ProviderAttachments) UnsetDocumentScan() {
	o.DocumentScan.Unset()
}

func (o ProviderAttachments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderAttachments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentScan.IsSet() {
		toSerialize["document-scan"] = o.DocumentScan.Get()
	}
	return toSerialize, nil
}

type NullableProviderAttachments struct {
	value *ProviderAttachments
	isSet bool
}

func (v NullableProviderAttachments) Get() *ProviderAttachments {
	return v.value
}

func (v *NullableProviderAttachments) Set(val *ProviderAttachments) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderAttachments) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderAttachments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderAttachments(val *ProviderAttachments) *NullableProviderAttachments {
	return &NullableProviderAttachments{value: val, isSet: true}
}

func (v NullableProviderAttachments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderAttachments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


