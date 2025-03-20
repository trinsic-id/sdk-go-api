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

// checks if the GetAttachmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAttachmentRequest{}

// GetAttachmentRequest struct for GetAttachmentRequest
type GetAttachmentRequest struct {
	// The Attachment Access Key to exchange for the raw file contents of the related Attachment
	AttachmentAccessKey string `json:"attachmentAccessKey"`
}

type _GetAttachmentRequest GetAttachmentRequest

// NewGetAttachmentRequest instantiates a new GetAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAttachmentRequest(attachmentAccessKey string) *GetAttachmentRequest {
	this := GetAttachmentRequest{}
	this.AttachmentAccessKey = attachmentAccessKey
	return &this
}

// NewGetAttachmentRequestWithDefaults instantiates a new GetAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAttachmentRequestWithDefaults() *GetAttachmentRequest {
	this := GetAttachmentRequest{}
	return &this
}

// GetAttachmentAccessKey returns the AttachmentAccessKey field value
func (o *GetAttachmentRequest) GetAttachmentAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttachmentAccessKey
}

// GetAttachmentAccessKeyOk returns a tuple with the AttachmentAccessKey field value
// and a boolean to check if the value has been set.
func (o *GetAttachmentRequest) GetAttachmentAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachmentAccessKey, true
}

// SetAttachmentAccessKey sets field value
func (o *GetAttachmentRequest) SetAttachmentAccessKey(v string) {
	o.AttachmentAccessKey = v
}

func (o GetAttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAttachmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attachmentAccessKey"] = o.AttachmentAccessKey
	return toSerialize, nil
}

func (o *GetAttachmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attachmentAccessKey",
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

	varGetAttachmentRequest := _GetAttachmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAttachmentRequest)

	if err != nil {
		return err
	}

	*o = GetAttachmentRequest(varGetAttachmentRequest)

	return err
}

type NullableGetAttachmentRequest struct {
	value *GetAttachmentRequest
	isSet bool
}

func (v NullableGetAttachmentRequest) Get() *GetAttachmentRequest {
	return v.value
}

func (v *NullableGetAttachmentRequest) Set(val *GetAttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAttachmentRequest(val *GetAttachmentRequest) *NullableGetAttachmentRequest {
	return &NullableGetAttachmentRequest{value: val, isSet: true}
}

func (v NullableGetAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


