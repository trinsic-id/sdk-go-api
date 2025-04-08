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

// checks if the SubProviderMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubProviderMetadata{}

// SubProviderMetadata Information about a sub-provider.
type SubProviderMetadata struct {
	// The ID of the sub-provider.              This cannot be used as a standalone Provider ID when creating a Session. It must be passed in via the Provider-specific input.
	Id string `json:"id"`
	// The name of the sub-provider
	Name string `json:"name"`
	// Flavor text for the sub-provider
	Description string `json:"description"`
	// A URL pointing to the logo on Trinsic's CDN.              May be a PNG, JPG, or SVG image.
	LogoUrl string `json:"logoUrl"`
}

type _SubProviderMetadata SubProviderMetadata

// NewSubProviderMetadata instantiates a new SubProviderMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubProviderMetadata(id string, name string, description string, logoUrl string) *SubProviderMetadata {
	this := SubProviderMetadata{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.LogoUrl = logoUrl
	return &this
}

// NewSubProviderMetadataWithDefaults instantiates a new SubProviderMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubProviderMetadataWithDefaults() *SubProviderMetadata {
	this := SubProviderMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *SubProviderMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubProviderMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubProviderMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SubProviderMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubProviderMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubProviderMetadata) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SubProviderMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SubProviderMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SubProviderMetadata) SetDescription(v string) {
	o.Description = v
}

// GetLogoUrl returns the LogoUrl field value
func (o *SubProviderMetadata) GetLogoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
func (o *SubProviderMetadata) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoUrl, true
}

// SetLogoUrl sets field value
func (o *SubProviderMetadata) SetLogoUrl(v string) {
	o.LogoUrl = v
}

func (o SubProviderMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubProviderMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["logoUrl"] = o.LogoUrl
	return toSerialize, nil
}

func (o *SubProviderMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"logoUrl",
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

	varSubProviderMetadata := _SubProviderMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubProviderMetadata)

	if err != nil {
		return err
	}

	*o = SubProviderMetadata(varSubProviderMetadata)

	return err
}

type NullableSubProviderMetadata struct {
	value *SubProviderMetadata
	isSet bool
}

func (v NullableSubProviderMetadata) Get() *SubProviderMetadata {
	return v.value
}

func (v *NullableSubProviderMetadata) Set(val *SubProviderMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSubProviderMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSubProviderMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubProviderMetadata(val *SubProviderMetadata) *NullableSubProviderMetadata {
	return &NullableSubProviderMetadata{value: val, isSet: true}
}

func (v NullableSubProviderMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubProviderMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


